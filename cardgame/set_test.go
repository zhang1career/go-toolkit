package cardgame_test

import (
	"github.com/zhang1career/golab/cardgame"
	"testing"
)

func TestGame_AddCards(t *testing.T) {
	var card1 = []cardgame.Card{
		{Value: 1, Suit: 1},
		{Value: 1, Suit: 2},
		{Value: 2, Suit: 3},
		{Value: 2, Suit: 4},
	}
	
	var card2 = []cardgame.Card{
		{Value: 1, Suit: 1},
		{Value: 1, Suit: 3},
		{Value: 2, Suit: 1},
		{Value: 2, Suit: 4},
	}
	
	game, err := cardgame.New()
	if err != nil {
		t.Error(err.Error())
	}
	
	result := game.AddCards(card1, card2)
	t.Log(result)
}

func TestGame_DelCards(t *testing.T) {
	var card1 = []cardgame.Card{
		{Value: 1, Suit: 1},
		{Value: 1, Suit: 2},
		{Value: 2, Suit: 3},
		{Value: 2, Suit: 4},
	}
	
	var card2 = []cardgame.Card{
		{Value: 1, Suit: 1},
		{Value: 1, Suit: 3},
		{Value: 2, Suit: 1},
		{Value: 2, Suit: 4},
	}
	
	game, err := cardgame.New()
	if err != nil {
		t.Error(err.Error())
	}
	
	result := game.DelCards(card1, card2)
	t.Log(result)
}

func TestGame_DealCards(t *testing.T) {
	var card = []cardgame.Card{
		{Value: 1, Suit: 1},
		{Value: 1, Suit: 2},
		{Value: 2, Suit: 3},
		{Value: 2, Suit: 4},
	}
	
	game, err := cardgame.New()
	if err != nil {
		t.Error(err.Error())
	}
	
	result, card := game.DealCards(card, 3)
	t.Log(result)
	t.Log(card)
}