package cardgame

import (
	"fmt"
	"sort"
	"strconv"
)

type Card struct {
	Value       int
	Suit        int
}

type Game struct {
	SuitMap     map[int]string
}

func New() (*Game, error) {
	return &Game{}, nil
}

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


type ValueCount map[int]int

func (this *Game) GroupByValue(cards []Card) ValueCount {
	rets := make(map[int]int, 0)
	for _, card := range cards {
		rets[card.Value] = rets[card.Value] + 1
	}
	return rets
}

func (this ValueCount) Sort(sort_type string) []KV {
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

func (this SuitCount) Sort(sort_type string) []KV {
	return sortMap(this, sort_type)
}


func sortMap(m map[int]int, sort_type string) []KV {
	var keys []int
	for k := range m {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	
	ret := make([]KV, len(m))
	if sort_type == "asc" {
		fmt.Println(sort_type)
		for i, k := range keys {
			ret[i] = KV{Key: k, Value: m[k]}
		}
	} else {
		length := len(keys)-1
		for i := 0; i <= length; i++ {
			k := keys[length-i]
			ret[i] = KV{Key: k, Value: m[k]}
		}
	}
	return ret
}

type KV struct {
	Key     int
	Value   int
}


func (this *Game) Show(cards []Card) []string {
	ret := make([]string, len(cards))
	for i, card := range cards {
		ret[i] = strconv.Itoa(card.Value) + this.SuitMap[card.Suit]
	}
	return ret
}