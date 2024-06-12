package tasks

/*
Задача:

	Дано n - сколько пар скобок нужно использовать.
	Необходимо вернуть всевозможные правильные комбинации n пар скобок

Пример:

	Input: n = 3
	Output: ["((()))","(()())","(())()","()(())","()()()"]
*/
func GenerateParenthesis(n int) []string {
	combinations := []string{}

	backtrack(&combinations, n, 0, 0, "")

	return combinations
}

func backtrack(res *[]string, n, opened, closed int, currentCombo string) {
	if len(currentCombo) == n*2 {
		*res = append(*res, currentCombo)
		return
	}

	if opened < n {
		backtrack(res, n, opened+1, closed, currentCombo+"(")
	}

	if opened > closed && closed < n {
		backtrack(res, n, opened, closed+1, currentCombo+")")
	}
}
