/*
On an n x n chessboard, a knight starts at the cell (row, column) and attempts to make exactly k moves.
The rows and columns are 0-indexed, so the top-left cell is (0, 0), and the bottom-right cell is (n - 1, n - 1).

A chess knight has eight possible moves it can make, as illustrated below. Each move is two cells in
a cardinal direction, then one cell in an orthogonal direction.

Each time the knight is to move, it chooses one of eight possible moves uniformly at random (even if the piece
would go off the chessboard) and moves there.

The knight continues moving until it has made exactly k moves or has moved off the chessboard.

Return the probability that the knight remains on the board after it has stopped moving.
*/

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
