/*
 * @Descripttion:
 * @Author:
 * @Date: 2022-09-22 09:07:25
 * @LastEditTime: 2022-09-22 09:40:01
 */
package removenthfromend

import (
	"fmt"
	"testing"
)

func TestRemoveNthFromEnd(t *testing.T) {
	head := &ListNode{Val: 1, Next: &ListNode{Val: 2, Next: &ListNode{Val: 3, Next: &ListNode{Val: 4, Next: &ListNode{Val: 5}}}}}

	final := removeNthFromEnd(head, 2)
	fmt.Println(final)
}
