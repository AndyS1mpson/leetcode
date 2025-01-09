package tasks

/*
Задача:

	Даны два массива nums1 и nums2. Нужно вернуть пересечение.
	Каждый элемент должен вернуться столько раз сколько раз он встречается в обоих массивах

Пример:

	Input: nums1 = [1,2,2,1], nums2 = [2,2]
	Output: [2,2]

	Input: nums1 = [4,9,5], nums2 = [9,4,9,8,4]
	Output: [4,9]
	Explanation: [9,4] также подходит.

	Сложность: O(n + m)
*/
func Intersect(nums1 []int, nums2 []int) []int {
	if len(nums1) == 0 || len(nums2) == 0 {
		return []int{}
	}

	if len(nums1) < len(nums2) {
		nums1, nums2 = nums2, nums1
	}

	result := make([]int, 0, len(nums2))

	m := make(map[int]int, 0)

	for _, num := range nums1 {
		m[num]++
	}

	for _, num := range nums2 {
		if v, ok := m[num]; ok {
			if v > 0 {
				result = append(result, num)
				m[num]--
			}
		}
	}

	return result
}
