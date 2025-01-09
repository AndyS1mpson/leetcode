package binary_search

/*
Задача:

	Дан массив nums, отсортированный в неубывающем порядке. И дан target
	Написать функцию поиска target в nums. Если target существует, вернуть индекс, иначе -1

Пример:

	Input: nums = [-1,0,3,5,9,12], target = 9
	Output: 4

Сложность: O(log(N))
*/
func Search(nums []int, target int) int {
	if len(nums) == 0 {
		return -1
	}

	l := 0
	r := len(nums)

	for {
		if r == l && nums[r] != target {
			return -1
		}

		mid := l + (r-l)/2

		if nums[mid] == target {
			return mid
		}

		if nums[mid] < target {
			l = mid
		} else {
			r = mid
		}
	}
}
