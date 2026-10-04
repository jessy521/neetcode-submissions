/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func mergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
    if list1 == nil && list2 != nil{return list2}
	if list2 == nil && list1 != nil{return list1}

	head := ListNode{}
	temp := &head

	l1,l2 := list1 , list2

	for l1 != nil && l2 != nil {
		min := min(l1.Val,l2.Val)
		tmp := &ListNode{Val:min}
		if temp != nil{
			temp.Next = tmp
			temp = temp.Next
		}else{
			temp = tmp
		}
		if l1.Val == min{
			l1 = l1.Next
		}else{
			l2 = l2.Next
		}
	}

	if l2 != nil {
		temp.Next = l2
	}

	if l1 != nil {
		temp.Next = l1
	}

	return head.Next
}
