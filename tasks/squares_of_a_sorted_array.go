package tasks
/*
Задача:
	Дан массив чисел, отсортированный в неубывающем порядке. Нужно вернуть массив квадратом чисел в отсортированном виде
Пример:
	Input: nums = [-4,-1,0,3,10]
	Output: [0,1,9,16,100]

	Input: nums = [-7,-3,2,3,11]
	Output: [4,9,9,49,121]
*/
func SortedSquares(nums []int) []int {
    left := 0
	right := len(nums) - 1
	res := make([]int,0,right+1)
	

	for left <= right {
		if nums[left]*nums[left] < nums[right]*nums[right] {
			res =append(res, nums[right]*nums[right])
			right--
		} else {
			res =append(res, nums[left]*nums[left])
			left++
		}
	}
	reverseSlice(res)
	return res
}

func reverseSlice(nums []int) {
	for i, j := 0, len(nums)-1; i < j; i, j = i+1, j-1 {
		nums[i], nums[j] = nums[j], nums[i]
	}
}