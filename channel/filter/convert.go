package filter

func ToInt(done <-chan interface{}, in <-chan interface{}) <-chan int {
	out := make(chan int)
	go func() {
		defer close(out)
		for i := range in {
			select {
			case <-done:
				return
			case out <- i.(int):
			}
		}
	}()
	return out
}

func ToString(done <-chan interface{}, in <-chan interface{}) <-chan string {
	out := make(chan string)
	go func() {
		defer close(out)
		for i := range in {
			select {
			case <-done:
				return
			case out <- i.(string):
			}
		}
	}()
	return out
}
