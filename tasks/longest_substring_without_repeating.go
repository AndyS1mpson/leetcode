package tasks

/* 
Задача:
	Найти самую длинную подстроку со всеми уникальными буквами внутри и вернуть ее длину

Пример:
	Input: abcsabsf
	Output: 4
*/
func LengthOfLongestSubstring(s string) int {
	var (
		store               = make(map[byte]int)
		left, right, result int
	)

	for right < len(s) {
		var r = s[right]
		store[r] += 1
		for store[r] > 1 {
			l := s[left]
			store[l] -= 1
			left += 1
		}
		result = max(result, right-left+1)

		right += 1
	}
	return result
}
