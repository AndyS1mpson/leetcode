package graph

import (
	"leetcode/tasks/custom_structs"
	"math"
)

/*
Задача:

	Дан корень бинарного дерева root, нужно определить валидное ли это бинарное дерево поиска (Binary Search Tree).
	Условия Binary Search Tree:
		1. Левое поддерево узла содержит только элементы с ключем меньшим чем ключ узла
		2. Справа содержатся все ключи больше
		3. Левое и правое поддерево также являются BST

Пример:

	Input: root = [2,1,3]		(2 -> 1, 2 -> 3)
	Output: true

	Input: root = [5,1,4,null,null,3,6]
	Output: false
*/
func IsValidBST(root *custom_structs.TreeNode) bool {
	maxVal := math.MaxInt
	minVal := math.MinInt

	return validate(root, minVal, maxVal)
}

func validate(root *custom_structs.TreeNode, min int, max int) bool {
	if root == nil {
		return true
	}

	if min < root.Val && root.Val < max {
		return validate(root.Left, min, root.Val) && validate(root.Right, root.Val, max)
	} else {
		return false
	}
}
