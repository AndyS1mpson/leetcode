package graph

import "leetcode/tasks/custom_structs"

/*
Задача:

	Дан корень бинарного дерева, нужно проверить симметрично ли оно или нет относительно корня

Пример:

	Input: root = [1,2,2,3,4,4,3]
	Output: true

	Input: root = [1,2,2,null,3,null,3]
	Output: false

Сложность: O(N)
*/
func IsSymmetric(root *custom_structs.TreeNode) bool {
	if root == nil {
		return false
	}

	q1 := []*custom_structs.TreeNode{root.Left}
	q2 := []*custom_structs.TreeNode{root.Right}

	for len(q1) > 0 && len(q2) > 0 {
		currentNode1 := q1[0]
		currentNode2 := q2[0]

		q1 = q1[1:]
		q2 = q2[1:]

		if currentNode1 == nil && currentNode2 == nil {
			continue
		}

		if currentNode1 == nil || currentNode2 == nil || currentNode1.Val != currentNode2.Val {
			return false
		}

		q1 = append(q1, currentNode1.Left, currentNode1.Right)
		q2 = append(q2, currentNode2.Right, currentNode2.Left)
	}

	return true
}
