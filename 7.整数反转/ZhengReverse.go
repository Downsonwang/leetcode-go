package __整数反转

import "math"

/*
   给你一个 32 位的有符号整数 x ，返回将 x 中的数字部分反转后的结果。
   如果反转后整数超过 32 位的有符号整数的范围 [−231,  231 − 1] ，就返回 0。
*/

func reverse(x int) int {
	var nums, newNums int
	if x > math.MaxInt32 || x < math.MinInt32 {
		return 0
	}
	for x <= math.MaxInt32 && x >= math.MinInt32 && x != 0 {

		a := x % 10
		newNums = nums*10 + a
		nums = newNums
		x = x / 10

	}
	if nums > math.MaxInt32 || nums < math.MinInt32 {
		return 0
	}
	return nums

}
