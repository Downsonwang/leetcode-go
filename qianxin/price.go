package qianxin



func maxProfit(price []int) int {
	if len(price) < 2 {
		return 0
	}
	var m = len(price)
	var dp = make([][]int,m)
	for i := range dp {
		dp[i] = make([]int,2)

	}
	dp[0][0] = 0
	dp[0][1]=-price[0]
	for i:=1;i<m;i++{
		dp[i][0]=max(dp[i-1][0],dp[i-1][1]+price[i])
		dp[i][1]=max(dp[i-1][1],-price[i])
	}
	return dp[m-1][0]

}
func max(x,y int)int{
	if x < y {
		return y
	}
	return x
}

func main(){

}