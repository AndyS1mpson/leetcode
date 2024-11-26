package graph

import "leetcode/structs"

/*
Задача:

	Дан корень бинарного дерева, необходимо вернуть последовательный обход значений его узлов

Пример:

	Input: root = [1,null,2,3], 1 -> 2, 2 -> 3
	Output: [1,3,2]

	Input: root = [1,2,3,4,5,null,8,null,null,6,7,9]
	Output: [4,2,6,5,7,1,3,9,8]
*/
func InorderTraversal(root *TreeNode) []int {
	result := []int{}

	if root == nil {
		return result
	}

	current := root

	stack := structs.NewStack[*TreeNode]()

	for current != nil || !stack.IsEmpty() {
		for current != nil {
			stack.Push(current)
			current = current.Left
		}

		var ok bool
		current, ok = stack.Pop()
		if ok {
			result = append(result, current.Val)
			current = current.Right
		}
	}

	return result
}

func InorderTraversalV2(root *TreeNode) []int {
	if root == nil {
		return nil
	}

	var values []int
	values = append(values, InorderTraversalV2(root.Left)...)
	values = append(values, root.Val)
	values = append(values, InorderTraversalV2(root.Right)...)

	return values
}
