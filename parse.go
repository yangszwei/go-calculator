package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

// Token represents a token in a mathematical expression.
type Token struct {
	Type  string
	Value string
}

func ParseExpression(expression string) (Evaluable, error) {
	var tokens []Token

	// Define regular expressions to match operators and operands
	operatorPattern := `[+\-*\/]`
	operandPattern := `\d+(\.\d+)?`

	// Combine the patterns into a single regular expression
	pattern := operatorPattern + `|` + operandPattern

	// Compile the regular expression
	regex := regexp.MustCompile(pattern)

	// Find all matches in the input string
	matches := regex.FindAllString(expression, -1)

	// Convert matches into tokens
	for _, match := range matches {
		tok := Token{}
		if strings.Contains("+-*/", match) {
			tok.Type = "Operator"
		} else {
			tok.Type = "Operand"
		}
		tok.Value = match
		tokens = append(tokens, tok)
	}

	fmt.Println(tokens)

	// Stack to keep track of operators and operands
	var operandStack []Evaluable
	var operatorStack []Token

	precedence := map[rune]int{'*': 2, '/': 2, '+': 1, '-': 1}

	for _, token := range tokens {
		if token.Type == "Operand" {
			value, err := strconv.ParseInt(token.Value, 10, 64)
			if err != nil {
				return nil, fmt.Errorf("invalid operand: %w", err)
			}
			operandStack = append(operandStack, NewConstantEvaluable(value))
		} else if token.Type == "Operator" {
			for len(operatorStack) > 0 {
				top := operatorStack[len(operatorStack)-1]
				if precedence[rune(top.Value[0])] >= precedence[rune(token.Value[0])] {
					operatorStack = operatorStack[:len(operatorStack)-1]
					right := operandStack[len(operandStack)-1]
					operandStack = operandStack[:len(operandStack)-1]
					left := operandStack[len(operandStack)-1]
					operandStack = operandStack[:len(operandStack)-1]
					operandStack = append(operandStack, NewComplexEvaluable(left, right, rune(top.Value[0])))
				} else {
					break
				}
			}
			operatorStack = append(operatorStack, token)
		}
	}

	for len(operatorStack) > 0 {
		top := operatorStack[len(operatorStack)-1]
		operatorStack = operatorStack[:len(operatorStack)-1]
		right := operandStack[len(operandStack)-1]
		operandStack = operandStack[:len(operandStack)-1]
		left := operandStack[len(operandStack)-1]
		operandStack = operandStack[:len(operandStack)-1]
		operandStack = append(operandStack, NewComplexEvaluable(left, right, rune(top.Value[0])))
	}

	if len(operandStack) != 1 {
		return nil, fmt.Errorf("invalid expression")
	}

	return operandStack[0], nil
}
