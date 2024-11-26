package tasks

import (
	"math/rand"
)

/*
Задача:

	Дано API rand7(), которое генерирует случайное число в диапозоне [1, 7].
	Напишите функцию rand10(), которая генерирует случайное число в диапозоне [1, 10].
	Вы можете использовать только rand7()

Пример:

	Input: n = 1
	Output: [2]

	Input: n = 3
	Output: [3,8,10]
*/
func Rand10() int {
	// Ссылка на объяснение: https://mywebcenter.ru/realizacziya-funkczii-rand10-s-ispolzovaniem-rand7/#Основы_генерации_случайных_чисел
	result := 40

	for result >= 40 {
		result = 7*(rand7()-1) + (rand7() - 1)
	}

	return result%10 + 1
}

func rand7() int {
	return rand.Intn(7-1) + 1
}
