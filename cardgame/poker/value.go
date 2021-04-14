package poker

import (
	"fmt"
	"github.com/zhang1career/golab/cardgame"
)

func (this *Poker) GetMostCards(cards []cardgame.Card, count int) ([]cardgame.Card, error) {
	if count <= 0 || count > len(cards) {
		return nil, fmt.Errorf("param out of range: count")
	}
	cards = this.prepareValues(cards)
	return cards[0:count], nil
}

func (this *Poker) GetLeastCards(cards []cardgame.Card, count int) ([]cardgame.Card, error) {
	if count <= 0 || count > len(cards) {
		return nil, fmt.Errorf("param out of range: count")
	}
	cards = this.prepareValues(cards)
	return cards[len(cards)-count:], nil
}

func (this *Poker) CalcScore(cards []cardgame.Card, count int) (int, []cardgame.Card, error) {
	cards, err := this.GetMostCards(cards, count); if err != nil {
		return 0, nil, err
	}
	score := 0
	for _, card := range cards {
		score += card.Value
	}
	return score, cards, nil
}