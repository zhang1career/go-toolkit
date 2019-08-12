package ast

type Valuable interface {
	Evaluate() interface{}
}

type Calculable interface {
	Calc([]Valuable) interface{}
}

type Item map[string]interface{}