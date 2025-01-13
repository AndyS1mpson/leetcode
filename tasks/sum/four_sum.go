package sum

import "sort"

/*
Дан массив чисел и target.

Необходимо найти все уникальные четверки чисел, такие что:

	nums[a] + nums[b] + nums[c] + nums[d] == target

# Ответ можно вернуть в любом порядке

Пример:

	Input: nums = [1,0,-1,0,-2,2], target = 0
	Output: [[-2,-1,1,2],[-2,0,0,2],[-1,0,0,1]]
*/
func FourSum(nums []int, target int) [][]int {
	sort.Ints(nums)
	res := [][]int{}

	for i := 0; i < len(nums)-3; i++ {
		// скипаем все дубли
		for i > 0 && i < len(nums) && nums[i] == nums[i-1] {
			i++
		}

		for j := i + 1; j < len(nums); j++ {
			// скипаем все дубли
			for j > i+1 && j < len(nums) && nums[j] == nums[j-1] {
				j++
			}

			start := j + 1
			end := len(nums) - 1

			// выше зафиксировали первые 2 переменные и проходимся по всевозможным значениям двух оставшихся
			for start < end {
				if nums[start]+nums[end]+nums[i]+nums[j] < target {
					start++
					// скипаем дубли
					for start < len(nums) && nums[start] == nums[start-1] {
						start++
					}
				} else if nums[start]+nums[end]+nums[i]+nums[j] > target {
					end--
					for end > 0 && nums[end] == nums[end+1] {
						end--
					}
				} else {
					res = append(res, []int{nums[i], nums[j], nums[start], nums[end]})
					start++
					end--
					for start < len(nums) && nums[start] == nums[start-1] {
						start++
					}
					for end > 0 && nums[end] == nums[end+1] {
						end--
					}
				}
			}

		}
	}

	return res
}

/*
Дано: [1,0,-1,0,-2,2], target = 0

1) Сортируем -> [-2, -1, 0, 0, 1, 2] -> O(n log(n))
2)

*/
