package tasks

import (
	"strconv"
)

/*
Задача:

	Дан массив символов chars, сожмите его с помощью следующего алгоритма:
	Начните с пустой строки s. Для каждой группы последовательно повторяющихся символов в chars:
		* Если длина группы равна 1, добавьте этот символ в s
		* В противном случае добавьте символ, за которым следует длина группы
	Сжатая строка s не должна возвращаться отдельно, а должна храниться в массиве входных символов chars.
	Обратите внимание, что длина группы, равная 10 или более, будет разбита на несколько символов в chars.
	После завершения модификации входного массива верните его новую длину.

Примеры:

	Input: chars = ["a","a","b","b","c","c","c"]
	Output: Вернет 6, and the первые 6 символов во введенном массиве должны быть: ["a","2","b","2","c","3"]
	Explanation: The groups are "aa", "bb", and "ccc". This compresses to "a2b2c3".

	Input: chars = ["a","b","b","b","b","b","b","b","b","b","b","b","b"]
	Output: Вернет 4 и: ["a","b","1","2"].
	Explanation: The groups are "a" and "bbbbbbbbbbbb". This compresses to "ab12".
*/
func Compress(chars []byte) int {
	l, r := 0, 0

	for ;r < len(chars); {
		j := r + 1
		for ; j < len(chars) && chars[j] == chars[j - 1]; j++ {}
		
		chars[l] = chars[r]
		l++

		if j != r + 1 {
			for _, number := range(strconv.Itoa(j - r)) {
				chars[l] = byte(number)
				l++
			}
		}
		r = j
	}

	chars = chars[:l]
	return len(chars)
}
