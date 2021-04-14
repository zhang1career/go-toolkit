package texas_test

import (
	"github.com/zhang1career/golab/cardgame"
	"github.com/zhang1career/golab/cardgame/poker"
	"github.com/zhang1career/golab/cardgame/poker/texas"
	"testing"
)

func TestTexas_HasOnePair(t *testing.T) {
	game, err := texas.New()
	if err != nil {
		t.Error(err.Error())
	}
	
	var cards = []cardgame.Card{
		{Value: 2, Suit: poker.Spade},
		{Value: 1, Suit: poker.Club},
		{Value: 2, Suit: poker.Diamond},
		{Value: 3, Suit: poker.Club},
		{Value: 3, Suit: poker.Heart},
	}
	
	has, value := game.HasOnePair(cards)
	t.Log(has)
	t.Log(game.Show(value))
}

func TestTexas_HasTwoPair(t *testing.T) {
	game, err := texas.New()
	if err != nil {
		t.Error(err.Error())
	}
	
	var cards = []cardgame.Card{
		{Value: 2, Suit: poker.Spade},
		{Value: 1, Suit: poker.Club},
		{Value: 2, Suit: poker.Diamond},
		{Value: 3, Suit: poker.Club},
		{Value: 3, Suit: poker.Heart},
	}
	
	has, values := game.HasTwoPair(cards)
	t.Log(has)
	for _, value := range values {
		t.Log(game.Show(value))
	}
}

func TestTexas_HasThreeOfAKind(t *testing.T) {
	game, err := texas.New()
	if err != nil {
		t.Error(err.Error())
	}
	
	var cards = []cardgame.Card{
		{Value: 3, Suit: poker.Heart},
		{Value: 2, Suit: poker.Spade},
		{Value: 1, Suit: poker.Club},
		{Value: 2, Suit: poker.Diamond},
		{Value: 3, Suit: poker.Club},
		{Value: 3, Suit: poker.Diamond},
	}
	
	has, value := game.HasThreeOfAKind(cards)
	t.Log(has)
	t.Log(game.Show(value))
}

func TestTexas_HasFourOfAKind(t *testing.T) {
	game, err := texas.New()
	if err != nil {
		t.Error(err.Error())
	}
	
	var cards = []cardgame.Card{
		{Value: 12, Suit: poker.Diamond},
		{Value: 3, Suit: poker.Heart},
		{Value: 2, Suit: poker.Spade},
		{Value: 1, Suit: poker.Club},
		{Value: 2, Suit: poker.Diamond},
		{Value: 3, Suit: poker.Club},
		{Value: 12, Suit: poker.Club},
		{Value: 3, Suit: poker.Diamond},
		{Value: 11, Suit: poker.Club},
		{Value: 12, Suit: poker.Spade},
		{Value: 13, Suit: poker.Club},
		{Value: 12, Suit: poker.Heart},
	}
	
	has, value := game.HasFourOfAKind(cards)
	t.Log(has)
	t.Log(game.Show(value))
}

func TestTexas_HasFullHouse(t *testing.T) {
	game, err := texas.New()
	if err != nil {
		t.Error(err.Error())
	}
	
	var cards = []cardgame.Card{
		{Value: 3, Suit: poker.Heart},
		{Value: 2, Suit: poker.Spade},
		{Value: 1, Suit: poker.Club},
		{Value: 2, Suit: poker.Diamond},
		{Value: 3, Suit: poker.Club},
		{Value: 12, Suit: poker.Club},
		{Value: 3, Suit: poker.Diamond},
	}
	
	has, values := game.HasFullHouse(cards)
	t.Log(has)
	for _, value := range values {
		t.Log(game.Show(value))
	}
}

