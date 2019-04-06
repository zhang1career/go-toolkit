package poker_test

import (
	"github.com/zhang1career/lib/cardgame"
	"github.com/zhang1career/lib/cardgame/poker"
	"testing"
)

var cards = []cardgame.Card{
	{Value: 3, Suit: poker.Heart},
	{Value: 2, Suit: poker.Spade},
	{Value: 1, Suit: poker.Spade},
	{Value: 2, Suit: poker.Diamond},
	{Value: 3, Suit: poker.Club},
	{Value: 3, Suit: poker.Diamond},
	{Value: 8, Suit: poker.Spade},
	{Value: 4, Suit: poker.Club},
	{Value: 10, Suit: poker.Club},
	{Value: 11, Suit: poker.Club},
	{Value: 12, Suit: poker.Club},
	{Value: 13, Suit: poker.Club},
}

func TestPokerGame_SortByValue(t *testing.T) {
	game, err := poker.New()
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

func TestPokerGame_SortBySuit(t *testing.T) {
	game, err := poker.New()
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

func TestPokerGame_HasPair(t *testing.T) {
	game, err := poker.New()
	if err != nil {
		t.Error(err.Error())
	}
	
	kvs := game.GroupByValue(cards).Sort("desc")
	t.Log(kvs)
	
	hasPair, pairs := game.HasPair(kvs, 2)
	t.Log(hasPair)
	t.Log(pairs)
}

func TestPokerGame_HasStraight(t *testing.T) {
	game, err := poker.New()
	if err != nil {
		t.Error(err.Error())
	}
	
	kvs := game.GroupByValue(cards).Sort("desc")
	t.Log(kvs)
	
	hasStraight, startValue := game.HasStraight(kvs, 5)
	t.Log(hasStraight)
	t.Log(startValue)
}

func TestPokerGame_HasSuit(t *testing.T) {
	game, err := poker.New()
	if err != nil {
		t.Error(err.Error())
	}
	
	num, suit := game.HasSuit(cards, 5)
	t.Log(num)
	t.Log(suit)
}