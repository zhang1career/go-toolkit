package cardgame_test

import (
	"github.com/zhang1career/lib/cardgame"
	"testing"
)

var cards = []cardgame.Card{
	{Value: 2, Suit: 2},
	{Value: 3, Suit: 4},
	{Value: 1, Suit: 3},
	{Value: 2, Suit: 3},
	{Value: 3, Suit: 2},
	{Value: 3, Suit: 1},
}

func TestSortByValue(t *testing.T) {
	game, err := cardgame.New()
	if err != nil {
		t.Error(err.Error())
	}
	values, err := game.SortByValue(cards, "asc")
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(game.Show(values))
	
	values, err = game.SortByValue(cards, "desc")
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(game.Show(values))
}

func TestSortBySuit(t *testing.T) {
	game, err := cardgame.New()
	if err != nil {
		t.Error(err.Error())
	}
	values, err := game.SortBySuit(cards, "asc")
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(game.Show(values))
	
	values, err = game.SortBySuit(cards, "desc")
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(game.Show(values))
}

func TestGroupByValue(t *testing.T) {
	game, err := cardgame.New()
	if err != nil {
		t.Error(err.Error())
	}
	values := game.GroupByValue(cards)
	t.Log(values)
}

func TestGroupByValueAndSort(t *testing.T) {
	game, err := cardgame.New()
	if err != nil {
		t.Error(err.Error())
	}
	values := game.GroupByValue(cards).Sort("asc")
	t.Log(values)
	values = game.GroupByValue(cards).Sort("desc")
	t.Log(values)
}

func TestGroupBySuit(t *testing.T) {
	game, err := cardgame.New()
	if err != nil {
		t.Error(err.Error())
	}
	values := game.GroupBySuit(cards)
	if err != nil {
		t.Error(err.Error())
	}
	t.Log(values)
}

func TestGroupBySuitAndSort(t *testing.T) {
	game, err := cardgame.New()
	if err != nil {
		t.Error(err.Error())
	}
	values := game.GroupBySuit(cards).Sort("asc")
	t.Log(values)
	values = game.GroupBySuit(cards).Sort("desc")
	t.Log(values)
}