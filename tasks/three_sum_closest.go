package tasks

import (
	"math"
	"sort"
)

/*
Дан массив nums длинной n и target. Нужно найти 3 числа из nums,
такие что сумма наиболее близка к target и вернуть сумму

Пример:
	Input: nums = [-1,2,1,-4], target = 1
	Output: 2
*/
func ThreeSumClosest(nums []int, target int) int {
	if len(nums) == 0 {
		return 0
	}

	sort.Ints(nums)

	max := math.MaxInt32
	closest := 0

	for idx1 := 0; idx1 < len(nums) - 2; idx1++ {
		l := idx1 + 1
		r := len(nums) - 1

		for l < r {
			sum := nums[idx1] + nums[l] + nums[r]

			if sum == target {
				return sum
			} else if sum < target {
				l++
			} else {
				r--
			}

			diff := math.Abs(float64(sum - target))
			if int(diff) < max {
				max = int(diff)
				closest = sum
			}
		}
	}

	return closest
}
