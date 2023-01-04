package _32queuebasedonstack

// 双指针链表实现
type Queue struct {
  Length int
  Pre *Value
  Suf *Value
}

type Value struct {
  Val int
  Next *Value
}

/** Initialize your data structure here. */

func MConstructor() Queue{
	return Queue{}
}

/** Push element x to the back of queue. */

func (this *Queue) Push(x int) {
    val := &Value{Val: x}
    // 入队 尾部不为nil
    if this.Suf != nil {
       this.Suf.Next = val //插入一个新结点
       this.Suf = val  //移动尾指针
    }else {
       this.Suf = val  //若尾部为nil 直接赋值
    }
    if this.Pre == nil{  //链表为nil 直接no 1
      this.Pre = val
    }
    this.Length++
}

/** Removes the element from in front of queue and returns that element.**/

func (this *Queue) Pop() int {
    //出队 从头部出队 指针变化 长度变化
    val := this.Pre
    this.Pre = val.Next
    this.Length--
    return val.Val
}

func (this *Queue) Peek() int{
   return this.Pre.Val
}

func (this *Queue) Empty() bool {
  return this.Length == 0
}