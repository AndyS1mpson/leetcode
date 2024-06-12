package tasks

import "fmt"

/*
Задача:
	Дан отсортированный массив уникальных чисел
	Дмапозоном [a, b] называется множество всех чисел между a и b (включительно).
	Верни наименьший отсортированный список диапозонов который покрывает все числа в массива.
	То есть каждый элемент nums покрывается ровно одним из диапазонов, и не существует целого числа x такого,
	что x находится в одном из диапазонов, но не в nums
Пример:
	Input: nums = [0,1,2,4,5,7]
	Output: ["0->2","4->5","7"]
	Explanation: The ranges are:
		[0,2] --> "0->2"
		[4,5] --> "4->5"
		[7,7] --> "7"

*/
func SummaryRanges(nums []int) []string {
	if len(nums) == 0 {
		return []string{}
	}

	result := make([]string, 0, len(nums))

	l := 0

	for r := range nums {
		if r == len(nums) - 1 || nums[r+1] != nums[r]+1 {
			if r == len(nums) || r == l {
				result = append(result, fmt.Sprintf("%d", nums[r]))
			} else {
				result = append(result, fmt.Sprintf("%d->%d", nums[l], nums[r]))
			}

			l = r + 1
		}
	}

	return result
}
