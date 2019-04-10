package texas

import (
	"github.com/zhang1career/lib/cardgame"
	"github.com/zhang1career/lib/cardgame/poker"
	"github.com/zhang1career/lib/log"
)

func Evaluate(own_cards []cardgame.Card, common_cards []cardgame.Card) (int, int) {
	game, err := New()
	if err != nil {
		log.Error(err.Error())
	}
	originCards := append(common_cards, own_cards...)
	// 皇家同花顺
	if has, cards := game.HasRoyalStraightFlush(originCards); has {
		score, _ := game.CalcSccore(cards, 5)
		return 9, score
	}
	// 同花顺
	if has, cards := game.HasStraightFlush(originCards); has {
		score, _ := game.CalcSccore(cards, 5)
		return 8, score
	}
	// 四条
	if has, cards := game.HasFourOfAKind(originCards); has {
		score, _ := game.CalcSccore(cards, 5)
		return 7, score
	}
	// 葫芦
	if has, pairs := game.HasFullHouse(originCards); has {
		cards := make([]cardgame.Card, 0)
		for _, pair := range pairs {
			cards = append(cards, pair...)
		}
		score, _ := game.CalcSccore(cards, 5)
		return 6, score
	}
	// 同花
	if has, cards := game.HasFlush(originCards); has {
		score, _ := game.CalcSccore(cards, 5)
		return 5, score
	}
	// 顺子
	if has, straights := game.HasStraight(originCards); has {
		cards := make([]cardgame.Card, 0)
		for _, straight := range straights {
			cards = append(cards, straight...)
		}
		score, _ := game.CalcSccore(cards, 5)
		return 4, score
	}
	// 三条
	if has, cards := game.HasThreeOfAKind(originCards); has {
		score, _ := game.CalcSccore(cards, 5)
		return 3, score
	}
	// 两对
	if has, pairs := game.HasTwoPair(originCards); has {
		cards := make([]cardgame.Card, 0)
		for _, pair := range pairs {
			cards = append(cards, pair...)
		}
		score, _ := game.CalcSccore(cards, 5)
		return 2, score
	}
	// 对子
	if has, cards := game.HasOnePair(originCards); has {
		score, _ := game.CalcSccore(cards, 5)
		return 1, score
	}
	// 高牌
	score, _ := game.CalcSccore(originCards, 5)
	return 0, score
}

func (this *Texas) HasOnePair(cards []cardgame.Card) (bool, []cardgame.Card) {
	hasPair, pairs := this.HasPair(cards, 2)
	if !hasPair {
		return false, nil
	}
	return true, pairs[0]
}

func (this *Texas) HasTwoPair(cards []cardgame.Card) (bool, [][]cardgame.Card) {
	hasPair, pair0 := this.HasOnePair(cards)
	if !hasPair {
		return false, nil
	}
	cards = this.DelCards(cards, pair0)
	hasPair, pair1 := this.HasOnePair(cards)
	if !hasPair {
		return false, nil
	}
	return true, [][]cardgame.Card{pair0, pair1}
}

func (this *Texas) HasThreeOfAKind(cards []cardgame.Card) (bool, []cardgame.Card) {
	hasPair, pairs := this.HasPair(cards, 3)
	if !hasPair {
		return false, nil
	}
	return true, pairs[0]
}

func (this *Texas) HasStraight(cards []cardgame.Card) (bool, [][]cardgame.Card) {
	return this.HasSerial(cards, 5)
}

func (this *Texas) HasFlush(cards []cardgame.Card) (bool, []cardgame.Card) {
	hasSuit, suits := this.HasSuit(cards, 5)
	if !hasSuit {
		return false, nil
	}
	ret := this.SortByValue(suits[0], "desc")
	return true, ret
}

func (this *Texas) HasFullHouse(cards []cardgame.Card) (bool, [][]cardgame.Card) {
	hasPair, pair3 := this.HasThreeOfAKind(cards)
	if !hasPair {
		return false, nil
	}
	cards = this.DelCards(cards, pair3)
	hasPair, pair2 := this.HasOnePair(cards)
	if !hasPair {
		return false, nil
	}
	return true, [][]cardgame.Card{pair3, pair2}
}

func (this *Texas) HasFourOfAKind(cards []cardgame.Card) (bool, []cardgame.Card) {
	hasPair, pairs := this.HasPair(cards, 4)
	if !hasPair {
		return false, nil
	}
	return true, pairs[0]
}

func (this *Texas) HasStraightFlush(cards []cardgame.Card) (bool, []cardgame.Card) {
	hasStraight, straights := this.HasStraight(cards)
	if !hasStraight {
		return false, nil
	}
	
	retStraight := make([]cardgame.Card, 0)
	for _, straight := range straights {
		hasFlush, _ := this.HasFlush(straight)
		if hasFlush {
			retStraight = append(retStraight, straight...)
			break
		}
	}
	
	return len(retStraight) > 0, retStraight
}

func (this *Texas) HasRoyalStraightFlush(cards []cardgame.Card) (bool, []cardgame.Card) {
	hasStraight, straight := this.HasStraightFlush(cards)
	if !hasStraight {
		return false, nil
	}
	
	game, err := poker.New()
	if err != nil {
		log.Error(err.Error())
	}
	straight = game.SortByValue(straight, "asc")
	if straight[0].Value != 10 {
		return false, nil
	}
	
	return true, straight
}