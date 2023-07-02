package _0Pow

func mPow(x float64,n int) float64 {
	var helper func(float64 , int) float64
	helper = func(x float64, n int) float64 {
		if n == 0 {
			return 1
		}
		t := helper(x,n/2)
		if n%2 == 1{
			return t * t * x
		}
		return t * t
	}
	if n >=0 {
		return helper(x,n)
	}
	return 1.0 / helper(x,-n)
}
