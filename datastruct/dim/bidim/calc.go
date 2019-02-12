package bidim

func (a *Coordinate) Add(b Coordinate) (sum Coordinate) {
	return Coordinate{a.X + b.X, a.Y + b.Y}
}
