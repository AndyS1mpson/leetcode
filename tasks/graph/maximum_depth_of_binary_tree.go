package graph

import "leetcode/tasks/custom_structs"

/*
Задача:

	Дан корень бинарного дерева, нужно вернуть максимальную глубину дерева

Пример:

	Input: root = [3,9,20,null,null,15,7]
	Output: 3

	Input: root = [1,null,2]
	Output: 2

Сложность: O(N)
*/
func MaxDepth(root *custom_structs.TreeNode) int {
	if root == nil {
		return 0
	}

	leftDepth := MaxDepth(root.Left)
	rightDepth := MaxDepth(root.Right)

	return 1 + max(leftDepth, rightDepth)
}
