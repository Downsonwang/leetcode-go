package _4swapnodesinpairs

type ListNode struct {
	Val int
	Next *ListNode
}

func swapPairs(head *ListNode) *ListNode{
	if head == nil {
		return nil
	}

	heads := &ListNode{Val: 0}
	l2 := &ListNode{}
	heads.Next = head  //哨兵结点指向了头结点
	l1 := heads

	for l1.Next != nil && l1.Next.Next != nil {
		l2 = l1.Next   // l2表达了l1的指向
		l1.Next = l2.Next  // l1指向结点2
		l2.Next = l2.Next.Next //对结点1指向结点3
		l1.Next.Next = l2 //指针反向
		l1 = l1.Next.Next //

	}
	return l1
}