package cardgame

func (this *Game) Show(cards []Card) []string {
	ret := make([]string, len(cards))
	for i, card := range cards {
		ret[i] = this.ValueMap[card.Value] + this.SuitMap[card.Suit]
	}
	return ret
}