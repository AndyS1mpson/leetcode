package graph

import "leetcode/tasks/custom_structs"

/*
Задача:

	Дан корень бинарного дерева root, вернуть зигзагообразный обход значений узлов в порядке уровней
	(т. е. слева направо, затем справа налево для следующего уровня и чередование между ними)

Пример:

	Input: root = [3,9,20,null,null,15,7]
	Output: [[3],[20,9],[15,7]]

Сложность:
*/
func ZigzagLevelOrder(root *custom_structs.TreeNode) [][]int {
	if root == nil {
		return [][]int{}
	}

	result := make([][]int, 0)

	q := []*custom_structs.TreeNode{root}

	isReverseLevel := false

	for len(q) > 0 {
		// массив для добавления следующих элементов
		nextLevel := make([]*custom_structs.TreeNode, 0)

		level := make([]int, 0, len(q))

		// добавляем элементы в результирующий массив и сохраняем элементы следующего уровня
		for len(q) > 0 {

			if q[0] != nil {
				nextLevel = append(nextLevel, q[0].Left, q[0].Right)
				level = append(level, q[0].Val)
			}
			q = q[1:]
		}

		if isReverseLevel {
			// меняем порядок следования в результирующем массиве
			for i, j := 0, len(level)-1; i < j; i, j = i+1, j-1 {
				level[i], level[j] = level[j], level[i]
			}
		}

		q = append(q, nextLevel...)
		isReverseLevel = !isReverseLevel

		if len(level) != 0 {
			result = append(result, level)
		}
	}

	return result
}
