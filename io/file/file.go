package file

import (
	"os"
)


func Open(filename string) (file *os.File, err error) {
	file, err = os.Open(filename)
	if err != nil {
		return nil, err
	}
	// @causion defer is needed in caller!
	
	return file, err
}

//func Read(p []byte) (n int, err error) {
//	out, err := build(obj.file)
//	if err != nil {
//		panic(err)
//	}
//	return out, err
//}

//
//type FileReader struct {
//	file     *os.File
//	reader   io.Reader
//}
//
//
//func (obj FileReader) Init(filename string) error {
//	file, err := os.Open(filename)
//	if err != nil {
//		panic(err)
//	}
//
//	obj.file = file
//	return nil
//}
//
//
//
//type FileWritor struct {
//	file     *os.File
//	writer   io.Writer
//}
//
//
//func (obj FileWritor) Init(filename string) error {
//	file, err := os.Create(filename)
//	if err != nil {
//		panic(err)
//	}
//	defer func() {
//		err = file.Close()
//		if err != nil {
//			panic(err)
//		}
//	}()
//
//	obj.file = file
//	return nil
//}
//
//func (obj FileWritor) Write(in <-chan int, build WriteBuild) error {
//	writer := bufio.NewWriter(obj.file)
//	defer func() {
//		err := writer.Flush()
//		if err != nil {
//			panic(err)
//		}
//	}()
//
//	err := build(writer, in)
//	if err != nil {
//		panic(err)
//	}
//	return nil
//}
