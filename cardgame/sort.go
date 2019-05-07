package cardgame

import (
	"sort"
)

func (this *Game) SortByValue(cards []Card, sort_type string) []Card {
	if sort_type == "asc" {
		sort.Slice(cards, func(i, j int) bool {
			return cards[i].Value < cards[j].Value
		})
	} else {
		sort.Slice(cards, func(i, j int) bool {
			return cards[i].Value > cards[j].Value
		})
	}
	return cards
}

func (this *Game) SortBySuit(cards []Card, sort_type string) []Card {
	if sort_type == "asc" {
		sort.Slice(cards, func(i, j int) bool {
			return cards[i].Suit < cards[j].Suit
		})
	} else {
		sort.Slice(cards, func(i, j int) bool {
			return cards[i].Suit > cards[j].Suit
		})
	}
	return cards
}


type ArrayCount struct {
	Key     int
	Value   int
}

type ValueCount map[int]int

func (this *Game) GroupByValue(cards []Card) ValueCount {
	rets := make(map[int]int, 0)
	for _, card := range cards {
		rets[card.Value] = rets[card.Value] + 1
	}
	return rets
}

func (this ValueCount) Sort(sort_type string) []ArrayCount {
	return sortMap(this, sort_type)
}


type SuitCount map[int]int

func (this *Game) GroupBySuit(cards []Card) SuitCount {
	rets := make(map[int]int, 0)
	for _, card := range cards {
		rets[card.Suit] = rets[card.Suit] + 1
	}
	return rets
}

func (this SuitCount) Sort(sort_type string) []ArrayCount {
	return sortMap(this, sort_type)
}


func sortMap(m map[int]int, sort_type string) []ArrayCount {
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	
	ret := make([]ArrayCount, len(m))
	if sort_type == "asc" {
		for i, k := range keys {
			ret[i] = ArrayCount{Key: k, Value: m[k]}
		}
	} else {
		for i := 0; i < len(keys); i++ {
			k := keys[len(keys)-1-i]
			ret[i] = ArrayCount{Key: k, Value: m[k]}
		}
	}
	return ret
}