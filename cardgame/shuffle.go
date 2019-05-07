package cardgame

import (
	"math/rand"
	"time"
)

func sprinkle(salt int)  {
	rand.Seed(time.Now().Unix() ^ (int64(salt)<<32 | int64(salt)))
}

func (this *Game) KnuthDurstenfeldShuffle(cards []Card, salt int) []Card {
	sprinkle(salt)
	for i := len(cards); i > 0; i-- {
		r := rand.Intn(i)
		cards[r], cards[i-1] = cards[i-1], cards[r]
	}
	return cards
}