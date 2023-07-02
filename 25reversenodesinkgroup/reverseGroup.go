package _5reversenodesinkgroup


// Leetcode 25. k个一组翻转链表
// 1. 如果节点总数不是k的整数倍,那么请将最后剩余节点保持原有顺序
// 2. 不能只是单纯的改变节点内部的值,而是需要实际进行节点交换
// 3. O(1)额外内存空间算法


type ListNode struct {
	Val int
	Next *ListNode
}

func reverseGroup(head *ListNode,k int) *ListNode{
    n := 0
	for cur := head; cur != nil;cur =cur.Next{
		n++ //统计节点个数
	}
	dummy := &ListNode{Next: head}
	p0 := dummy
	var pre,cur *ListNode = nil,p0.Next
	for ; n >= k ; n-=k {
		for i :=0; i < k ;i++{
			nxt := cur.Next
			cur.Next = pre
			pre = cur
			cur = nxt
		}
		// 见视频
		nxt := p0.Next
		p0.Next.Next = cur
		p0.Next = pre
		p0 = nxt
	}
	return dummy.Next
}
