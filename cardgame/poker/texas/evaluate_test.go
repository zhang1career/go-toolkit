package texas_test

import (
	"github.com/zhang1career/lib/cardgame"
	"github.com/zhang1career/lib/cardgame/poker"
	"github.com/zhang1career/lib/cardgame/poker/texas"
	"testing"
)

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

func TestTexas_HasOnePair(t *testing.T) {
	//game, err := poker.New()
	//if err != nil {
	//	t.Error(err.Error())
	//}
	//
	//kvs := game.GroupByValue(cards).Sort("desc")
	//t.Log(kvs)
	
	has, value := texas.HasOnePair(cards)
	t.Log(has)
	t.Log(value)
}

func TestTexas_HasFlush(t *testing.T) {
	//game, err := poker.New()
	//if err != nil {
	//	t.Error(err.Error())
	//}
	//
	//kvs := game.GroupByValue(cards).Sort("desc")
	//t.Log(kvs)
	
	has, value := texas.HasFlush(cards)
	t.Log(has)
	t.Log(value)
}

func TestTexas_HasStraightFlush(t *testing.T) {
	game, err := poker.New()
	if err != nil {
		t.Error(err.Error())
	}
	
	has, value := texas.HasStraightFlush(cards)
	t.Log(has)
	t.Log(game.Show(value))
}

func TestTexas_HasRoyalStraightFlush(t *testing.T) {
	game, err := poker.New()
	if err != nil {
		t.Error(err.Error())
	}
	
	has, value := texas.HasRoyalStraightFlush(cards)
	t.Log(has)
	t.Log(game.Show(value))
}