func TestTexas_HasFlush(t *testing.T) {
	game, err := texas.New()
	if err != nil {
		t.Error(err.Error())
	}
	
	var cards = []cardgame.Card{
		{Value: 12, Suit: poker.Diamond},
		{Value: 3, Suit: poker.Heart},
		{Value: 2, Suit: poker.Spade},
		{Value: 1, Suit: poker.Club},
		{Value: 2, Suit: poker.Diamond},
		{Value: 3, Suit: poker.Club},
		{Value: 12, Suit: poker.Club},
		{Value: 3, Suit: poker.Diamond},
		{Value: 10, Suit: poker.Club},
		{Value: 4, Suit: poker.Club},
		{Value: 12, Suit: poker.Spade},
		{Value: 13, Suit: poker.Club},
		{Value: 5, Suit: poker.Heart},
	}
	
	has, value := game.HasFlush(cards)
	t.Log(has)
	t.Log(game.Show(value))
}

func TestTexas_HasStraight(t *testing.T) {
	game, err := texas.New()
	if err != nil {
		t.Error(err.Error())
	}
	
	var cards = []cardgame.Card{
		{Value: 12, Suit: poker.Diamond},
		{Value: 3, Suit: poker.Heart},
		{Value: 2, Suit: poker.Spade},
		{Value: 1, Suit: poker.Club},
		{Value: 2, Suit: poker.Diamond},
		{Value: 3, Suit: poker.Club},
		{Value: 12, Suit: poker.Club},
		{Value: 3, Suit: poker.Diamond},
		{Value: 8, Suit: poker.Spade},
		{Value: 4, Suit: poker.Club},
		{Value: 10, Suit: poker.Club},
		{Value: 11, Suit: poker.Club},
		{Value: 12, Suit: poker.Spade},
		{Value: 13, Suit: poker.Club},
		{Value: 11, Suit: poker.Heart},
	}
	
	has, values := game.HasStraight(cards)
	t.Log(has)
	for _, value := range values {
		t.Log(game.Show(value))
	}
}

func TestTexas_HasStraightFlush(t *testing.T) {
	game, err := texas.New()
	if err != nil {
		t.Error(err.Error())
	}
	
	var cards = []cardgame.Card{
		{Value: 12, Suit: poker.Diamond},
		{Value: 3, Suit: poker.Heart},
		{Value: 2, Suit: poker.Spade},
		{Value: 1, Suit: poker.Club},
		{Value: 2, Suit: poker.Diamond},
		{Value: 3, Suit: poker.Club},
		{Value: 12, Suit: poker.Club},
		{Value: 3, Suit: poker.Diamond},
		{Value: 8, Suit: poker.Spade},
		{Value: 4, Suit: poker.Club},
		{Value: 10, Suit: poker.Club},
		{Value: 11, Suit: poker.Club},
		{Value: 12, Suit: poker.Spade},
		{Value: 13, Suit: poker.Club},
		{Value: 11, Suit: poker.Heart},
	}
	
	has, value := game.HasStraightFlush(cards)
	t.Log(has)
	t.Log(game.Show(value))
}

func TestTexas_HasRoyalStraightFlush(t *testing.T) {
	game, err := texas.New()
	if err != nil {
		t.Error(err.Error())
	}
	
	var cards = []cardgame.Card{
		{Value: 12, Suit: poker.Diamond},
		{Value: 3, Suit: poker.Heart},
		{Value: 2, Suit: poker.Spade},
		{Value: 1, Suit: poker.Club},
		{Value: 2, Suit: poker.Diamond},
		{Value: 3, Suit: poker.Club},
		{Value: 12, Suit: poker.Club},
		{Value: 3, Suit: poker.Diamond},
		{Value: 8, Suit: poker.Spade},
		{Value: 4, Suit: poker.Club},
		{Value: 10, Suit: poker.Club},
		{Value: 11, Suit: poker.Club},
		{Value: 12, Suit: poker.Spade},
		{Value: 13, Suit: poker.Club},
		{Value: 11, Suit: poker.Heart},
	}
	
	has, value := game.HasRoyalStraightFlush(cards)
	t.Log(has)
	t.Log(game.Show(value))
}