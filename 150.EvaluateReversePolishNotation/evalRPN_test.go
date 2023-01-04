package _50_EvaluateReversePolishNotation

import "testing"

func TestEvalRPN(t *testing.T) {
	var s = []string{"2","1","+","3","*"}
	res := evalRPN(s)
	t.Log(res)
}
