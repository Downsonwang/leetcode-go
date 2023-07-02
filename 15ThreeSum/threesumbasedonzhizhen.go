package threesum

import "sort"

func threeSumB(nums []int) (res [][]int) {
	sort.Ints(nums)
	for i := 0 ; i < len(nums) - 1 ; i++{
	    if i > 0 && nums[i] == nums[i-1]{
			continue
		}
		L := i + 1
		R := len(nums) - 1
		for L < R {
			sum := nums[i] + nums[L] + nums[R]
			if sum == 0 {
				res = append(res,[]int{nums[i],nums[L],nums[R]})
				for L < R && nums[L] == nums[L+1] {
					L++
				}
				for L < R && nums[R] == nums[R-1]{
					R--
				}
				L++
				R--
			}else if sum > 0 {
				R--
			}else if sum < 0 {
				L++
			}
		}
	}
	return res
}