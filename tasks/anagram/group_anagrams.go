package tasks

import "sort"

/*
Задача:

	Дан массив строк. Нужно сгруппировать анаграммы вместе и вернуть ответ в любом результате

Пример:

	Input: strs = ["eat","tea","tan","ate","nat","bat"]
	Output: [["bat"],["nat","tan"],["ate","eat","tea"]]

	Input: strs = [""]
	Output: [[""]]
*/
func GroupAnagrams(strs []string) [][]string {
	res := make([][]string, 0, len(strs))

	mp := make(map[string][]string, len(strs))

	for _, str := range strs {
		sortedWord := sortString(str)

		if _, ok := mp[string(sortedWord)]; !ok {
			mp[string(sortedWord)] = make([]string, 0)
		}

		mp[string(sortedWord)] = append(mp[string(sortedWord)], str)
	}

	for _, v := range mp {
		res = append(res, v)
	}

	return res
}

func sortString(s string) string {
	characters := []rune(s)
	sort.Slice(characters, func(i, j int) bool {
		return characters[i] < characters[j]
	})
	return string(characters)
}
