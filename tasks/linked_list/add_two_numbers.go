package linked_list

import "leetcode/tasks/custom_structs"

/*
Задача:
	На вход подается два числа, представленных в обратном порядке:
	425: 5 -> 2 -> 4
	126: 6 -> 2 -> 1

	Необходимо вернуть сумму этих чисел в таком же виде:
	1 -> 5 -> 5
*/
func AddTwoNumbers(l1 *custom_structs.ListNode, l2 *custom_structs.ListNode) *custom_structs.ListNode {
	increase := 0

	res := &custom_structs.ListNode{}
	tmpRes := res

	for l1 != nil || l2 != nil || increase != 0 {
		tmpRes.Next = &custom_structs.ListNode{}
		tmpRes = tmpRes.Next

		if l1 != nil && l2 != nil {
			increase, tmpRes.Val = count(l1.Val, l2.Val, increase)

			l1 = l1.Next
			l2 = l2.Next
			continue
		}

		if l1 != nil {
			increase, tmpRes.Val = count(l1.Val, 0, increase)
			l1 = l1.Next

			continue
		}

		if l2 != nil {
			increase, tmpRes.Val = count(0, l2.Val, increase)
			l2 = l2.Next

			continue
		}

		if increase != 0 {
			tmpRes.Val = increase
			break
		}

	}

	return res.Next
}

func count(a, b, increase int) (int, int) {
	res := a + b + increase
	tenCount := res / 10

	return tenCount, res % 10
}
