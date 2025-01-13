package sum

import "sort"

/*
Необходмо найти все тройки чисел [nums[i], nums[j], nums[k]], такие что:
	i != j, i != k, j != k,
	и nums[i] + nums[j] + nums[k] == 0

В результате не должно быть дубликатов триплетов.
*/
func ThreeSum(nums []int) [][]int {
	// Сортируем массив
	sort.Ints(nums)

	var result [][]int

	for numIdx := 0; numIdx < len(nums) - 2; numIdx++ {
		// Пропускаем все дубликаты слева
		// numIdx > 0 обеспечивает проверку только со второго элемента и далее иначе на первой итерации была бы паника
		if numIdx > 0 && nums[numIdx] == nums[numIdx-1] {
			continue
		}

		num2Idx := numIdx + 1
		num3Idx := len(nums) - 1

		for num2Idx < num3Idx {
			sum := nums[numIdx] + nums[num2Idx] + nums[num3Idx]
			if sum == 0 {
				result = append(result, []int{ nums[numIdx], nums[num2Idx], nums[num3Idx]})
			
				num3Idx--

				// Пропускаем все дубликаты справа
				for num2Idx < num3Idx && nums[num3Idx] == nums[num3Idx + 1] {
					num3Idx--
				}
			} else if sum > 0 {
				num3Idx--
			} else {
				num2Idx++
			}
			
		}
	}

	return result
}
