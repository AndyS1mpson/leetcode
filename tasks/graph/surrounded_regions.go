package graph

/*
Задача:

	Дана матрица m x n содержащая буквы 'Х' и 'О', захватите области которые окружены по правилам:
		* Клетка соединяется с соседними клетками по горизонтали или вертикали
		* Чтобы сформировать регион, соедините все клетки с буквой 'О'
		* Регион окружен клетками 'X', если вы можете соединить регион с клетками 'X и ни одна из клеток региона не находится на краю доски
	Захват окруженного региона осуществляется заменой всех 'О' на 'Х' на входной матричной доске

Пример:

	Input: board = [
			["X","X","X","X"],
			["X","O","O","X"],
			["X","X","O","X"],
			["X","O","X","X"],
		]
	Output: [
		["X","X","X","X"],
		["X","X","X","X"],
		["X","X","X","X"],
		["X","O","X","X"],
	]

	Input: board = [
			["X","X","X"],
			["X","O","X"],
			["X","X","X"],
		]

	Output: board = [
			["X","X","X"],
			["X","X","X"],
			["X","X","X"],
		]

Сложность:
*/
func Solve(board [][]byte) {
	if len(board) == 0 {
		return
	}

	// Проходим по всем 'O' которые находятся по краям и такие регионы помечаем '*'
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if i == 0 || j == 0 || i == len(board)-1 || j == len(board[i])-1 {
				if board[i][j] == 'O' {
					solveDFS(board, i, j)
				}
			}
		}
	}

	// Все регионы которые состоят из '*' меняем на 'O', а все 'O' на 'X'
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board[i]); j++ {
			if board[i][j] == '*' {
				board[i][j] = 'O'
			} else if board[i][j] == 'O' {
				board[i][j] = 'X'
			}
		}
	}
}

func solveDFS(board [][]byte, i int, j int) {
	if i < 0 || j < 0 || i >= len(board) || j >= len(board[i]) || board[i][j] == 'X' || board[i][j] == '*' {
		return
	}

	board[i][j] = '*'
	solveDFS(board, i-1, j)
	solveDFS(board, i+1, j)
	solveDFS(board, i, j-1)
	solveDFS(board, i, j+1)

}
