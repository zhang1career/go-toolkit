package bidim

import "fmt"

func Printf(message string, aa [][]int) {
	for _, a := range aa {
		fmt.Printf(message, a)
	}
}
