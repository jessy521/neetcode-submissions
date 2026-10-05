/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */

func removeNthFromEnd(head *ListNode, n int) *ListNode {
    count := 0;
	temp := head

	for temp != nil {
		count++;
		temp = temp.Next
	}

	if count-n == 0{
		return head.Next
	}

	temp = head
	i := 1

	for temp != nil{
		// fmt.Printf( " %d - %d == %d\n" ,count , i,n)
		if count - i == n{
			temp.Next = temp.Next.Next
		}
		temp = temp.Next
		i++
	}

	return head
}
