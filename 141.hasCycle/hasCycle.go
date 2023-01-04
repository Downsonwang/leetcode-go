/*
 * @Descripttion: 141 判断环形链表 - 快慢指针
 * @Author: downson
 *类似 “追及问题”
   两个人在环形跑道上赛跑，同一个起点出发，一个跑得快一个跑得慢，在某一时刻，跑得快的必定会追上跑得慢的，只要是跑道是环形的，不是环形就肯定追不上。
 * @Date: 2022-09-19 09:01:20
 * @LastEditTime: 2022-09-19 10:11:57
*/
package hascycle

type ListNode struct {
	Val  int
	Next *ListNode
}

func hasCycle(head *ListNode) bool {
	slow, fast := head, head

	for fast != nil {
		// 没环 或没值
		if fast.Next == nil {
			return false
		}
		slow = slow.Next      //慢的 走一步
		fast = fast.Next.Next //快的走两步
		if slow == fast {
			return true
		}
	}
	return false

}
