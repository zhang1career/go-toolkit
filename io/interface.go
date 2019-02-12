package io

import "io"

type IO interface {
	Read() (<-chan int, error)
	Write(in <-chan int) error
}

type ReadBuild  func(reader io.Reader) (<-chan int, error)
type WriteBuild func(writer io.Writer, in <-chan int) error
