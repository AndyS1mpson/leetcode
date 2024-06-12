package tasks

var buttons = map[byte][]byte{
	'2': {'a', 'b', 'c'},
	'3': {'d', 'e', 'f'},
	'4': {'g', 'h', 'i'},
	'5': {'j', 'k', 'l'},
	'6': {'m', 'n', 'o'},
	'7': {'p', 'q', 'r', 's'},
	'8': {'t', 'u', 'v'},
	'9': {'w', 'x', 'y', 'z'},
}

/*
Задача:
	Необходимо вернуть всевозможные комбинации букв для введенной последовательности цифр телефона.

Пример:

	Input: digits = "23"
	Output: ["ad","ae","af","bd","be","bf","cd","ce","cf"]
*/
func LetterCombinations(digits string) []string {
	result := []string{}

	for len(digits) > 0 {
		cur := buttons[digits[len(digits)-1]]
		digits = digits[:len(digits)-1]

		new_comb := []string{}

		if len(result) > 0 {
			for i := 0; i < len(cur); i++ {
				for j := 0; j < len(result); j++ {
					new_comb = append(new_comb, string(cur[i])+result[j])
				}
			}
		} else {
			for _, val := range cur {
				new_comb = append(new_comb, string(val))
			}
		}

		result = make([]string, len(new_comb))

		copy(result, new_comb)
	}

	return result
}
