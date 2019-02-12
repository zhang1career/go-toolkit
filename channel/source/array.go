package source

func ArraySource(inputs []interface{}) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		defer close(out)
		for _, v := range inputs {
			out <- v
		}
	}()
	return out
}
