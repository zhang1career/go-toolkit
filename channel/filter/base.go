package filter

func Iterate(done <-chan interface{}, in <-chan interface{}, fn func(interface{}) interface{}) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		defer close(out)
		for i := range in {
			select {
			case <-done:
				return
			case out <- fn(i):
			}
		}
	}()
	return out
}
