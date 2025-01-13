package sum

// Необходимо вернуть индексы двух любых элементов, сумма которых == target
func TwoSum(nums []int, target int) []int {
	m := make(map[int]int, len(nums))

	for i, n := range nums {
		if v, ok := m[target-n]; ok {
			return []int{v, i}
		}

		m[n] = i
	}

	return []int{}
}
