package texas

import (
	"github.com/zhang1career/golab/calc"
	"github.com/zhang1career/golab/cardgame"
	"github.com/zhang1career/golab/cardgame/poker"
	"github.com/zhang1career/golab/log"
)

func (this *Texas) Evaluate(originCards []cardgame.Card) (winCate int, winScore float64) {
	// 皇家同花顺
	if has, cards := this.HasRoyalStraightFlush(originCards); has {
		score, _, _ := this.CalcScore(cards, 1)
		return RoyalStraightFlush, calc.Normalize(score, 5, 14)
	}
	// 同花顺
	if has, cards := this.HasStraightFlush(originCards); has {
		score, _, _ := this.CalcScore(cards, 1)
		return StraightFlush, calc.Normalize(score, 5, 14)
	}
	// 四条
	if has, cards := this.HasFourOfAKind(originCards); has {
		score, _, _ := this.CalcScore(cards, 1)
		return FourOfAKind, calc.Normalize(score, 2, 14)
	}
	// 葫芦
	if has, pairs := this.HasFullHouse(originCards); has {
		sum := 0
		for _, pair := range pairs {
			sum = sum << 4
			score, _, _ := this.CalcScore(pair, 1)
			sum = sum + score - 1
		}
		return FullHouse, calc.Normalize(sum, 0, 256)
	}
	// 同花
	if has, cards := this.HasFlush(originCards); has {
		sum := 0
		for i := 0; i < 5; i++ {
			sum = sum << 4
			score, _, _ := this.CalcScore(cards[i:i+1], 1)
			sum = sum + score - 1
		}
		return Flush, calc.Normalize(sum, 0, 1048576)
	}
	// 顺子
	if has, straights := this.HasStraight(originCards); has {
		cards := make([]cardgame.Card, 0)
		for _, straight := range straights {
			cards = append(cards, straight...)
		}
		score, _, _ := this.CalcScore(cards, 1)
		return Straight, calc.Normalize(score, 5, 14)
	}
	// 三条
	if has, cards := this.HasThreeOfAKind(originCards); has {
		score, _, _ := this.CalcScore(cards, 1)
		return ThreeOfAKind, calc.Normalize(score, 2, 14)
	}
	// 两对
	if has, pairs := this.HasTwoPair(originCards); has {
		sum := 0
		for _, pair := range pairs {
			sum = sum << 4
			score, _, _ := this.CalcScore(pair, 1)
			sum = sum + score - 1
		}
		return TwoPair, calc.Normalize(sum, 0, 256)
	}
	// 一对
	if has, cards := this.HasOnePair(originCards); has {
		score, _, _ := this.CalcScore(cards, 1)
		return OnePair, calc.Normalize(score, 2, 14)
	}
	// 高牌
	sum := 0
	for i := 0; i < 5; i++ {
		sum = sum << 4
		score, _, _ := this.CalcScore(originCards[i:i+1], 1)
		sum = sum + score - 1
	}
	return HighCard, calc.Normalize(sum, 0, 1048576)
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