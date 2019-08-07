package gotime

import (
	"os"
)

func Dir() string {
	dir, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return dir
}
