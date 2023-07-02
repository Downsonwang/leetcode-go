package _5reverseKGroup

type ListNode struct {
	Val int
	Next *ListNode
}

func reverseKGroup(head *ListNode,k int) *ListNode{
	cur := head
	for i :=0; i < k ; i++ {
		if cur == nil {
			return head
		}
		cur = cur.Next
	}
	newHead := reverse(head,cur)
	head.Next = reverseKGroup(cur,k)
	return newHead
}

func reverse(start,end *ListNode) *ListNode{
	var pre *ListNode
	cur := start
	for cur != end {
		nxt := cur.Next
		cur.Next = pre
		pre = cur
		cur = nxt
	}
	return pre
}