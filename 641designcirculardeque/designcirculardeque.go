package _41designcirculardeque


type MyCircularDeque struct {
    items []int
	head int
	tail int
}

func Constructor(k int) MyCircularDeque {
  return MyCircularDeque{items: make([]int,k+1),head: 0,tail: 0}
}


func (this *MyCircularDeque) InsertFront(value int) bool  {
      if this.IsFull() {
		  return false
	  }
	  this.head = (this.head - 1 + len(this.items)) % len(this.items) //移动队头指针
	  this.items[this.head] = value
	  return true

}

func (this *MyCircularDeque) InsertLast(value int) bool  {
    if this.IsFull(){
		return false
	}
	this.items[this.tail] = value
	this.tail = (this.tail + 1 ) % len(this.items)
	return true
}

func (this *MyCircularDeque) DeleteFront() bool {
   if this.IsEmpty(){
	   return false
   }

   //修改队头指针,将其加一,表示删除队头元素 因为是
   this.head = (this.head + 1) % len(this.items)
   return true
}


func (this *MyCircularDeque) DeleteLast() bool {
	if this.IsEmpty(){
		return false
	}
	this.tail = (this.tail - 1 + len(this.items)) % len(this.items)
	return true
}


func (this *MyCircularDeque) GetFront() int {
    if this.IsEmpty() {
		return -1
	}
	front := this.items[this.head]
	return front
}


func (this *MyCircularDeque) GetRear() int {
    if this.IsEmpty(){
		return -1
	}
	rear := this.items[(this.tail -1 + len(this.items)) % len(this.items) ]    // 循环队列 并不是对尾部单纯减1 要取余数
	return rear
}


func (this *MyCircularDeque) IsEmpty() bool {
     return this.head == this.tail
}


func (this *MyCircularDeque) IsFull() bool {
   return (this.tail+1)%len(this.items) == this.head
}