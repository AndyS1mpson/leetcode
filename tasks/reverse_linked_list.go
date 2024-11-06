package tasks

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }

Задача:
	Дана голова однонаправленного массива, нужно развернуть лист
Пример:
	Input: head = [1,2,3,4,5]
	Output: [5,4,3,2,1]
 */
 func ReverseList(head *ListNode) *ListNode {
	if head == nil || head.Next == nil {
        return head
    }

	reversed := &ListNode{
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
