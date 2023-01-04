package _43_ReorderList


// 重排链表: 1 2 3 4

type ListNode struct {
	Val int  // 1
	Next *ListNode // 2
}

func reorderList(head *ListNode) {

	if head == nil {
		return
	}
    stack := []*ListNode{}
	node := head
	for node != nil {
		 stack = append(stack,node)
		 node = node.Next
	}

	cur,nxt := head,head.Next
	for nxt != nil {
		tail := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if tail == nxt{
			break
		}
		stack[len(stack)-1].Next = nil
		cur.Next = tail
		tail.Next = nxt
		cur = nxt
		nxt = nxt.Next
	}


}