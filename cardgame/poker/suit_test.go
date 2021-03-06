package poker_test

import (
	"github.com/zhang1career/golab/cardgame"
	"github.com/zhang1career/golab/cardgame/poker"
	"testing"
)

func TestPoker_HasSuit(t *testing.T) {
	var cards = []cardgame.Card{
		{Value: 3, Suit: poker.Heart},
		{Value: 2, Suit: poker.Spade},
		{Value: 1, Suit: poker.Spade},
		{Value: 2, Suit: poker.Diamond},
		{Value: 3, Suit: poker.Club},
		{Value: 3, Suit: poker.Diamond},
	}
	game, err := poker.New()
	if err != nil {
		t.Error(err.Error())
	}
	
	hasSuit, suits := game.HasSuit(cards, 0)
	t.Log(hasSuit)
	for _, suit := range suits {
		t.Log(game.Show(suit))
	}
	
	hasSuit, suits = game.HasSuit(cards, 1)
	t.Log(hasSuit)
	for _, suit := range suits {
		t.Log(game.Show(suit))
	}
	
	hasSuit, suits = game.HasSuit(cards, 2)
	t.Log(hasSuit)
	for _, suit := range suits {
		t.Log(game.Show(suit))
	}
	
	hasSuit, suits = game.HasSuit(cards, 3)
	t.Log(hasSuit)
	for _, suit := range suits {
		t.Log(game.Show(suit))
	}
	
	hasSuit, suits = game.HasSuit(cards, 6)
	t.Log(hasSuit)
	for _, suit := range suits {
		t.Log(game.Show(suit))
	}
	hasSuit, suits = game.HasSuit(cards, 7)
	t.Log(hasSuit)
	for _, suit := range suits {
		t.Log(game.Show(suit))
	}
	
	hasSuit, suits = game.HasSuit(cards, 999999999999999999)
	t.Log(hasSuit)
	for _, suit := range suits {
		t.Log(game.Show(suit))
	}
}