package sum

/*
Задача:

	Дан массив целых чисел, который уже отсортирован в неубывающем порядке.
	Найдите 2 числа, которые в сумме дают target.
	Пусть эти два числа будут  numbers[index1] и numbers[index2] где 1 <= index1 < index2 <= numbers.length

	Ваше решение должно использовать только постоянное дополнительное пространство.

Пример:

	Input: numbers = [2,7,11,15], target = 9
	Output: [1,2]
	Explanation: Сумма 2 и 7 это 9. Поэтому index1 = 1, index2 = 2. Возвращаем [1, 2].

Сложность: O(N)
*/
func TwoSum2(numbers []int, target int) []int {
	for l, r := 0, len(numbers)-1; ; {
		switch {
		case numbers[l]+numbers[r] > target:
			r--
		case numbers[l]+numbers[r] < target:
			l++
		default:
			return []int{l + 1, r + 1}
		}
	}
}
