package poker_test

import (
	"github.com/zhang1career/lib/cardgame"
	"github.com/zhang1career/lib/cardgame/poker"
	"testing"
)

func TestPoker_HasPair(t *testing.T) {
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
	
	game, err := poker.New()
	if err != nil {
		t.Error(err.Error())
	}
	
	hasPair, pairs := game.HasPair(cards, 2)
	t.Log(hasPair)
	for _, pair := range pairs {
		t.Log(game.Show(pair))
	}
}

func TestPoker_HasSerial(t *testing.T) {
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
	
	game, err := poker.New()
	if err != nil {
		t.Error(err.Error())
	}
	
	hasStraight, straights := game.HasSerial(cards, 5)
	t.Log(hasStraight)
	for _, straight := range straights {
		t.Log(game.Show(straight))
	}
}