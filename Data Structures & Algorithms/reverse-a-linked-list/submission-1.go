/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil{
		return head
	}
    st := []int{}
	tmp := head

	for tmp != nil{
		st = append(st,tmp.Val)
		tmp = tmp.Next
	}


	newHead := &ListNode{Val: st[len(st)-1]}
	temp := newHead

	for i := len(st)-2;i>=0;i--{
		tmp := &ListNode{Val: st[i]}
		temp.Next = tmp
		temp = tmp
	}

	return newHead
}
