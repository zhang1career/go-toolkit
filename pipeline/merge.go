package pipeline


func Merge(in1, in2 <-chan int) <-chan int {
	out := make(chan int, ChanBuffSize)
	go func() {
		v1, ok1 := <-in1
		v2, ok2 := <-in2
		for ok1 || ok2  {
			if ok1 && ok2 {
				if v1 <= v2 {
					out <- v1
					v1, ok1 = <-in1
				} else {
					out <- v2
					v2, ok2 = <-in2
				}
			} else if !ok2 {
				out <- v1
				v1, ok1 = <-in1
			} else {
				out <- v2
				v2, ok2 = <-in2
			}
		}
		close(out)
	}()
	return out
}

func MergeN(ins ...<-chan int) <-chan int {
	n := len(ins)
	if n == 1 {
		return ins[0]
	}

	m := len(ins) >> 1
	return Merge(MergeN(ins[:m]...), MergeN(ins[m:]...))
}
