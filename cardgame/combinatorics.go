package cardgame

func (this *Game) Permutate(cards [][]Card) [][]Card {
	ret := make([][]Card, 0)
	
	length := len(cards)
	if length <= 1 {
		for _, card := range cards[length-1] {
			ret = append(ret, []Card{card})
		}
		return ret
	}
	
	for _, straight := range this.Permutate(cards[1:]) {
		for _, head := range cards[0] {
			tmpStraight := make([]Card, len(straight))
			copy(tmpStraight, straight)
			tmpStraight = append(tmpStraight, head)
			ret = append(ret, tmpStraight)
		}
	}
	
	return ret
}

func (this *Game) Combinate(cards []Card, count int) [][]Card {
	if len(cards) < count || count <= 0 {
		return nil
	}
	
	ret := make([][]Card, 0)
	if count == 1 {
		for _, card := range cards {
			ret = append(ret, []Card{card})
		}
		return ret
	}
	
	for i := 0; i <= len(cards) - count; i++ {
		for _, com := range this.Combinate(cards[i+1:], count-1) {
			tmpCom := make([]Card, len(com))
			copy(tmpCom, com)
			tmpCom = append(tmpCom, cards[i])
			ret = append(ret, tmpCom)
		}
	}
	return ret
}