package tasks

/*
Задача:
	Даны 2 отсортированных односвязных списка
	Нужно объединить в один отсортированный список
Пример:
	Input: list1 = [1,2,4], list2 = [1,3,4]
	Output: [1,1,2,3,4,4]
Сложность: O(n)
*/
func MergeTwoLists(list1 *ListNode, list2 *ListNode) *ListNode {
	result := &ListNode{}

	temp := result

	for list1 != nil || list2 != nil {
		if list1 == nil {
			temp.Next = list2
			list2 = list2.Next

			temp = temp.Next
			continue
		}

		if list2 == nil {
			temp.Next = list1
			list1 = list1.Next

			temp = temp.Next
			continue
		}

		if list1.Val <= list2.Val {
			temp.Next = list1
			list1 = list1.Next

			temp = temp.Next
			continue
		}

		if list2.Val < list1.Val {
			temp.Next = list2
			list2 = list2.Next

			temp = temp.Next
			continue
		}
	}

	return result.Next
}
