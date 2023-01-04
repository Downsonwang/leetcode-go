package _25MyStack



// 仅用一个队列实现栈  切片  不断取队头到队尾
type OnlyStack struct {
	items []int
}

func MyConstructor() MyStack {
  return MyStack{items: nil}
}

func (this *OnlyStack) MyPush(x int){
       this.items = append(this.items,x)
	   for i := len(this.items); i > 1 ; i--{
		   this.items = append(this.items,this.items[0])  //队头到队尾
		   this.items = this.items[1:] //新的栈顶元素
	   }

}

func (this *OnlyStack) MyPop() int{
     top := this.items[0]
	 this.items = this.items[1:len(this.items)]
	 return top
}

func (this *OnlyStack) MyTop() int{
    return this.items[0]
}

func (this *OnlyStack) Empty() bool{
  return len(this.items) == 0
}