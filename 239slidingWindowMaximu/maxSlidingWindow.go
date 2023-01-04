package _39slidingWindowMaximu


// 用双Queue实现,保持下标右移,window存下标,每增加一个下标,值比前面的值大就将前面的下标删除,确保最左值最大.
// 下标实现
func MaxSlidingWindow(nums []int,k int) []int{
	if nums == nil {
		return nil
	}
	// window 存下标
	window,res := []int{},[]int{}
    for index,value := range nums {
		if index >= k && window[0] <= index - k {
			window = window[1:]
		}
		// 循坏 对比window里的最后一次元素, 比x小的都移除,永远确保最左的是最大值的下标
		for len(window) > 0 && nums[window[len(window)-1]] <= value {
			window = window[:len(window)-1]
		}
		//把下标添加到window
		window = append(window,index)
		//从k-1开始 只要i>= k-1 可以将最左元素添加到结果
        if index >= k-1{
			res = append(res,nums[window[0]])
		}
	}
	return res
}