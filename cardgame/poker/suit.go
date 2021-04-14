package poker


import (
	"github.com/zhang1career/golab/cardgame"
)

func (this *Poker) HasSuit(cards []cardgame.Card, count int) (bool, [][]cardgame.Card) {
	// params checkage
	if count <= 0 {
		return false, nil
	}
	
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