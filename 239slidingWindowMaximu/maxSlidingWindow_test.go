package _39slidingWindowMaximu

import "testing"

//测试

func TestMaxSlidingWindow(t *testing.T) {
	 nums := []int{1,-1}
	 k := 1
	res := MaxSlidingWindow(nums,k)
	t.Log(res)
}