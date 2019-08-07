package ast

type Valuable interface {
	GetValue() interface{}
}

type Calculable interface {
	Calc() interface{}
}