package calc

import (
	"fmt"
	"github.com/zhang1career/lib/gotime"
)

type Eval func(params ...int) (result int, err error)


func Apply(op Eval, params ...int) (result int, err error) {
	fmt.Printf("Call %s\n", gotime.WhichFunc(op))
	return op(params...)
}

func Add(params ...int) (result int, err error) {
	result = 0
	for i := 0; i <len(params); i++ {
		result += params[i]
	}
	return result, nil
}

func Sub(params ...int) (result int, err error) {
	result = params[0]
	for i := 1; i <len(params); i++ {
		result -= params[i]
	}
	return result, nil
}

func Mul(params ...int) (result int, err error) {
	result = 1
	for i := 0; i <len(params); i++ {
		result *= params[i]
	}
	return result, nil
}

func Div(params ...int) (result int, err error) {
	result = params[0]
	for i := 1; i <len(params); i++ {
		result /= params[i]
	}
	return result, nil
}

func Normalize(value int, min int, max int) float64 {
	return float64(value-min) / float64(max-min)
}