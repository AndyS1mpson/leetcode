package tasks

import (
	"fmt"
	"unicode/utf8"
)

/*
Задача:

	Дана строка s. Мы хотим разбить строку на как можно большее количество частей,
	чтобы каждая буква встречалась не более чем в одной части.

	Обратите внимание, что разбиение проводится таким образом, чтобы после конкатенации всех частей по порядку получилась строка s.

	Вернуть список целых чисел, обозначающих размер этих частей

Пример:

	Input: s = "ababcbacadefegdehijhklij"
	Output: [9,7,8]
	Explanation:
		Разбивка: «ababcbaca», «defegde», «hijhklij».
		Это такое разбиение, что каждая буква встречается не более чем в одной части.
		Такое разбиение, как «ababcbacadefegde», «hijhklij», является неправильным, так как разбивает s на меньшее количество частей.

Сложность: O(N)
*/
func PartitionLabels(s string) []int {
	lastIdx := make(map[rune]int, utf8.RuneCountInString(s))

	// сохраняем последнее встречание каждой буквы
	for idx, c := range s {
		lastIdx[c] = idx
	}

	var (
		start int
		end int
		ans []int
	)

	// движемся по строке пока не дойдем до самой последней позиции из букв встретившихся на пути
	// потом повторяем 
	for idx, _ := range s {
		end = max(end, lastIdx[rune(s[idx])])
		fmt.Println(string(s[idx]))
		if idx == end {
			ans = append(ans, idx - start + 1)
			start = idx + 1
		}
	}

	return ans
}
