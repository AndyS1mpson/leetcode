package tasks

import "sort"

/*
Задача:

	Даны интервалы где intervals[i] = [start_i, end_i],
	Соедините все пересекающиеся интервалы и верните массив непересекающихся интервалов,
	которые покрывают все интервалы во входных данных

Пример:

	    Input: intervals = [[1,3],[2,6],[8,10],[15,18]]
		Output: [[1,6],[8,10],[15,18]]
		Explanation: Интервалы [1,3] [2,6] перекрываются, соедините их в [1,6].

Сложность: O(N*log(N))
*/
func Merge(intervals [][]int) [][]int {
	// сортируем массив по возрастанию начал отрезков
	sort.Slice(intervals, func(i, j int) bool { return intervals[i][0] < intervals[j][0] })

	result := make([][]int, 0, len(intervals))
	result = append(result, intervals[0])

	for _, interval := range intervals[1:] {
		if top := result[len(result)-1]; interval[0] > top[1] {
			result = append(result, interval)
		} else if interval[1] > top[1] {
			top[1] = interval[1]
		}
	}

	return result
}
