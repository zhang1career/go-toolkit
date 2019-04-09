package texas

import (
	"github.com/zhang1career/lib/cardgame/poker"
)

type Texas struct {
	poker.Poker
}

func New() (*Texas, error) {
	p, err := poker.New()
	if err != nil {
		return nil, err
	}
	
	texas := Texas{
		*p,
	}
	return &texas, nil
}
