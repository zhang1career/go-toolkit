package poker

import (
	"github.com/zhang1career/lib/cardgame"
)

func (this *Poker) HasPair(cards []cardgame.Card, count int) (bool, [][]cardgame.Card) {
	cards = this.prepareValues(cards)
	// pooling
	maybePair := [][]cardgame.Card{{cards[0]}}
	
	for i := 1; i < len(cards); i++ {
		card := cards[i]
		// add to last maybe
		if cards[i-1].Value == card.Value {
			maybePair[len(maybePair)-1] = append(maybePair[len(maybePair)-1], card)
			continue
		}
		// delete last maybe when not a pair
		if len(maybePair[len(maybePair)-1]) < count {
			maybePair = maybePair[0:len(maybePair)-1]
		}
		maybePair = append(maybePair, []cardgame.Card{card})
	}
	// last check
	if len(maybePair[len(maybePair)-1]) < count {
		maybePair = maybePair[0:len(maybePair)-1]
	}
	// permutate the result
	ret := make([][]cardgame.Card, 0)
	for _, pair := range maybePair {
		ret = append(ret, this.Combinate(pair, count)...)
	}
	return len(ret) > 0, ret
}

func (this *Poker) HasSerial(cards []cardgame.Card, length int) (bool, [][]cardgame.Card) {
	cards = this.prepareValues(cards)
	// pooling
	poolStraights := make([][][]cardgame.Card, 0)
	maybeStraight := [][]cardgame.Card{{cards[0]}}
	
	for i := 1; i < len(cards); i++ {
		card := cards[i]
		// add to last maybe
		if cards[i-1].Value == card.Value {
			j := len(maybeStraight)
			maybeStraight[j-1] = append(maybeStraight[j-1], card)
			continue
		}
		// add to pool when meet straight length
		if len(maybeStraight) >= length {
			tmpStraight := maybeStraight
			poolStraights = append(poolStraights, tmpStraight)
			maybeStraight = maybeStraight[1:]
		}
		// reset when meet a value gap
		if cards[i-1].Value > card.Value + 1 {
			maybeStraight = [][]cardgame.Card{{card}}
			continue
		}
		// grow maybe
		maybeStraight = append(maybeStraight, []cardgame.Card{card})
	}
	// last check
	if len(maybeStraight) >= length {
		tmpStraight := maybeStraight
		poolStraights = append(poolStraights, tmpStraight)
	}
	// permutate the result
	ret := make([][]cardgame.Card, 0)
	for _, poolStraight := range poolStraights {
		ret = append(ret, this.Permutate(poolStraight)...)
	}
	return len(ret) > 0, ret
}