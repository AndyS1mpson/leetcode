package tasks

/*
Задача:

	Дан массив nums, найти подмассив максимальной суммы и вернуть сумму

Пример:

	Input: nums = [-2,1,-3,4,-1,2,1,-5,4]
	Output: 6
	Explanation: The subarray [4,-1,2,1] has the largest sum 6.

	Input: nums = [5,4,-1,7,8]
	Output: 23
	Explanation: The subarray [5,4,-1,7,8] has the largest sum 23.

Идея алгоритма:

	Движемся по массиву слева направо и накапливаем в переменной curSum текущую частичную сумму.
	Если в какой-то момент curSum окажется отрицательной, то присвоим curSum = 0.
	Максимум из всех значений переменной curSum за время прохода по массиву и будет ответом на задачу.

Сложность: O(n)
*/
func MaxSubArray(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	globalMax, curSum := nums[0], 0

	for _, num := range nums {
		curSum += num
		if curSum > globalMax {
			globalMax = curSum
		}
		if curSum < 0 {
			curSum = 0
		}
	}

	return globalMax
}
