package _25MyStack

// 用队列实现栈
// 队列是先进先出 栈是后进先出
// 两个队列实现栈

type MyStack struct {
    items []int  //永久队列
	other []int  //暂时队列
}


func Constructor() MyStack {
  return MyStack{items:nil ,other: nil}
}


func (this *MyStack) Push(x int)  {

	 // 新来的加入到暂时队列
	 this.other = append(this.other,x)
	 // 永久队列遍历加到暂时队列中
	 for len(this.items) > 0 {
		 //加入暂时队列
		 this.other = append(this.other,this.items[0])
		 //更改永久队列中的TOP
		 this.items = this.items[1:]
	 }
	 //转还给永久队列
	 this.items,this.other = this.other,this.items

}


func (this *MyStack) Pop() int {
	v := this.items[0]  //取栈顶元素
	this.items = this.items[1:]  //移除并更新下一个栈顶元素
	return v
}


func (this *MyStack) Top() int {
      v := this.items[0]
	  return v
}


func (this *MyStack) Empty() bool {
     if this.items == nil {
		 return true
	 }
	 return false
}
