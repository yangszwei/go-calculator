package main

import "fmt"

func main() {
	input := "1 + 1 - (5 * 2)"

	evaluable, err := ParseExpression(input)
	if err != nil {
		panic(err)
	}

	fmt.Println(evaluable.Evaluate())
}
