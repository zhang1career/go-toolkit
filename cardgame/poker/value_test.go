package poker_test

import (
	"github.com/zhang1career/lib/cardgame"
	"github.com/zhang1career/lib/cardgame/poker"
	"testing"
)

func TestPoker_GetMostCards(t *testing.T) {
	cards := []cardgame.Card{
		{Value: 1, Suit: poker.Heart},
		{Value: 1, Suit: poker.Diamond},
		{Value: 2, Suit: poker.Spade},
		{Value: 3, Suit: poker.Diamond},
		{Value: 3, Suit: poker.Club},
	}
	
	game, err := poker.New()
	if err != nil {
		t.Error(err.Error())
	}
	
	m1, ok := game.GetMostCards(cards, 1)
	t.Log(m1)
	t.Log(ok)
	
	m2, ok := game.GetMostCards(cards, 2)
	t.Log(m2)
	t.Log(ok)
	
	m3, ok := game.GetMostCards(cards, 3)
	t.Log(m3)
	t.Log(ok)
	
	m4, ok := game.GetMostCards(cards, 4)
	t.Log(m4)
	t.Log(ok)
	
	m5, ok := game.GetLeastCards(cards, 5)
	t.Log(m5)
	t.Log(ok)
	
	m6, ok := game.GetLeastCards(cards, 6)
	t.Log(m6)
	t.Log(ok)
	
	m7, ok := game.GetLeastCards(cards, 7)
	t.Log(m7)
	t.Log(ok)
}

func TestPoker_GetLeastCards(t *testing.T) {
	cards := []cardgame.Card{
		{Value: 1, Suit: poker.Heart},
		{Value: 1, Suit: poker.Diamond},
		{Value: 2, Suit: poker.Spade},
		{Value: 3, Suit: poker.Diamond},
		{Value: 3, Suit: poker.Club},
	}
	
	game, err := poker.New()
	if err != nil {
		t.Error(err.Error())
	}
	
	l1, ok := game.GetLeastCards(cards, 1)
	t.Log(l1)
	t.Log(ok)
	
	l2, ok := game.GetLeastCards(cards, 2)
	t.Log(l2)
	t.Log(ok)
	
	l3, ok := game.GetLeastCards(cards, 3)
	t.Log(l3)
	t.Log(ok)
	
	l4, ok := game.GetLeastCards(cards, 4)
	t.Log(l4)
	t.Log(ok)
	
	l5, ok := game.GetLeastCards(cards, 5)
	t.Log(l5)
	t.Log(ok)
	
	l6, ok := game.GetLeastCards(cards, 6)
	t.Log(l6)
	t.Log(ok)
	
	l7, ok := game.GetLeastCards(cards, 7)
	t.Log(l7)
	t.Log(ok)
}

func TestPoker_CalcScore(t *testing.T) {
	cards := []cardgame.Card{
		{Value: 1, Suit: poker.Heart},
		{Value: 1, Suit: poker.Diamond},
		{Value: 2, Suit: poker.Spade},
		{Value: 3, Suit: poker.Diamond},
		{Value: 3, Suit: poker.Club},
	}
	
	game, err := poker.New()
	if err != nil {
		t.Error(err.Error())
	}
	
	s1, cards, ok := game.CalcScore(cards, 4)
	t.Log(s1)
	t.Log(cards)
	t.Log(ok)
	
}