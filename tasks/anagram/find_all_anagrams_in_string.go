package tasks

import "fmt"

/*
Задача:
	Даны 2 строки s и p, нужно вернуть массив всех начальных индексов анограм p в s в любом порядке
Пример:
	Input: s = "cbaebabacd", p = "abc"
	Output: [0,6]

	Input: s = "abab", p = "ab"
	Output: [0,1,2]
*/
func FindAnagrams(s string, p string) []int {
	var list []int

	if len(s) < len(p) {
		return list
	}

	freqS := make([]int, 26)
	freqP := make([]int, 26)

	for i := 0; i < len(p); i++ {
		freqS[s[i]-'a']++
        freqP[p[i]-'a']++
	}

	start := 0
	end := len(p)

	if fmt.Sprint(freqS) == fmt.Sprint(freqP) {
		list = append(list, start)
	}

	for end < len(s) {
		freqS[s[start]-'a']--
		freqS[s[end]-'a']++

		if fmt.Sprint(freqS) == fmt.Sprint(freqP) {
                list = append(list, start+1)
        }

		start++
		end++
	}

	return list
}
