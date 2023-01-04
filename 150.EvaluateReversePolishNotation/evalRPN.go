package _50_EvaluateReversePolishNotation

import "strconv"

func evalRPN(tokens []string) int {

	var stack []int
	for _ ,token := range tokens {
		val,err := strconv.Atoi(token)
		if err == nil {
			stack = append(stack,val)
		} else {
			num1,num2 := stack[len(stack)-2],stack[len(stack)-1]
			stack = stack[:len(stack)-2]
			switch token {
			case "+" :
				stack = append(stack,num1+num2)
			case "-":
				stack = append(stack,num1-num2)
			case "*":
				stack = append(stack,num1*num2)
			case "/":
				stack = append(stack,num1/num2)
			}

		}

	}
    return stack[0]
}
