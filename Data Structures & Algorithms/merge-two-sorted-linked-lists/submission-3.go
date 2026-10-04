/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
    if l1 == nil {return l2}
	if l2 == nil {return l1}

	head := &ListNode{}
	temp := head

	for l1 != nil && l2 != nil {
		if l1.Val < l2.Val{
			temp.Next = l1
			l1 = l1.Next
		}else{
			temp.Next = l2
			l2 = l2.Next
		}
		temp = temp.Next
	}

	temp.Next = l1
	if l1 == nil {
		temp.Next = l2
	}

	return head.Next
}
