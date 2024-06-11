package tasks

type ListNode struct {
	Val  int
	Next *ListNode
}

func count(a, b, increase int) (int, int) {
	res := a + b + increase
	tenCount := res / 10

	return tenCount, res % 10
}

func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	increase := 0

	res := &ListNode{}
	tmpRes := res

	for l1 != nil || l2 != nil || increase != 0 {
		tmpRes.Next = &ListNode{}
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
