package graph

/*
Задача:

	Дана матрица row x col где grid[i][j] = 1 - земля, 0 - вода
	Ячейки сетки соединены по горизонтали/вертикали (не по диагонали).
	Сетка полностью окружена водой, и на ней есть ровно один остров.

	Определите периметр острова

Пример:

	Input: grid = [
		[0,1,0,0],
		[1,1,1,0],
		[0,1,0,0],
		[1,1,0,0],
	]
	Output: 16

Сложность:
*/
func IslandPerimeter(grid [][]int) int {
	if len(grid) == 0 {
		return 0
	}

	perimeter := 0

	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[i]); j++ {
			if grid[i][j] == 1 {
				perimeter += calculateSurroundingWaterArea(i, j, grid)
			}
		}
	}

	return perimeter
}

func calculateSurroundingWaterArea(i, j int, grid [][]int) int {
	totalBeachAreaofCube := 0
	// checking 4 sides
	// up
	if i-1 < 0 || grid[i-1][j] == 0 {
		totalBeachAreaofCube++
	}
	// down
	if i+1 >= len(grid) || grid[i+1][j] == 0 {
		totalBeachAreaofCube++
	}
	// left
	if j-1 < 0 || grid[i][j-1] == 0 {
		totalBeachAreaofCube++
	}
	// right
	if j+1 >= len(grid[0]) || grid[i][j+1] == 0 {
		totalBeachAreaofCube++
	}
	return totalBeachAreaofCube
}
