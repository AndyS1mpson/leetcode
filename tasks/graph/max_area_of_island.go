package graph

/*
Задача:

	Дана бинарная матрица m x n. Остров это группа из 1, соединенных 4 направлениями (горизонтально и вертикально).
	Верните максимальную площадь острова в сетке. Если острова нет, верните 0.

Пример:

	Input: grid = [
		[0,0,1,0,0,0,0,1,0,0,0,0,0],
		[0,0,0,0,0,0,0,1,1,1,0,0,0],
		[0,1,1,0,1,0,0,0,0,0,0,0,0],
		[0,1,0,0,1,1,0,0,1,0,1,0,0],
		[0,1,0,0,1,1,0,0,1,1,1,0,0],
		[0,0,0,0,0,0,0,0,0,0,1,0,0],
		[0,0,0,0,0,0,0,1,1,1,0,0,0],
		[0,0,0,0,0,0,0,1,1,0,0,0,0],
	]
	Output: 6
	Explanation: The answer is not 11, because the island must be connected 4-directionally.

Сложность: O(MN)
*/
func MaxAreaOfIsland(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	maxIslandLen := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				currIslandLen := maxAreadBFS(grid, i, j)
				maxIslandLen = max(maxIslandLen, currIslandLen)
			}
		}
	}

	return maxIslandLen
}

func maxAreadBFS(grid [][]int, r int, c int) int {
	// очередь для bfs
	q := [][2]int{{r, c}}
	grid[r][c] = 2
	islandLen := 1

	offsets := [4][2]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

	for len(q) > 0 {
		curr := q[0]
		q = q[1:]

		for _, offset := range offsets {
			x := curr[0] + offset[0]
			y := curr[1] + offset[1]

			if x >= 0 && x < len(grid) && y >= 0 && y < len(grid[x]) && grid[x][y] == 1 {
				islandLen++
				q = append(q, [2]int{x, y})
				grid[x][y] = 2
			}
		}
	}

	return islandLen
}
