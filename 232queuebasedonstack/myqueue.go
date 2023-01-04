package _32queuebasedonstack

type MyQueue struct {
      stack1 []int //永恒栈
	  stack2 []int //暂时栈
}


func Constructor() MyQueue {
  return MyQueue{}
}


func (this *MyQueue) Push(x int)  {
   this.stack1 = append(this.stack1,x)
}


func (this *MyQueue) Pop() int {
   for len(this.stack1) > 0 {
	   top := this.stack1[0]
	   this.stack1 = this.stack1[1:]
	   this.stack2 = append(this.stack2,top)
   }
   this.stack1,this.stack2 = this.stack2,this.stack1
    ret := this.stack1[0]
	this.stack1 = this.stack1[1:]
	return ret
}


func (this *MyQueue) Peek() int {
    return this.stack1[0]

}


func (this *MyQueue) Empty() bool {
     return len(this.stack1) == 0
}