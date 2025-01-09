package linked_list

import "leetcode/tasks/custom_structs"

/*
Задача:

	Дана голова однонаправленного массива, нужно развернуть лист

Пример:

	Input: head = [1,2,3,4,5]
	Output: [5,4,3,2,1]

Сложность: O(N)
*/
func ReverseList(head *custom_structs.ListNode) *custom_structs.ListNode {
	if head == nil || head.Next == nil {
		return head
	}

	reversed := &custom_structs.ListNode{
		Next: head,
	}

	pointer := reversed.Next

	for pointer.Next != nil {
		cur := pointer.Next
		pointer.Next = cur.Next

		cur.Next = reversed.Next
		reversed.Next = cur

	}

	return reversed.Next
}
