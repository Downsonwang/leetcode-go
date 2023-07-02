package _39slidingWindowMaximu

import "testing"

//测试

func TestMaxSlidingWindow(t *testing.T) {
	 nums := []int{1,3,-1,-3,5,3,6,7}
	 k := 3
	res := MaxSlidingWindow(nums,k)
	t.Log(res)
}