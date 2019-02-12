package pipeline

import (
	"encoding/binary"
	"io"
)


func (file File) ReadSource() (<-chan int, error) {
	out := make(chan int, ChanBuffSize)
	go func() {
		buffer := make([]byte, IOBuffSize)
		for {
			n, err := file.reader.Read(buffer)
			if n > 0 {
				v := int(binary.BigEndian.Uint64(buffer))
				out <- v
			}
			if err != nil {
				break
			}
		}
		close(out)
	}()
	return out, nil
}

func (file File) WriteSink() error {
	buffer := make([]byte, IOBuffSize)
	for v := range file.in {
		binary.BigEndian.PutUint64(buffer, uint64(v))
		_, err := file.writer.Write(buffer)
		if err != nil {
			panic(err)
		}
	}
	return nil
}
