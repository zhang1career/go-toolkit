package cardgame

type Card struct {
	Value       int
	Suit        int
}

type Game struct {
	ValueMap    map[int]string
	SuitMap     map[int]string
}

func New() (*Game, error) {
	return &Game{}, nil
}