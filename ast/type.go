package ast

type Valuable interface {
	Evaluate() interface{}
}

type Calculable interface {
	Calc(map[string]Valuable) interface{}
}

type Item map[string]interface{}