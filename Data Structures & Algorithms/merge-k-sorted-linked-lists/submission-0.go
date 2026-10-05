/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func merge(l1 *ListNode,l2 *ListNode)*ListNode{
	if l1 == nil{return l2}
	if l2 == nil{return l1}

	head := &ListNode{}
	temp := head

	for l1 != nil && l2 != nil{
		if l1.Val <= l2.Val{
			temp.Next = l1
			l1 = l1.Next
		}else{
			temp.Next = l2
			l2 = l2.Next
		}
		temp = temp.Next
	}

	if l1 != nil{
		temp.Next = l1
	}else{
		temp.Next = l2
	}

	return head.Next
}
func divide(lists []*ListNode,left,right int) * ListNode{
	if left>right{
		return nil
	}

	if left == right{
		return lists[left]
	}

	mid := left +(right-left)/2
	l1 := divide(lists,left,mid)
	l2 := divide(lists,mid+1,right)

	return merge(l1,l2)
}

func mergeKLists(lists []*ListNode) *ListNode {
	if len(lists) == 0{return nil}
	
	return divide(lists,0,len(lists)-1)
}
