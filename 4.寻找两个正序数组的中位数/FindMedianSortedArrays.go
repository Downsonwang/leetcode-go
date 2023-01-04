package __寻找两个正序数组的中位数

import (
	"fmt"
	"sort"
)

// nums1 =  [1,3] , nums2 = [2]   [1,2,3] 2
// nums2 = [1,2] , nums2 = [3,4]  [1,2,3,4] （2 + 3 ）/ 2 = 2.5
// nums1 = [0,0] , nums2 = [0,0]

func FindMedianSortedArrays(nums1 []int, nums2 []int) float64 {
	//var s float64
	var num []int
	tmp := 0
	half := 0.0
	if nil == nums1 && nil == nums2 {
		return 0.0
	}
	if nil == nums1 && nil != nums2 {
		for i := 0; i < len(nums2); i++ {
			return float64(nums2[i])
		}
	}
	if nil != nums1 && nil == nums2 {
		for i := 0; i < len(nums1); i++ {
			return float64(nums1[i])
		}
	}

	if nil != nums1 && nil != nums2 {
		num = append(num, nums1...)
		num = append(num, nums2...)
		sort.Ints(num)
		fmt.Println(num)

		for i := 0; i < len(num)-1; i++ {
			if num[i] > num[i+1] {
				tmp = num[i]
				num[i] = num[i+1]
				num[i+1] = tmp
			}

		}
		if len(num)%2 == 0 {
			firstLen := len(num) / 2
			secondLen := len(num)/2 + 1
			i2 := num[firstLen-1]
			fmt.Println(i2)
			i3 := num[secondLen-1]
			fmt.Println(i3)
			i4 := float64(i2)/2.0 + float64(i3)/2.0
			half = float64(i4)
			fmt.Println(half)

		} else {
			halfLen := (len(num) + 1) / 2
			i2 := num[halfLen-1]
			half = float64(i2)
			fmt.Println("half", half)
		}
		fmt.Println(" show all num ", num)
	}

	return half
}
