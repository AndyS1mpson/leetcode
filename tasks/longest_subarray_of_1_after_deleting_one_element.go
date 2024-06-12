package tasks

/*
Задача:

	Дан бинарный массив nums, нужно удалить 1 элемент из него.
	Вернуть длину наибольшего подмассива, который содержит только 1. Вернуть 0 если такого подмассива не существует

Пример:

	Input: nums = [1,1,0,1]
	Output: 3
	Explanation: после удаления номера на позиции 2, [1, 1, 1] содержит 3 числа 1.
*/
func LongestSubarray(nums []int) int {
	l := 0
	maxOnes := 0
	delePos := -1

	for r := range nums {
		if nums[r] == 0 {
			maxOnes = max(maxOnes, r - l - 1)
			l = delePos + 1
			delePos = r
		} else if r == len(nums) - 1 {
			maxOnes = max(maxOnes, r - l)
		}
	}

	return maxOnes
}
