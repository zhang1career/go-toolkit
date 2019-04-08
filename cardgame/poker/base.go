package poker

import (
	"fmt"
	"github.com/zhang1career/lib/cardgame"
)

const (
	Spade   = 4     // ♠
	Heart   = 3     // ♥
	Club    = 2     // ♣
	Diamond = 1     // ♦
)

type PokerGame struct {
	cardgame.Game
}

func New() (*PokerGame, error) {
	game := PokerGame{
		cardgame.Game{
			SuitMap: map[int]string{
				Spade:   "♠",
				Heart:   "♥",
				Club:    "♣",
				Diamond: "♦",
			},
		},
	}
	return &game, nil
}

func (this *PokerGame) HasPair(cards []cardgame.KV, count int) (int, []cardgame.KV) {
	ret := make([]cardgame.KV, 0)
	
	for _, kv := range cards {
		if kv.Value < count {
			continue
		}
		ret = append(ret, kv)
	}
	
	return len(ret), ret
}

func (this *PokerGame) HasStraight(cards []cardgame.Card, length int) (bool, [][]cardgame.Card) {
	cards = appendHighAce(cards)
	cards = this.SortByValue(cards, "desc")
	// pooling
	poolStraights := make([][][]cardgame.Card, 0)
	maybeStraight := [][]cardgame.Card{{cards[0]}}
	
	for i := 1; i < len(cards); i++ {
		card := cards[i]
		// add to last maybe
		j := len(maybeStraight)
		if cards[i-1].Value == card.Value {
			maybeStraight[j-1] = append(maybeStraight[j-1], card)
			continue
		}
		// add to pool when meet straight length
		if len(maybeStraight) >= length {
			tmpStraight := make([][]cardgame.Card, len(maybeStraight))
			copy(tmpStraight, maybeStraight)
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
		tmpStraight := make([][]cardgame.Card, len(maybeStraight))
		copy(tmpStraight, maybeStraight)
		poolStraights = append(poolStraights, tmpStraight)
	}
	// permutate the result
	ret := make([][]cardgame.Card, 0)
	for _, poolStraight := range poolStraights {
		fmt.Println(poolStraight)
		s := Permutate(poolStraight)
		ret = append(ret, s...)
	}
	return len(ret) > 0, ret
}

func appendHighAce(cards []cardgame.Card) []cardgame.Card {
	// do not append twice
	for _, card := range cards {
		if card.Value >= 14 {
			return cards
		}
	}
	
	highAces := make([]cardgame.Card, 0)
	for _, card := range cards {
		if card.Value != 1 {
			continue
		}
		highAces = append(highAces, cardgame.Card{Value: 14, Suit: card.Suit})
	}
	return append(cards, highAces...)
}

func Permutate(cards [][]cardgame.Card) [][]cardgame.Card {
	ret := make([][]cardgame.Card, 0)
	
	length := len(cards)
	if length <= 1 {
		for _, card := range cards[length-1] {
			ret = append(ret, []cardgame.Card{card})
		}
		return ret
	}
	
	for _, straight := range Permutate(cards[0:length-1]) {
		for _, head := range cards[length-1] {
			tmpStraight := make([]cardgame.Card, len(straight))
			copy(tmpStraight, straight)
			tmpStraight = append(tmpStraight, head)
			ret = append(ret, tmpStraight)
		}
	}
	
	return ret
}

func (this *PokerGame) HasSuit(cards []cardgame.Card, count int) (int, []int) {
	ret := make([]int, 0)
	
	kvs := this.GroupBySuit(cards)
	for k, v := range kvs {
		if v < count {
			continue
		}
		ret = append(ret, k)
	}
	return len(ret), ret
}