package tasks

/*
Задача:

	Даны 2 строки s и t, верните true если s это подстрока t, иначе false

	Подпоследовательность строки - это новая строка, которая образуется из исходной строки путем удаления
	некоторых (можно ни одного) символов без нарушения взаимного расположения оставшихся символов.
	(например, «ace» является подпоследовательностью «abcde», а «aec» - нет).

Пример:

	Input: s = "abc", t = "ahbgdc"
	Output: true

	Input: s = "axc", t = "ahbgdc"
	Output: false

Сложность:
*/
func IsSubsequence(s string, t string) bool {
	sPointer := 0
	tPointer := 0

	for sPointer < len(s) && tPointer < len(t) {
		if s[sPointer] == t[tPointer] {
			sPointer++
		}

		tPointer++
	}

	return sPointer == len(s)
}
