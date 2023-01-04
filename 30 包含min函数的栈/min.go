package main

type MinStack struct {
	minStack []int
	Stack    []int
}

/** initialize your data structure here. */
func Constructor() MinStack {
	return MinStack{
		minStack: []int{},
		Stack:    []int{},
	}

}
func diff(x, y int) int {
	if x > y {
		return x
	}
	return y
}
func (this *MinStack) Push(x int) {
	this.Stack = append(this.Stack, x)
	top := this.minStack[len(this.Stack)-1]
	this.minStack = append(this.minStack, diff(x, top))
}

func (this *MinStack) Pop() {
	this.Stack = this.Stack[:len(this.Stack)-1]
	this.minStack = this.minStack[:len(this.minStack)-1]

}

func (this *MinStack) Top() int {
	return this.Stack[len(this.Stack)-1]
}

func (this *MinStack) Min() int {
	return this.minStack[len(this.minStack)-1]

}
