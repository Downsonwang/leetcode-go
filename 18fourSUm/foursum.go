package _8fourSUm

import "sort"

func fourSum(nums []int,target int) [][]int{
 sort.Ints(nums)
 ans := make([][]int,0)
 if len(nums) < 4 {
  return ans
 }
 for i := 0 ; i < len(nums)-3;{
  for j := i +1; j < len(nums) -2;{
   left := j + 1
   right := len(nums) -1
   for left < right{
    if nums[i] + nums[j] + nums[left] + nums[right] > target {
     right--
    }else if nums[i] + nums[j] + nums[left] + nums[right] < target{
     left++
    }else {
      ans = append(ans,[]int{nums[i],nums[j],nums[left],nums[right]})
      left = sort.SearchInts(nums,nums[left]+1)
    }
   }
   j = sort.SearchInts(nums,nums[j]+1)
  }
  i = sort.SearchInts(nums,nums[i]+1)
 }
 return ans
}
