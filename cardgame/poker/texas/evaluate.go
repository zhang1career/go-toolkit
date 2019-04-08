package texas

import (
	"github.com/zhang1career/lib/cardgame"
	"github.com/zhang1career/lib/cardgame/poker"
	"github.com/zhang1career/lib/log"
)

//func Evaluate(own_cards []cardgame.Card, common_cards []cardgame.Card) bool {
//	game, err := poker.New()
//	if err != nil {
//		log.Error(err.Error())
//	}
//	cards := append(common_cards, own_cards...)
//	values := game.GroupByValue(cards).Sort("desc")
//
//	has_straight, start_value := game.HasStraight(values, 5)
//
//	return has_straight
//}

// high card

func HasOnePair(cards []cardgame.Card) (bool, int) {
	game, err := poker.New()
	if err != nil {
		log.Error(err.Error())
	}
	kvs := game.GroupByValue(cards).Sort("desc")
	
	num, pairs := game.HasPair(kvs, 2)
	if num <= 0 {
		return false, 0
	}
	
	return true, pairs[0].Key
}

func HasTwoPair(cards []cardgame.Card) (bool, []int) {
	game, err := poker.New()
	if err != nil {
		log.Error(err.Error())
	}
	kvs := game.GroupByValue(cards).Sort("desc")
	
	ret := make([]int, 0)
	
	num, pairs := game.HasPair(kvs, 2)
	if num <= 1 {
		return false, ret
	}
	ret = append(ret, pairs[0].Key, pairs[1].Key)
	
	return true, ret
}

func HasThreeOfAKind(cards []cardgame.Card) (bool, int) {
	game, err := poker.New()
	if err != nil {
		log.Error(err.Error())
	}
	kvs := game.GroupByValue(cards).Sort("desc")
	
	num, pairs := game.HasPair(kvs, 3)
	if num <= 0 {
		return false, 0
	}
	
	return true, pairs[0].Key
}

func HasStraight(cards []cardgame.Card) (bool, [][]cardgame.Card) {
	game, err := poker.New()
	if err != nil {
		log.Error(err.Error())
	}
	
	return game.HasStraight(cards, 5)
}

func HasFlush(cards []cardgame.Card) (bool, int) {
	game, err := poker.New()
	if err != nil {
		log.Error(err.Error())
	}
	
	hasSuit, suits := game.HasSuit(cards, 5)
	if !hasSuit {
		return false, 0
	}
	
	return true, suits[0]
}

func HasFullHouse(cards []cardgame.Card) (bool, []int) {
	game, err := poker.New()
	if err != nil {
		log.Error(err.Error())
	}
	kvs := game.GroupByValue(cards).Sort("desc")
	
	ret := make([]int, 0)
	
	threeNum, threePairs := game.HasPair(kvs, 3)
	if threeNum <= 0 {
		return false, ret
	}
	ret = append(ret, threePairs[0].Key)
	
	twoNum, twoPairs := game.HasPair(kvs, 2)
	if twoNum <= 0 {
		return false, ret
	}
	ret = append(ret, twoPairs[0].Key)
	
	return true, ret
}

func HasFourOfAKind(cards []cardgame.Card) (bool, int) {
	game, err := poker.New()
	if err != nil {
		log.Error(err.Error())
	}
	kvs := game.GroupByValue(cards).Sort("desc")
	
	num, pairs := game.HasPair(kvs, 4)
	if num <= 0 {
		return false, 0
	}
	
	return true, pairs[0].Key
}

func HasStraightFlush(cards []cardgame.Card) (bool, []cardgame.Card) {
	hasStraight, straights := HasStraight(cards)
	if !hasStraight {
		return false, nil
	}
	
	retStraight := make([]cardgame.Card, 0)
	for _, straight := range straights {
		hasFlush, _ := HasFlush(straight)
		if hasFlush {
			retStraight = append(retStraight, straight...)
			break
		}
	}
	
	return len(retStraight) > 0, retStraight
}

func HasRoyalStraightFlush(cards []cardgame.Card) (bool, []cardgame.Card) {
	hasStraight, straight := HasStraightFlush(cards)
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