package _0Pow


//实现pow(x,n）,即计算x的整数n次幂函数(即 xN)
// myPow 简单递归解决,递归调用n次,所以时间复杂度为O（N）,空间复杂度O(1)
func myPow(x float64,n int) float64{
  var helper func(float64,int) float64
  helper = func(x float64, n int) float64 {
	  if n == 0 {
		  return 1
	  }
	  return helper(x,n-1) * x
  }
  if n >= 0 {
	  return helper(x,n)
  }
  return 1.0 / helper(x,-n)
}
