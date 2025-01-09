package tasks

import "strings"

/*
Задача:

	Дана строка s, измените порядок символов в каждом слове в предложении, сохранив пробелы и начальный порядок слов.

Пример:

	Input: s = "Let's take LeetCode contest"
	Output: "s'teL ekat edoCteeL tsetnoc"

	Input: s = "Mr Ding"
	Output: "rM gniD"

Сложность:
*/
func ReverseWords(s string) string {
	words := strings.Fields(s)

	for i, word := range words {
		reverserWord := reverseString(word)
		words[i] = reverserWord
	}

	return strings.Join(words, " ")
}

func reverseString(s string) string {
	runes := []rune(s)

	left, right := 0, len(runes)-1

	for left < right {
		runes[left], runes[right] = runes[right], runes[left]
		left++
		right--
	}

	return string(runes)
}
