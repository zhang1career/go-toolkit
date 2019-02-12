package manager

import "fmt"

func OrDone(done <-chan interface{}, varIns ...<-chan interface{}) (<-chan interface{}, error) {
	inSize := len(varIns)
	if inSize <= 0 {
		return nil, fmt.Errorf("no input given")
	}
	if inSize == 1 {
		return varIns[0], nil
	}
	
	ins := make([]<-chan interface{}, 0)
	for _, varIn := range varIns {
		ins = append(ins, varIn)
	}
	
	if inSize == 2 {
		return orDone2(done, ins[0], ins[1]), nil
	}
	return OrDone(done, append(varIns[2:], orDone2(done, ins[0], ins[1]))...)
}

func orDone2(done <-chan interface{}, in1, in2 <-chan interface{}) <-chan interface{} {
	var out = make(chan interface{})
	go func() {
		defer close(out)
		select {
		case <-done:
			return
		case <-in1:
		case <-in2:
		}
	}()
	return out
}