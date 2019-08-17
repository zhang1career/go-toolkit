package ast

type Serializable interface {
	GetValue() string
}

type Unserializable interface {
	SetValue(string)
}

type Evaluable interface {
	Evaluate() interface{}
}

type Calculable interface {
	Calc([]Evaluable) interface{}
}

type Item map[string]interface{}