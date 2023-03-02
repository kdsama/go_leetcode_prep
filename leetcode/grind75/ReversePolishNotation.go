package grind75

import (
	"fmt"
	"strconv"
)

func evalRPN(tokens []string) int {

	stack := []int{}
	symbol_arr := map[string]string{"+": "+", "*": "*", "/": "/", "-": "-"}
	for i := 0; i < len(tokens); {
		if _, ok := symbol_arr[tokens[i]]; !ok {
			x, _ := strconv.Atoi(tokens[i])
			stack = append(stack, x)

		} else {
			len_a := len(stack)

			a := stack[len_a-1]
			b := stack[len_a-2]
			fmt.Println(a, b, tokens[i])
			val := getVal(b, a, tokens[i])

			stack = stack[:len_a-2]

			stack = append(stack, val)
		}
		i++
	}
	return stack[0]
}

func getVal(a int, b int, operator string) int {

	switch operator {
	case "+":
		return a + b
	case "/":
		return a / b
	case "-":
		return a - b

	}
	return a * b
}
