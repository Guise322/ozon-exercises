package main

import (
	"fmt"
)

func main() {
	fmt.Printf("%e\n", knightProbability(3, 2, 0, 0))
}

func knightProbability(n int, k int, row int, column int) float64 {
	if row > n || column > n {
		return 0.0
	}
	board := make([][][]*float64, n)
	for i := 0; i < n; i++ {
		board[i] = make([][]*float64, n)
		for j := 0; j < n; j++ {
			board[i][j] = make([]*float64, k+1)
		}
	}
	return evaluateMove(n, k, row, column, board)
}

func evaluateMove(n, k, row, column int, board [][][]*float64) float64 {
	if k == 0 && row < n && column < n && row >= 0 && column >= 0 {
		return 1.0
	}
	if row > n-1 || column > n-1 || row < 0 || column < 0 {
		return 0.0
	}
	d := board[row][column][k]
	if d == nil {
		res := evaluateMove(n, k-1, row+1, column+2, board)
		res += evaluateMove(n, k-1, row+1, column-2, board)
		res += evaluateMove(n, k-1, row+2, column+1, board)
		res += evaluateMove(n, k-1, row+2, column-1, board)
		res += evaluateMove(n, k-1, row-1, column+2, board)
		res += evaluateMove(n, k-1, row-1, column-2, board)
		res += evaluateMove(n, k-1, row-2, column+1, board)
		res += evaluateMove(n, k-1, row-2, column-1, board)
		res = res / 8.0
		board[row][column][k] = &res
		return res
	}
	return *d
}
