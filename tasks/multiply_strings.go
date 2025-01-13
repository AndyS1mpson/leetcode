package tasks

import (
	"strings"
)

/*
Задача:

	Даны 2 неотрицательных целых числа, представленных как строки.
	Вернуть произведение num1 и num2 также в виде строки

	Примечение: Нельзя использовать встроенную библиотеку Biginteger
	иди преобразовывать входные данные в целое число

Пример:

	Input: num1 = "2", num2 = "3"
	Output: "6"

	Input: num1 = "123", num2 = "456"
	Output: "56088"

Сложность:
*/
func Multiply(num1 string, num2 string) string {
	if num1 == "0" || num2 == "0" {
		return "0"
	}

	sum := make([]byte, len(num1)+len(num2))

	for i := len(num1) - 1; i >= 0; i-- {
		for j := len(num2) - 1; j >= 0; j-- {
			n1 := num1[i] - '0'
			n2 := num2[j] - '0'

			curr := sum[i+j+1] + n1*n2
			sum[i+j] += curr / 10
			sum[i+j+1] = curr % 10
		}
	}

	start := 0
	if sum[0] == 0 {
		start = 1
	}

	var sb strings.Builder
	for i := start; i < len(sum); i++ {
		sb.WriteByte(sum[i] + '0')
	}

	return sb.String()
}
