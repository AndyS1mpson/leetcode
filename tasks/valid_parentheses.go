package tasks

import "leetcode/structs"

/*
Задача:

	Дана строка s, состоящая из '(', ')', '{', '}', '[', ']'
	Нужно проверить правильная ли s скобачная последовательность или нет

Пример:

	Input: s = "()[]{}"
	Output: true

	Input: s = "(]"
	Output: false
*/
func IsValid(s string) bool {
	stack := structs.NewStack[rune]()

	for _, letter := range s {
		if letter == '(' || letter == '{' || letter == '[' {
			stack.Push(letter)
		} else {
			if stack.Len() == 0 {
				return false
			}

			top, _ := stack.Pop()
			if (letter == ')' && top != '(') ||
				(letter == '}' && top != '{') ||
				(letter == ']' && top != '[') {
				return false
			}
		}
	}

	if stack.Len() != 0 {
        return false
    }

	return true
}
