/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	var prev *ListNode
	curr := head

	for curr != nil{
		front := curr.Next
		curr.Next = prev
		prev = curr
		curr = front
	}

	return prev
}
