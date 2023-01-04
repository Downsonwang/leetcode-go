package main

import (
	"container/list"
)

type CQueue struct {
	appendStack *list.List
	deleteStack *list.List
}

func Constructor() CQueue {
	return CQueue{
		appendStack: list.New(),
		deleteStack: list.New(),
	}
}

func (this *CQueue) AppendTail(value int) {
	this.appendStack.PushBack(value)
}

func (this *CQueue) DeleteHead() int {
	if this.deleteStack.Len() == 0 {
		for this.appendStack.Len() > 0 {
			this.deleteStack.PushBack(this.appendStack.Remove(this.appendStack.Back()))
		}
	}
	if this.deleteStack.Len() > 0 {
		this.deleteStack.Remove(this.deleteStack.Back())
	}

	return -1
}

func main() {

}
