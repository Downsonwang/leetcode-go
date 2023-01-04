/*
 * @Descripttion: MID 19 删除链表的倒数第N个结点,并且返回链表的头结点.
 * @Author: Dowoson
 * @Date: 2022-06-01 09:52:41
 * @LastEditTime: 2022-06-01 11:21:18
 */
package deletelinklist

// 给你个链表,删除链表的倒数第n个节点,并返回链表的头结点
type ListNode struct {
	Val  int
	Next *ListNode
}

func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	//目标初始声明
	target := n
	var stack []*ListNode
	//链表为空
	if head == nil {
		return head
	}
	d := &ListNode{0, head}

	//链表不为空则遍历
	for d.Next != nil {
		stack = append(stack, d)
		d = d.Next
	}

	// 将指针指向下一个节点
	length := len(stack) - target
	for r := range stack {
		if r == length {
			stack[r].Next = stack[r].Next.Next
		}
		//return d.Next
	}
	return d.Next
}
