package gate

import (
	"fmt"
)

func Or(done <-chan interface{}, varIns ...<-chan interface{}) (<-chan interface{}, error) {
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
		return or2(done, ins[0], ins[1]), nil
	}
	return Or(done, append(varIns[2:], or2(done, ins[0], ins[1]))...)
}

func or2(done <-chan interface{}, in1, in2 <-chan interface{}) <-chan interface{} {
	var out = make(chan interface{})
	var bufferIns []interface{}
	var bufferIn interface{}
	var inSize = 2
	go func() {
		defer close(out)
		for inSize > 0 || len(bufferIns) > 0 {
			var maybeOut chan<- interface{}
			
			if len(bufferIns) > 0 {
				maybeOut = out
				bufferIn = bufferIns[0]
			}
			
			select {
			case <-done:
				return
			case in, ok := <-in1:
				if ok {
					bufferIns = append(bufferIns, in)
				} else {
					in1 = nil
					inSize = inSize - 1
				}
			case in, ok := <-in2:
				if ok {
					bufferIns = append(bufferIns, in)
				} else {
					in2 = nil
					inSize = inSize - 1
				}
			case maybeOut <- bufferIn:
				bufferIns = bufferIns[1:]
			}
		}
	}()
	return out
}
