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

func (this *Poker) CalcSccore(cards []cardgame.Card, count int) (int, []cardgame.Card) {
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
		ret = append(ret, this.combinate(pair, count)...)
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
		ret = append(ret, this.permutate(poolStraight)...)
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

func (this *Poker) permutate(cards [][]cardgame.Card) [][]cardgame.Card {
	ret := make([][]cardgame.Card, 0)
	
	length := len(cards)
	if length <= 1 {
		for _, card := range cards[length-1] {
			ret = append(ret, []cardgame.Card{card})
		}
		return ret
	}
	
	for _, straight := range this.permutate(cards[1:]) {
		for _, head := range cards[0] {
			tmpStraight := make([]cardgame.Card, len(straight))
			copy(tmpStraight, straight)
			tmpStraight = append(tmpStraight, head)
			ret = append(ret, tmpStraight)
		}
	}
	
	return ret
}

func (this *Poker) combinate(cards []cardgame.Card, count int) [][]cardgame.Card {
	if len(cards) < count || count <= 0 {
		return nil
	}
	
	ret := make([][]cardgame.Card, 0)
	if count == 1 {
		for _, card := range cards {
			ret = append(ret, []cardgame.Card{card})
		}
		return ret
	}
	
	for i := 0; i <= len(cards) - count; i++ {
		for _, com := range this.combinate(cards[i+1:], count-1) {
			tmpCom := make([]cardgame.Card, len(com))
			copy(tmpCom, com)
			tmpCom = append(tmpCom, cards[i])
			ret = append(ret, tmpCom)
		}
	}
	return ret
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