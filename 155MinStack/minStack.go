package _55MinStack

// 1. 单纯的切片实现栈
type MinStack struct {

	items []int //将队头作为栈顶

}

func Constructor() MinStack{
   return MinStack{}
}

func (this *MinStack) Push(val int) {
	//插入
	this.items = append(this.items,val)
}

func (this *MinStack) Pop(){
	this.items = this.items[:len(this.items)-1]
}

func (this *MinStack) Top() int {

	top := this.items[len(this.items)-1]

	return top
}

func (this *MinStack) GetMin() int {
	 min := this.items[0]
	 for _,v := range this.items{
		 if v < min {
			 min = v
		 }
	 }
   return min
}