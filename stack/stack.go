package stack

import "fmt"

// 数组实现栈 时间复杂度都是O(1)
//对于入栈操作来说，最好情况时间复杂度是 O(1)，最坏情况时间复杂度是 O(n)

type ArrStack struct {
	items []string //数组
	//count int      //栈中元素个数
	//n     int      //栈的大小
	// 栈顶指针
    top int
}

// 初始化数组
func NewArrayStack() *ArrStack{
	return &ArrStack{
		items: []string{"a","b"},
		top: -1,
	}
}

func (this *ArrStack) IsEmpty() bool{
	// 如果栈顶为负数 证明nil
	if this.top < 0 {
		return true
	}
	return false
}
// 入栈
func (this *ArrStack) Push(v string) {
	if this.top < 0 {
		this.top = 0
	}else {
		this.top +=1
	}
	if this.top > len(this.items)-1 {
		this.items = append(this.items,v)
	}else {
		this.items[this.top] = v
	}
}



// 出栈操作
func (this *ArrStack) Pop() string  {
	if this.IsEmpty(){
		return ""
	}
	lv := this.items[this.top]
	this.top -=1
	return lv
}


func (this *ArrStack) Top() string {
	if this.IsEmpty(){
		return ""
	}
	return this.items[this.top]
}

func (this *ArrStack) Flush(){
	this.top = -1
}

func (this *ArrStack) Print() {
	if this.IsEmpty(){
		fmt.Println("empty stack")
	}else {
		for i:= this.top; i >= 0 ; i-- {
			fmt.Println(this.items[i])
		}
	}
}