package ArrayQueue

// 循环队列
// 一般情况下只需要 tail + 1 == head 即可判定队列已满，结合特殊情况（head == 0，tail == n - 1）可使用公式 (tail + 1) % n == head
//判定队列已满，head == tail 判断队列为空，但这种公式判断循环队列会浪费一个数组的存储空间。
//另也可使用一个变量 size 记录当前队列中的数量（入队 +1，出队 -1），
//当 head == tail 时，如果 size == n，则队列已满，如果 size == 0 则队列为空。

// 避免了数据的搬移

type CircularListQueue struct {
	items []string
	head int
	tail int
	length int
}

func NewCircularListQueue() *CircularListQueue{
	return &CircularListQueue{items:nil,head: 0,tail: 0,length: 0}
}

func (this *CircularListQueue) Enqueue(v string) bool {
    // 队列为满
	if ((this.tail + 1) % this.length == this.head){
		return false
	}
	this.items[this.tail] = v
	this.tail = (this.tail + 1) % this.length
	return true
}

//出队
func (this *CircularListQueue) DeQueue() interface{}{
	if this.head == this.tail {
		return nil
	}
	ret := this.items[this.head]
	this.head = (this.head + 1) % this.length
	return ret
}