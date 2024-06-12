package tasks

import "sort"

/*
Найти самый длинный общий префикс у входных слов
Пример:
	Input: strs = ["flower","flow","flight"]
	Output: "fl"
*/
func LongestCommonPrefix(strs []string) string {
	if len(strs) == 0 {
		return strs[0]
	}

	sort.Strings(strs)

	l := len(strs)

	for i := range strs[0] {
		if strs[0][i] != strs[l - 1][i] {
			return strs[0][:i]
		}
	}

	return strs[0]
}
