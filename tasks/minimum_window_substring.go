package tasks

import (
	"math"
)

/*
Задача:

	Даны две строки s и t длиной m и n соответственно.
	Вернуть минимальное окно подстроки s в которой каждая буква из t (включая дубликаты) включена в окно.
	Если такого окна нет, вернуть пустую строку.

Пример:

	Input: s = "ADOBECODEBANC", t = "ABC"
	Output: "BANC"
	Explanation: Минимальная подстрока "BANC" включает 'A', 'B', и 'C' из строки t.

	Input: s = "a", t = "a"
	Output: "a"
	Explanation: The entire string s is the minimum window.

Сложность:
*/
func MinWindow(s string, t string) string {
	targetLetters := make(map[rune]int, len(t))

	for _, let := range t {
		targetLetters[let]++
	}

	start, minStart, minLen := 0, 0, math.MaxInt32
	runeS := []rune(s)
	count := len(t)

	// скипаем начало строки в котором нет букв из таргета
	for _, let := range s {
		if _, ok := targetLetters[let]; !ok {
			start++
			minStart++
		} else {
			break
		}
	}

	for end := start; end < len(s); end++ {
		if val, ok := targetLetters[runeS[end]]; ok {
			if val > 0 {
				count--
			}
			targetLetters[runeS[end]]--
		}

		// пытаемся минимизировать окно
		for count == 0 {
			if end-start+1 < minLen {
				minStart = start
				minLen = end - start + 1
			}

			if _, ok := targetLetters[runeS[start]]; ok {
				targetLetters[runeS[start]]++
				if targetLetters[runeS[start]] > 0 {
					count++
				}
			}

			start++
		}
	}

	if minLen == math.MaxInt32 {
		return ""
	}

	return s[minStart : minStart+minLen]
}
