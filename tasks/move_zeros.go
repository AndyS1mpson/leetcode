package tasks

/*
Задача:
	Дан массив nums, передвиньте все 0 в конец, сохранив относительный порядок ненулевых элементов.
Пример:
	Input: nums = [0,1,0,3,12]
	Output: [1,3,12,0,0]
*/
func MoveZeroes(nums []int)  {
    l := 0

	for r := range nums {
		if nums[r] == 0 {
			if nums[l] != 0 {
				l = r
			}
		} else {
			nums[l], nums[r] = nums[r], nums[l]
			l++
		}
	}
}
