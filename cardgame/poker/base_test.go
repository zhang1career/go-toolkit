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
	{Value: 10, Suit: poker.Diamond},
	{Value: 11, Suit: poker.Club},
	{Value: 5, Suit: poker.Spade},
	{Value: 12, Suit: poker.Club},
	{Value: 6, Suit: poker.Heart},
	{Value: 13, Suit: poker.Club},
}

func TestPokerGame_SortByValue(t *testing.T) {
	game, err := poker.New()
	if err != nil {
		t.Error(err.Error())
	}
	
	values := game.SortByValue(cards, "asc")
	t.Log(game.Show(values))
	
	values = game.SortByValue(cards, "desc")
	t.Log(game.Show(values))
}

func TestPokerGame_SortBySuit(t *testing.T) {
	game, err := poker.New()
	if err != nil {
		t.Error(err.Error())
	}
	
	values := game.SortBySuit(cards, "asc")
	t.Log(game.Show(values))
	
	values = game.SortBySuit(cards, "desc")
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
	
	hasStraight, straights := game.HasStraight(cards, 5)
	t.Log(hasStraight)
	for _, straight := range straights {
		t.Log(game.Show(straight))
	}
}

func TestPokerGame_HasSuit(t *testing.T) {
	game, err := poker.New()
	if err != nil {
		t.Error(err.Error())
	}
	
	hasSuit, suit := game.HasSuit(cards, 5)
	t.Log(hasSuit)
	t.Log(suit)
}

func TestPermutate(t *testing.T) {
	var card5 = []cardgame.Card{
		{Value: 5, Suit: poker.Diamond},
	}
	var card4 = []cardgame.Card{
		{Value: 4, Suit: poker.Diamond},
	}
	var card3 = []cardgame.Card{
		{Value: 3, Suit: poker.Diamond},
		{Value: 3, Suit: poker.Club},
	}
	var card2 = []cardgame.Card{
		{Value: 2, Suit: poker.Diamond},
		{Value: 2, Suit: poker.Club},
	}
	//var card1 = []cardgame.Card{
	//	{Value: 1, Suit: poker.Heart},
	//	{Value: 1, Suit: poker.Spade},
	//}
	var cards = [][]cardgame.Card{card5, card4, card3, card2}
	t.Log(cards)
	ret := poker.Permutate(cards)
	t.Log(ret)
}