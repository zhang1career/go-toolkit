package tcp

import (
	"bufio"
	"net"
)


type Net struct {
	addr    string
	listen  net.Listener
}


func (obj Net) Init(addr string) error {
	obj.addr = addr
	return nil
}

func (obj Net) Read(build ReadBuild) (<-chan int, error) {
	out := make(chan int)
	go func() {
		conn, err := net.Dial("tcp", obj.addr)
		if err != nil {
			panic(err)
		}

		r, err := build(bufio.NewReader(conn))
		for v := range r {
			out <- v
		}
		close(out)
	}()
	return out, nil
}

func (obj Net) Write(in <-chan int, build WriteBuild) error {
	listener, err := net.Listen("tcp", addr)
	if err != nil {
		panic(err)
	}
	defer func() {
		err = listener.Close()
		if err != nil {
			panic(err)
		}
	}()

	obj.listener = listener
	return nil

	conn, err := obj.listener.Accept()
	if err != nil {
		panic(err)
	}
	defer func() {
		err = conn.Close()
		if err != nil {
			panic(err)
		}
	}()

	writer := bufio.NewWriter(conn)
	defer func() {
		err := writer.Flush()
		if err != nil {
			panic(err)
		}
	}()

	err = build(writer, in)
	if err != nil {
		panic(err)
	}
	return nil
}
