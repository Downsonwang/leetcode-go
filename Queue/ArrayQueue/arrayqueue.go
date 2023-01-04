package ArrayQueue

// 数组实现队列 —— Golang ABCDEFG

type ArrayQue struct{

	  items []string
	  n int  //数组大小
      head int  //队头下标
	  tail int  //队尾下标
}

// 初始化队列
func NewArrayQueue() *ArrayQue {
	return &ArrayQue{items: []string{"A","B","C","D"}}
}

//入队操作
func (this *ArrayQue) Enqueue(v string) bool{
	// tail== n 表示队列末尾没有空间了.
	if this.tail == this.n{
		// tail == n && head == 0 表示整个队列都占满了

		if this.head == 0 {
			return false
		}
		// 数据搬移
		for i := this.head ;i < this.tail;i++{
			this.items[i-this.head] = this.items[i]
		}
		// 搬移完之后重新更新head和tail
		this.tail = this.tail - this.head
		this.head = 0
	}
	this.items[this.tail] = v
	this.tail++
	return true
}


// 出队操作
func (this *ArrayQue) DeQueue() interface{}{
    //如果head == tail 表示队列为空
	if this.head == this.tail {
	    return nil
	}
	ret := this.items[this.head]
	this.head++
	return ret
}