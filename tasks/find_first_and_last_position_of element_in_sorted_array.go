package tasks

/*
Задача:

	Дан массив nums отсортированный в неубывающем порядке.
	Нужно найти начальную и конечную позиции заданного target.
	Если target не найден, вернуть [-1, -1]

	Сложность: O(log n)

Пример:

	Input: nums = [5,7,7,8,8,10], target = 8
	Output: [3,4]
*/
func SearchRange(nums []int, target int) []int {
	if len(nums) == 0 {
		return []int{-1, -1}
	}

	return []int{binaryLeft(nums, target), binaryRight(nums, target)}
}

func binaryLeft(nums []int, target int) int {
	left, right := 0, len(nums)-1

	for left < right {
		mid := (right + left) / 2
		if nums[mid] < target {
			left = mid + 1
		} else {
			right = mid
		}
	}
	if nums[left] == target {
		return left
	}

	return -1
}

func binaryRight(nums []int, target int) int {
    left, right := 0, len(nums)-1
    
    for left < right {
        mid := (right + left)/2 + 1
        if nums[mid] > target {
            right = mid - 1
        } else if nums[mid] <= target {
            left = mid
        }
    }
    if nums[left] == target {
        return left
    } 

    return -1
}
