package poker

import (
	"github.com/zhang1career/lib/cardgame"
)

const (
	Spade   = 4     // ♠
	Heart   = 3     // ♥
	Club    = 2     // ♣
	Diamond = 1     // ♦
)

type Poker struct {
	cardgame.Game
}

func New() (*Poker, error) {
	game := Poker{
		cardgame.Game{
			ValueMap: map[int]string{
				1:      "A",
				2:      "2",
				3:      "3",
				4:      "4",
				5:      "5",
				6:      "6",
				7:      "7",
				8:      "8",
				9:      "9",
				10:     "10",
				11:     "J",
				12:     "Q",
				13:     "K",
				14:     "A",
			},
			SuitMap: map[int]string{
				Spade:  "♠",
				Heart:  "♥",
				Club:   "♣",
				Diamond:"♦",
			},
		},
	}
	return &game, nil
}

func (this *Poker) GetAllCards() []cardgame.Card {
	return []cardgame.Card {
		{1,1}, {1,2}, {1,3}, {1,4},
		{2,1}, {2,2}, {2,3}, {2,4},
		{3,1}, {3,2}, {3,3}, {3,4},
		{4,1}, {4,2}, {4,3}, {4,4},
		{5,1}, {5,2}, {5,3}, {5,4},
		{6,1}, {6,2}, {6,3}, {6,4},
		{7,1}, {7,2}, {7,3}, {7,4},
		{8,1}, {8,2}, {8,3}, {8,4},
		{9,1}, {9,2}, {9,3}, {9,4},
		{10,1}, {10,2}, {10,3}, {10,4},
		{11,1}, {11,2}, {11,3}, {11,4},
		{12,1}, {12,2}, {12,3}, {12,4},
		{13,1}, {13,2}, {13,3}, {13,4},
	}
}

func (this *Poker) prepareValues(cards []cardgame.Card) []cardgame.Card {
	cards = this.addHighAce(cards)
	cards = this.SortByValue(cards, "desc")
	return cards
}

func (this *Poker) addHighAce(cards []cardgame.Card) []cardgame.Card {
	// do not append twice
	for _, card := range cards {
		if card.Value >= 14 {
			return cards
		}
	}
	
	highAces := make([]cardgame.Card, 0)
	for _, card := range cards {
		if card.Value != 1 {
			continue
		}
		highAces = append(highAces, cardgame.Card{Value: 14, Suit: card.Suit})
	}
	return append(cards, highAces...)
}