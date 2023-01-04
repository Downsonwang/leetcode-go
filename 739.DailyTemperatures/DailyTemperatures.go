package _39_DailyTemperatures

//来源：力扣（LeetCode）
//链接：https://leetcode.cn/problems/daily-temperatures
// 单调栈: 通常是一维数组,
// 739
func dailyTemperatures(temperatures []int) []int{

     ans := make([]int,len(temperatures))
     stack := []int{}
     for i,v := range temperatures{
         // 栈不空,且遍历当前元素
         for len(stack) != 0 && v > temperatures[stack[len(stack)-1]]{
             //pop
             top := stack[len(stack)-1]
             stack = stack[:len(stack)-1]
             ans[top] = i - top
         }
         stack = append(stack,i)
     }
    return ans
}
