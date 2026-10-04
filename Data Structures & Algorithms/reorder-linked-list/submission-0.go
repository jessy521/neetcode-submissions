/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func reorderList(head *ListNode) {
    nodes := []*ListNode{}
	temp := head
	for temp!=nil{
		nodes = append(nodes,temp)
		temp = temp.Next
	}

	i, j := 0,len(nodes)-1

	for i<j{
		nodes[i].Next = nodes[j]
		i++
		if i>= j{
			break
		}
		nodes[j].Next = nodes[i]
		j--
	}

	nodes[i].Next = nil
}
