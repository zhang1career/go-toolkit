package poker

import (
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

func (this *PokerGame) HasStraight(cards []cardgame.KV, length int) (bool, int) {
	cards = prependHighAce(cards)
	
	straight_num := 0
	start_value  := 0
	
	for _, kv := range cards {
		// params check
		if kv.Value <= 0 {
			straight_num = 0
			start_value  = 0
			continue
		}
		// reset when meet a number gap
		if start_value != kv.Key + 1 {
			straight_num = 0
		}
		straight_num++
		start_value = kv.Key
		// return when meet straight length
		if (straight_num >= length) {
			break
		}
	}
	
	return (straight_num >= length), start_value
}

func prependHighAce(cards []cardgame.KV) []cardgame.KV {
	// do not prepend twice
	if cards[0].Key >= 14 {
		return cards
	}
	
	lowAce := cards[len(cards)-1]
	// there is no ace
	if lowAce.Key != 1 {
		return cards
	}
	
	highAce := cardgame.KV{
		Key:    14,
		Value:  lowAce.Value,
	}
	return append([]cardgame.KV{highAce}, cards...)
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