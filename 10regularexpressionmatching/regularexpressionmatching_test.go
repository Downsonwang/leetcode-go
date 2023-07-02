package _0regularexpressionmatching

import "testing"

func TestIsMatch(t *testing.T){
	s := "aa"
	p := "a"
  bol := isMatch(s,p)
  t.Log(bol)
}
