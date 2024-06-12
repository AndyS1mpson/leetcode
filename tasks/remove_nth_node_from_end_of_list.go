package tasks

/*
Задача: Дан указатель на массив, необходимо удалить из него n'ый членс конца

Пример:
	Input: head = [1,2,3,4,5], n = 2
	Output: [1,2,3,5]
*/
func RemoveNthFromEnd(head *ListNode, n int) *ListNode {
	if head == nil {
		return nil
	}

	res := &ListNode{
		Next: head,
	}

	lead := res

	for i := 0; i <= n; i++ {
		lead = lead.Next
	}

	cur := res

	for lead != nil {
		cur = cur.Next
		lead = lead.Next
	}

	cur.Next = cur.Next.Next

	return res.Next
}
