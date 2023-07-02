package recusion

// 迭代实现阶乘
type Fac struct {
	val map[int]int
}

func NewFactorial(n int) *Fac{
	return &Fac{
		make(map[int]int,n),
	}
}

func (this *Fac) Factorial(n int) int{
   if this.val[n] != 0 {
	   return this.val[n]
   }
   if n <= 1 {
	   this.val[n] = 1
	   return 1
   }else {
	   res := n * this.Factorial(n-1)
	   this.val[n] = res
	   return res
   }
}

func (this *Fac) Print(n int){
	println(this.val[n])
}