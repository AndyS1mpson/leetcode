package tasks

/*
Задача:

	Дан отсортированный в неубывающем порядке массив nums, нужно удалить дубликаты на месте так,
	чтобы каждый уникальны йэлемент встречался только один раз

	Относительный порядок элементов должен остаться прежним, вернуть количество уникальных элементов в nums

Пример:

	Input: nums = [1,1,2]
	Output: 2, nums = [1,2,_]

	Input: nums = [0,0,1,1,1,2,2,3,3,4]
	Output: 5, nums = [0,1,2,3,4,_,_,_,_,_]

Сложность:
*/
func RemoveDuplicates(nums []int) int {
	if len(nums) == 0 {
		return 0
	}

	lastUnique := 0

	for i := 1; i < len(nums); i++ {
		if nums[i] != nums[i-1] {
			nums[lastUnique + 1] = nums[i]
			lastUnique++
		}
	}

	return lastUnique + 1
}
