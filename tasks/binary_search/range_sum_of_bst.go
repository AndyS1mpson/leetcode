package binary_search

type TreeNode struct {
	Val   int
	Left  *TreeNode
	Right *TreeNode
}

/*
Задача:

	Дан корень дерева бинарного поиска и 2 числа log и high,
	вернуть сумму значений всех нод со значениями, которые включены в диапазон [low, high].

Пример:

	Input: root = [10,5,15,3,7,null,18], low = 7, high = 15
	Output: 32
	Explanation: Ноды 7, 10, и 15 принадлежат [7, 15]. 7 + 10 + 15 = 32.

Сложность:
*/
func RangeSumBST(root *TreeNode, low int, high int) int {
	if root == nil {
		return 0
	}

	if root.Val >= low && root.Val <= high {
		return root.Val + RangeSumBST(root.Left, low, high) + RangeSumBST(root.Right, low, high)
	}

	return RangeSumBST(root.Left, low, high) + RangeSumBST(root.Right, low, high)
}
