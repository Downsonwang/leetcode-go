package __寻找两个正序数组的中位数

import "testing"

func TestFindMedianSortedArrays(t *testing.T) {

	nums1 := []int{2, 2, 4, 4}
	nums2 := []int{2, 2, 4, 4}
	r := FindMedianSortedArrays(nums1, nums2)
	t.Log(r)

}
