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

func (this *Poker) GetMostCards(cards []cardgame.Card, count int) []cardgame.Card {
	if count <= 0 {
		return nil
	}
	if count >= len(cards) {
		return cards
	}
	cards = this.prepareValues(cards)
	return cards[0:count]
}

func (this *Poker) GetLeastCards(cards []cardgame.Card, count int) []cardgame.Card {
	if count <= 0 {
		return nil
	}
	if count >= len(cards) {
		return cards
	}
	cards = this.prepareValues(cards)
	return cards[len(cards)-count:]
}

func (this *Poker) CalcScore(cards []cardgame.Card, count int) (int, []cardgame.Card) {
	cards = this.GetMostCards(cards, count)
	score := 0
	for _, card := range cards {
		score += card.Value
	}
	return score, cards
}

func (this *Poker) HasPair(cards []cardgame.Card, count int) (bool, [][]cardgame.Card) {
	cards = this.prepareValues(cards)
	// pooling
	maybePair := [][]cardgame.Card{{cards[0]}}
	
	for i := 1; i < len(cards); i++ {
		card := cards[i]
		// add to last maybe
		if cards[i-1].Value == card.Value {
			maybePair[len(maybePair)-1] = append(maybePair[len(maybePair)-1], card)
			continue
		}
		// delete last maybe when not a pair
		if len(maybePair[len(maybePair)-1]) < count {
			maybePair = maybePair[0:len(maybePair)-1]
		}
		maybePair = append(maybePair, []cardgame.Card{card})
	}
	// last check
	if len(maybePair[len(maybePair)-1]) < count {
		maybePair = maybePair[0:len(maybePair)-1]
	}
	// permutate the result
	ret := make([][]cardgame.Card, 0)
	for _, pair := range maybePair {
		ret = append(ret, this.Combinate(pair, count)...)
	}
	return len(ret) > 0, ret
}

func (this *Poker) HasSerial(cards []cardgame.Card, length int) (bool, [][]cardgame.Card) {
	cards = this.prepareValues(cards)
	// pooling
	poolStraights := make([][][]cardgame.Card, 0)
	maybeStraight := [][]cardgame.Card{{cards[0]}}
	
	for i := 1; i < len(cards); i++ {
		card := cards[i]
		// add to last maybe
		if cards[i-1].Value == card.Value {
			j := len(maybeStraight)
			maybeStraight[j-1] = append(maybeStraight[j-1], card)
			continue
		}
		// add to pool when meet straight length
		if len(maybeStraight) >= length {
			tmpStraight := make([][]cardgame.Card, len(maybeStraight))
			copy(tmpStraight, maybeStraight)
			poolStraights = append(poolStraights, tmpStraight)
			maybeStraight = maybeStraight[1:]
		}
		// reset when meet a value gap
		if cards[i-1].Value > card.Value + 1 {
			maybeStraight = [][]cardgame.Card{{card}}
			continue
		}
		// grow maybe
		maybeStraight = append(maybeStraight, []cardgame.Card{card})
	}
	// last check
	if len(maybeStraight) >= length {
		tmpStraight := make([][]cardgame.Card, len(maybeStraight))
		copy(tmpStraight, maybeStraight)
		poolStraights = append(poolStraights, tmpStraight)
	}
	// permutate the result
	ret := make([][]cardgame.Card, 0)
	for _, poolStraight := range poolStraights {
		ret = append(ret, this.Permutate(poolStraight)...)
	}
	return len(ret) > 0, ret
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

func (this *Poker) HasSuit(cards []cardgame.Card, count int) (bool, [][]cardgame.Card) {
	cards = this.prepareSuits(cards)
	// pooling
	maybeSuit := [][]cardgame.Card{{cards[0]}}
	
	for i := 1; i < len(cards); i++ {
		card := cards[i]
		// add to last maybe
		if cards[i-1].Suit == card.Suit {
			maybeSuit[len(maybeSuit)-1] = append(maybeSuit[len(maybeSuit)-1], card)
			continue
		}
		// delete last maybe when not enough the same suit
		if len(maybeSuit[len(maybeSuit)-1]) < count {
			maybeSuit = maybeSuit[0:len(maybeSuit)-1]
		}
		maybeSuit = append(maybeSuit, []cardgame.Card{card})
	}
	// last check
	if len(maybeSuit[len(maybeSuit)-1]) < count {
		maybeSuit = maybeSuit[0:len(maybeSuit)-1]
	}
	return len(maybeSuit) > 0, maybeSuit
}

func (this *Poker) prepareSuits(cards []cardgame.Card) []cardgame.Card {
	cards = this.SortBySuit(cards, "desc")
	return cards
}