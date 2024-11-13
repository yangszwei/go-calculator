package main

import "fmt"

type Evaluable interface {
	Evaluate() (int64, error)
}

type SimpleEvaluable struct {
	left  int64
	right int64
	op    rune
}

func NewSimpleEvaluable(left int64, right int64, op rune) *SimpleEvaluable {
	return &SimpleEvaluable{left: left, right: right, op: op}
}

func (se *SimpleEvaluable) Evaluate() (int64, error) {
	return Evaluate(se.left, se.right, se.op)
}

type ComplexEvaluable struct {
	left  Evaluable
	right Evaluable
	op    rune
}

func NewComplexEvaluable(left Evaluable, right Evaluable, op rune) *ComplexEvaluable {
	return &ComplexEvaluable{left: left, right: right, op: op}
}

func (ce *ComplexEvaluable) Evaluate() (int64, error) {
	left, err := ce.left.Evaluate()
	if err != nil {
		return 0, fmt.Errorf("error evaluating left node: %w", err)
	}

	right, err := ce.right.Evaluate()
	if err != nil {
		return 0, fmt.Errorf("error evaluating right node: %w", err)
	}

	return Evaluate(left, right, ce.op)
}

type ConstantEvaluable struct {
	value int64
}

func NewConstantEvaluable(value int64) *ConstantEvaluable {
	return &ConstantEvaluable{value: value}
}

func (ce *ConstantEvaluable) Evaluate() (int64, error) {
	return ce.value, nil
}
