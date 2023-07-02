package _6threeSumClosest

import "sort"

func threeSumClosest(nums []int,target int) int {
	sort.Ints(nums)
	length := len(nums) -1
	sumClosest := nums[0] + nums[1]+nums[2]
	for i := 0 ; i < length; i++ {
		L := i + 1
		R := length
		for L < R {
			res := nums[i] + nums[L] + nums[R]
			if res == target{
				return target
			}
			if abs(res-target) < abs(sumClosest-target){
				sumClosest = res
			}
			if res < target{
				L++
			}else {
				R--
			}
		}
	}
	return sumClosest
}

func abs(x int )int {
	if x < 0 {
		return -x
	}
	return x
}