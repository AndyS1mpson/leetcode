package graph

import "leetcode/tasks/custom_structs"

/*
Задача:

	Дан корень бинарного дерева root, вернуть сумму всех левых листьев

Пример:

	Input: root = [3,9,20,null,null,15,7]
	Output: 24
	Explanation: There are two left leaves in the binary tree, with values 9 and 15 respectively.

Сложность: O(N)
*/
func SumOfLeftLeaves(root *custom_structs.TreeNode) int {
	if root == nil {
		return 0
	}

	sum := 0

	q := []*custom_structs.TreeNode{root}

	for len(q) > 0 {
		node := q[0]
		q = q[1:]

		if node == nil {
			continue
		}

		if node.Left != nil {
			if node.Left.Left == nil && node.Left.Right == nil {
				sum += node.Left.Val
			} else {
				q = append(q, node.Left)
			}
		}

		q = append(q, node.Right)

	}

	return sum
}
