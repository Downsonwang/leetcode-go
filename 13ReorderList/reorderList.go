package _3ReorderList

type ListNode struct {
	Val int
	Next *ListNode
}

func reorderList(head *ListNode){

	if head == nil {
		return
	}
    length :=0
	stack := []*ListNode{}
	node := head
	if node != nil {
		length++
		stack = append(stack,node)
		node = node.Next
	}

	for i := 0 ; i < length ; i++{
	  stack[length-i] = node.Next
	  node.Next.Next = stack[length-i].Next


    }


}
