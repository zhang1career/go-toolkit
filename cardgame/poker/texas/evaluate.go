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

func HasStraight(cards []cardgame.Card) (bool, int) {
	game, err := poker.New()
	if err != nil {
		log.Error(err.Error())
	}
	kvs := game.GroupByValue(cards).Sort("desc")
	
	return game.HasStraight(kvs, 5)
}

func HasFlush(cards []cardgame.Card) (bool, int) {
	game, err := poker.New()
	if err != nil {
		log.Error(err.Error())
	}
	
	num, suits := game.HasSuit(cards, 5)
	if num <= 0 {
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

// @todo HasStraight return cards
func HasStraightFlush(cards []cardgame.Card) (bool, []int) {
	ret := make([]int, 0)
	
	hasStraight, startValue := HasStraight(cards)
	if hasStraight == false {
		return false, ret
	}
	ret = append(ret, startValue)
	
	hasFlush, suit := HasFlush(cards)
	if hasFlush == false {
		return false, ret
	}
	ret = append(ret, suit)
	
	return true, ret
}

// royal straight flush

