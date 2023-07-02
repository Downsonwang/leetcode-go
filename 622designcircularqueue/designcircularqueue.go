package _22designcircularqueue

type MyCircularQueue struct {
	items []int

	head int
	tail int
}


func Constructor(k int) MyCircularQueue {
      return MyCircularQueue{
		  items: make([]int,k+1),
	  }
}


func (this *MyCircularQueue) EnQueue(value int) bool {
	if this.IsFull(){
		return false
	}
	this.items[this.tail] = value
	this.tail = (this.tail + 1) % len(this.items)
	return true
}


func (this *MyCircularQueue) DeQueue() bool {
	if this.IsEmpty() {
		return false
	}

	this.head = (this.head + 1) % len(this.items)
	return true
}


func (this *MyCircularQueue) Front() int {
     if this.IsEmpty() {
		 return -1
	 }
	 ret := this.items[this.head]
	//this.head = (this.head + 1) % this.length
    return ret
}


func (this *MyCircularQueue) Rear() int {
	if this.IsFull(){
		return -1
	}
	rep := this.items[(this.tail - 1 + len(this.items))%len(this.items)]
	return rep
}


func (this *MyCircularQueue) IsEmpty() bool {
    if this.head == this.tail {
		return true
	}
	return false
}


func (this *MyCircularQueue) IsFull() bool {
	if ((this.tail + 1) % len(this.items) == this.head){
		return true
	}
	return false
}
