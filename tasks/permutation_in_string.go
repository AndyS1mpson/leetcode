package tasks

/*
Задача:

	Даны 2 строки s1 и s2. Вернуть true если s2 содержит перестановку s1, иначе false

Примеры:

	Input: s1 = "ab", s2 = "eidbaooo"
	Output: true
	Explanation: s2 содержит одну перестановку s1 ("ba").

	Input: s1 = "ab", s2 = "eidboaoo"
	Output: false
*/
func CheckInclusion(s1 string, s2 string) bool {
	l := 0
	count := [26]int{}

	for i := range s1 {
		count[s1[i] - 'a']++
	}

	for r := range s2 {
		count[s2[r]-'a']--
		if count == [26]int{} { return true}

		if r + 1 >= len(s1) {
			count[s2[l]-'a']++
			l++
		}
	}

	return false
}
