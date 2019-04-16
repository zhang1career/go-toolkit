package texas_test

import (
	"github.com/zhang1career/lib/cardgame"
	"github.com/zhang1career/lib/cardgame/poker"
	"github.com/zhang1career/lib/cardgame/poker/texas"
	"testing"
)

//func TestGetResult(t *testing.T) {
//	var card0 = []cardgame.Card{
//		{Value: 10, Suit: poker.Diamond},
//		{Value: 11, Suit: poker.Diamond},
//		{Value: 12, Suit: poker.Diamond},
//		{Value: 13, Suit: poker.Diamond},
//		{Value: 13, Suit: poker.Club},
//	}
//	var card1 = []cardgame.Card{
//		{Value: 14, Suit: poker.Diamond},
//		{Value: 14, Suit: poker.Club},
//	}
//	var card2 = []cardgame.Card{
//		{Value: 4, Suit: poker.Heart},
//		{Value: 10, Suit: poker.Spade},
//	}
//	results := texas.GetResult(card0, [][]cardgame.Card{card1, card2})
//	t.Log(results)
//}

func TestGetStat(t *testing.T) {
	var card0 = []cardgame.Card{
		{Value: 10, Suit: poker.Diamond},
		{Value: 11, Suit: poker.Diamond},
		{Value: 12, Suit: poker.Diamond},
	}
	var card1 = []cardgame.Card{
		{Value: 14, Suit: poker.Diamond},
		{Value: 2,  Suit: poker.Club},
	}
	
	results := texas.GetStat(card0, card1)
	t.Log(results)
}