package tasks

import (
	"strconv"
)

/*
Задача:

	Дан массив токенов которые представляют арифметическое выражение в Польской нотации.
	Нужно вычислить выражение.
	Напоминание:
	- Валидные операторы: "+", "-", "*", "/"
	- Каждый операнд можешь быть числом или другим выражением
	- Деление двух целых чисел всегда усекается до нуля.
	- Деления на ноль не будет.
	- Входные данные представляют собой правильное арифметическое выражение в обратной польской нотации.
	- Ответ и все промежуточные вычисления могут быть представлены в виде 32-битного целого числа.

Пример:

	Input: tokens = ["2","1","+","3","*"]
	Output: 9
	Explanation: ((2 + 1) * 3) = 9

	Input: tokens = ["4","13","5","/","+"]
	Output: 6
	Explanation: (4 + (13 / 5)) = 6

	Input: tokens = ["10","6","9","3","+","-11","*","/","*","17","+","5","+"]
	Output: 22
	Explanation: ((10 * (6 / ((9 + 3) * -11))) + 17) + 5
		= ((10 * (6 / (12 * -11))) + 17) + 5
		= ((10 * (6 / -132)) + 17) + 5
		= ((10 * 0) + 17) + 5
		= (0 + 17) + 5
		= 17 + 5
		= 22

Сложность: O(n)
*/
func EvalRPN(tokens []string) int {
	operators := map[string]func(int, int) int{
		"+": func(a, b int) int { return a + b },
		"-": func(a, b int) int { return a - b },
		"*": func(a, b int) int { return a * b },
		"/": func(a, b int) int { return a / b },
	}

	stack := []int{}

	for _, token := range tokens {
		if calculate, exist := operators[token]; exist {
			a, b := stack[len(stack)-2], stack[len(stack)-1]
			stack = append(stack[:len(stack)-2], calculate(a, b))
		} else {
			num, _ := strconv.Atoi(token)
			stack = append(stack, num)
		}
	}

	return stack[0]
}
