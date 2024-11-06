package tasks

/*
Задача:

	Вам даны 2 списка закрытых интервалов firstList и secondList,
	где firstList[i] = [start_i, end_i] and secondList[j] = [start_j, end_j]
	Каждый лист интервалов попарно расходится и располагается в отсортированном порядке.

	Вернуть пересечение этих двух списков.

	Замкнутый интервал [a, b] (при a <= b) обозначает множество вещественных чисел x таких что a <= x <= b.

	Пересечение двух замкнутых интервалов - множество вещественных чисел, которое либо пусто, либо представлено в виде замкнутого интервала.
	Например, пересечением [1, 3] и [2, 4] является [2, 3]

Пример:

	Input:
		firstList = [[0,2],[5,10],[13,23],[24,25]],
		secondList = [[1,5],[8,12],[15,24],[25,26]]
	Output: [[1,2],[5,5],[8,10],[15,23],[24,24],[25,25]]
*/
func IntervalIntersection(firstList [][]int, secondList [][]int) [][]int {
	res := make([][]int, 0, len(firstList)+len(secondList))

	for f, s := 0, 0; f < len(firstList) && s < len(secondList); {
		fStart := firstList[f][0]
		fEnd := firstList[f][1]

		sStart := secondList[s][0]
		sEnd := secondList[s][1]

		start := max(fStart, sStart)
		end := min(fEnd, sEnd)

		if start <= end {
			res = append(res, []int{start, end})
		}

		if fEnd <= sEnd {
			f++
		} else {
			s++
		}
	}

	return res
}
