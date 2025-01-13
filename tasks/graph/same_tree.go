package graph

import "leetcode/tasks/custom_structs"

/*
Задача:

	Даны корни 2 бинарных деревьев p и q, нужно проверить одинаковые ли они

Пример:

	Input: p = [1,2,3], q = [1,2,3]
	Output: true

	Input: p = [1,2], q = [1,null,2]
	Output: false

Сложность: O(N+M)
*/
func IsSameTree(p *custom_structs.TreeNode, q *custom_structs.TreeNode) bool {
	q1 := []*custom_structs.TreeNode{p}
	q2 := []*custom_structs.TreeNode{q}

	for len(q1) > 0 && len(q2) > 0 {
		elem1 := q1[0]
		q1 = q1[1:]
		elem2 := q2[0]
		q2 = q2[1:]

		if elem1 == nil && elem2 == nil {
            continue
        }

		if elem1 != nil && elem2 != nil {
			if elem1.Val != elem2.Val {
				return false
			}

			q1 = append(q1, elem1.Left, elem1.Right)
			q2 = append(q2, elem2.Left, elem2.Right)
		} else {
			return false
		}

	}

	return true
}
