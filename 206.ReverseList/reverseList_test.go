/*
 * @Descripttion: 206 测试反转链表
 * @Author: downson
 * @Date: 2022-09-18 10:49:26
 * @LastEditTime: 2022-09-18 14:41:38
 */
package reverselist

import "testing"

func TestReverseList(t *testing.T) {
	lis := *&ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: nil}}}
	res := reverseList(&lis)
	t.Log(res)
}
