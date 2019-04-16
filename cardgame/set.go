package cardgame

func (this *Game) AddCards(addedCards []Card, addingCards []Card) []Card {
	for j := 0; j < len(addingCards); j++ {
		isThere := false;
		for i := 0; i < len(addedCards); i++ {
			if addedCards[i].Value == addingCards[j].Value && addedCards[i].Suit == addingCards[j].Suit {
				isThere = true
				break
			}
		}
		if !isThere {
			addedCards = append(addedCards, addingCards[j])
		}
	}
	return addedCards
}

func (this *Game) DelCards(subedCards []Card, subingCards []Card) []Card {
	for j := 0; j < len(subingCards); j++ {
		for i := 0; i < len(subedCards); i++ {
			if subedCards[i].Value == subingCards[j].Value && subedCards[i].Suit == subingCards[j].Suit {
				subedCards = append(subedCards[:i], subedCards[i+1:]...)
				break
			}
		}
	}
	return subedCards
}