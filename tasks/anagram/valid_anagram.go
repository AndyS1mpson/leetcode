package anagram

/*
Задача:

	Даны 2 строки s и t, вернуть true если t это анаграмма s, иначе false

Пример:

	Input: s = "anagram", t = "nagaram"
	Output: true

	Input: s = "rat", t = "car"
	Output: false

Сложность: O(N)
*/
func IsAnagram(s string, t string) bool {
	letS := make(map[rune]int, len(s))

	for _, let := range s {
		letS[let]++
	}

	for _, let := range t {
		if _, ok := letS[let]; !ok {
			return false
		}

		letS[let]--
		if letS[let] == 0 {
			delete(letS, let)
		}
	}

	return len(letS) == 0
}
