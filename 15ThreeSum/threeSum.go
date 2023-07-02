/*
 * @Descripttion: 三数之和MID
 * @Author: downson
 * @Date: 2022-09-07 14:36:08
 * @LastEditTime: 2022-09-08 12:34:43
 */
package threesum

func threeSum(nums []int) [][]int{
    if nums == nil {
        return nil
    }
    length := len(nums) - 1
    separateSort(nums,0,length)
    res := make([][]int,0)


    for i := 0 ; i < length; i++{

        if i > 0 && nums[i] == nums[i-1]{
            continue
        }
        L := i+1
        R := length
        for L < R {
            sum := nums[i] + nums[L]+nums[R]
            if sum == 0 {
                res = append(res,[]int{nums[i],nums[L],nums[R]})
              for L < R && nums[L] == nums[L+1] {
                  L++
              }
              for L < R && nums[R] == nums[R-1] {
                  R--
              }
              L++
              R--
            }else if sum < 0 {
                L++
            }else if sum > 0 {
                R--
            }
        }


    }
    return  res
}

func separateSort(arr []int,start,end int)   {
    if start >= end {
        return
    }
    i := partition(arr,start,end)
    separateSort(arr,start,i-1)
    separateSort(arr,i+1,end)
}

func partition(arr []int,start,end int) int{
    pivot := arr[end]
    var i = start
    for j := start; j < end; j++{
        if arr[j] < pivot{
            if !(i==j){
                // 交换位置
                arr[i],arr[j] = arr[j],arr[i]
            }
            i++
        }
    }
    arr[i],arr[end] = arr[end],arr[i]
    return i
}