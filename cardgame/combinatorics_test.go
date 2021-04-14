package cardgame_test

import (
	"github.com/zhang1career/golab/cardgame"
	"testing"
)

func TestGame_Permutate(t *testing.T) {
	var card1 = []cardgame.Card{
		{Value: 1, Suit: 1},
		{Value: 1, Suit: 2},
		{Value: 1, Suit: 3},
		{Value: 1, Suit: 4},
	}
	
	var card2 = []cardgame.Card{
		{Value: 2, Suit: 1},
		{Value: 2, Suit: 2},
		{Value: 2, Suit: 3},
		{Value: 2, Suit: 4},
	}
	
	game, err := cardgame.New()
	if err != nil {
		t.Error(err.Error())
	}
	
	result := game.Permutate([][]cardgame.Card{card1, card2})
	t.Log(result)
}

func TestGame_Combinate(t *testing.T) {
	var cards = []cardgame.Card{
		{Value: 1, Suit: 1},
		{Value: 2, Suit: 1},
		{Value: 3, Suit: 1},
		{Value: 4, Suit: 1},
	}
	
	game, err := cardgame.New()
	if err != nil {
		t.Error(err.Error())
	}
	
	result := game.Combinate(cards, 2)
	t.Log(result)
}