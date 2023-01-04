package fanzhuanlianbiao

type ListNode struct {
	Val int
	Next *ListNode
}

// 反转链表
func ReverseList(head *ListNode) *ListNode {
	  cur := head
	  var prev *ListNode = nil
	  for cur != nil {
		   prev,cur,cur.Next = cur,cur.Next,prev
	  }
	  return prev
}
