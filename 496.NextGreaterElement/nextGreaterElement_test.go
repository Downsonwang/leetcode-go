package _96_NextGreaterElement

import "testing"

func TestNextGreaterElement(t *testing.T) {
	 nums1 :=[]int{4,1,2}
	 nums2 :=[]int{1,3,4,2}
	res := NextGreaterElement(nums1,nums2)
    t.Log(res)
}
