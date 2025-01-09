package linked_list

import "leetcode/tasks/custom_structs"

/*
Задача:

	Дан указатель на массив, необходимо удалить из него n'ый членс конца

Пример:

	Input: head = [1,2,3,4,5], n = 2
	Output: [1,2,3,5]

Сложность: O(N)
*/
func RemoveNthFromEnd(head *custom_structs.ListNode, n int) *custom_structs.ListNode {
	if head == nil {
		return nil
	}

	res := &custom_structs.ListNode{
		Next: head,
	}

	// Идея в том что мы в начале уводим более быстрый указатель на N позиций вперед
	lead := res
	for i := 0; i <= n; i++ {
		lead = lead.Next
	}

	// Более медленный указатель стоит на начальной позиции
	cur := res

	// После этого мы двигаем указатели одновременно, таким образом когда быстрый дойдет до конца, медленный будет также на расстоянии N от него
	for lead != nil {
		cur = cur.Next
		lead = lead.Next
	}

	cur.Next = cur.Next.Next

	return res.Next
}
