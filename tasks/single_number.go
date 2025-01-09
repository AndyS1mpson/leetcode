package tasks

/*
Задача:

	Дан непустой массив nums, каждый элемент встречается дважды кроме одного
	Нужно его найти

	Нужно реализовать решение с линейной сложностью во время выполнения и
	использовать только постоянное дополнительное пространство.

Пример:

	Input: nums = [2,2,1]
	Output: 1

	Input: nums = [4,1,2,1,2]
	Output: 4
*/
func SingleNumber(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	res := 0

	for _, num := range nums {
		res ^= num
	}

	return res
}
