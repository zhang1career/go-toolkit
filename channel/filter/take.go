package filter

func Take(done <-chan interface{}, in <-chan interface{}, num int) <-chan interface{} {
	out := make(chan interface{})
	go func() {
		defer close(out)
		for i := 0; i < num; i++ {
			select {
			//@todo 关闭不是原子操作，关闭后可能有1个数据进入out
			case <-done:
				return
			case out <-<- in:
			}
		}
	}()
	return out
}
