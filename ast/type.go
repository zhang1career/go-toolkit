package ast

type Observable interface {
	GetValue() string
}

type Evaluable interface {
	Evaluate() interface{}
}

type Calculable interface {
	Calc([]Evaluable) interface{}
}

type Item map[string]interface{}