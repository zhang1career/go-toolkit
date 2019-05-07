package cardgame_test

import (
	"fmt"
	"github.com/zhang1career/lib/cardgame"
	"testing"
)

var allCards = []cardgame.Card{
	{Value: 0, Suit: 0},
	{Value: 1, Suit: 0},
	{Value: 2, Suit: 0},
	{Value: 3, Suit: 0},
	{Value: 4, Suit: 0},
	{Value: 5, Suit: 0},
	{Value: 6, Suit: 0},
	{Value: 7, Suit: 0},
	{Value: 8, Suit: 0},
	{Value: 9, Suit: 0},
}

func TestGame_Shuffle(t *testing.T) {
	game, err := cardgame.New()
	if err != nil {
		t.Error(err.Error())
	}
	
	t.Log(allCards)
	shuffledCards := game.Shuffle(allCards, 0)
	t.Log(shuffledCards)
}

func TestGame_Shuffle_Distribution(t *testing.T) {
	game, err := cardgame.New()
	if err != nil {
		t.Error(err.Error())
	}
	
	count := make([][]int, 10)
	for k := 0; k < 10; k++ {
		count[k] = make([]int, 10)
	}
	
	for t := 0; t < 100000; t++ {
		// create all cards
		allCards := make([]cardgame.Card, 10)
		for i := 0; i < 10; i++ {
			allCards[i] = cardgame.Card{Value:i, Suit:0}
		}
		// shuffle
		shuffledCards := game.Shuffle(allCards, t)
		// count
		for j := 0; j < 10; j++ {
			card := shuffledCards[j]
			count[j][card.Value] = count[j][card.Value] + 1
		}
	}
	
	fmt.Println(count)
}