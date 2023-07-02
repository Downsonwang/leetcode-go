package _0Pow

import "testing"

func TestMyPow(t *testing.T) {
	x := 2.00000
	res := myPow(x,5)
	mres := mPow(x,5)
	t.Log(mres)
	t.Log(res)
}
