package main

import "fmt"

func Evaluate(left int64, right int64, op rune) (int64, error) {
	switch op {
	case '+':
		return left + right, nil
	case '-':
		return left - right, nil
	case '*':
		return left * right, nil
	case '/':
		return left / right, nil
	}
	return 0, fmt.Errorf("unrecognized operator: %c", op)
}
