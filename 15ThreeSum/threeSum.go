/*
 * @Descripttion: 三数之和MID
 * @Author: downson
 * @Date: 2022-09-07 14:36:08
 * @LastEditTime: 2022-09-08 12:34:43
 */
package threesum

import (
	"sort"
)

func ThreeSum(nums []int) [][]int {

	goal := 0

	sort.Ints(nums) //排序 目的:避免重复
	for i := 0; i < len(nums); i++ {
		//从小到大并且跳过重复值
		if i > 0 && nums[i] == nums[i-1] {
			continue
		}

	}

}

func twoSum(taget, value, begin, end int) []int {

}
