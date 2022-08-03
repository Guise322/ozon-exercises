/*
Given an m x n grid of characters board and a string word, return true if word exists in the grid.

The word can be constructed from letters of sequentially adjacent cells, where adjacent cells are horizontally or vertically
neighboring. The same letter cell may not be used more than once.

Input: board = [["A","B","C","E"],["S","F","C","S"],["A","D","E","E"]], word = "ABCCED"
Output: true
*/

package main

import "fmt"

func main() {
	word := "ABC"
	board := [][]string{
		{"A", "F", "G"},
		{"B", "F", "G"},
		{"C", "F", "G"},
	}
	bBoard := [][]byte{}
	for _, row := range board {
		bRow := []byte{}
		for _, r := range row {
			bRow = append(bRow, byte([]rune(r)[0]))
		}
		bBoard = append(bBoard, bRow)
	}
	fmt.Printf("%t\n", exist(bBoard, word))
}

func exist(board [][]byte, word string) bool {
	bWord := []byte(word)
	if len(bWord) > len(board[0])*len(board) {
		return false
	}
	for i := range board {
		for j := range board[0] {
			if ex(board, bWord, i, j) {
				return true
			}
		}
	}
	return false
}

func ex(board [][]byte, word []byte, i, j int) bool {
	colLen := len(board)
	rowLen := len(board[0])
	if i < 0 || j < 0 || colLen <= i || rowLen <= j {
		return false
	}
	currRune := board[i][j]
	if len(word) == 1 && currRune == word[0] {
		return true
	}
	if currRune != word[0] {
		return false
	}
	board[i][j] = '$'
	res := ex(board, word[1:], i+1, j) ||
		ex(board, word[1:], i-1, j) ||
		ex(board, word[1:], i, j+1) ||
		ex(board, word[1:], i, j-1)
	board[i][j] = currRune
	return res
}
