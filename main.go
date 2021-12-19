package main

import (
	"fmt"
)

// Welcome the player to the game and state the rules of the game
func welcome() {
	fmt.Println("Welcome to Tic Tac Toe!")
	fmt.Println("You will make your move by entering the row and column number of the space you want to move to.")
	fmt.Println("Your moves should follow the format 'row column'. For example, to place a mark on the space in the 2nd row, 3rd column, enter '2 3'.")
	fmt.Println("The first player to get three of their marks in a row, column, or diagonal wins!")
	fmt.Println("Good luck!")
}

// Initialize a new empty board as a 3x3 matrix
func initBoard() [][]string {
	return [][]string{
		{"_", "_", "_"},
		{"_", "_", "_"},
		{"_", "_", "_"},
	}
}

// Print the tic tac toe board to the console
func printBoard(board [][]string) {
	fmt.Println("")
	for i := 0; i < len(board); i++ {
		fmt.Printf("%s|%s|%s\n", board[i][0], board[i][1], board[i][2])
	}
}

// Get the player move in a valid format (row column)
func getMove(board [][]string) (int, int) {
	var row, col int
	fmt.Print("Enter your next move: ")
	fmt.Scanf("%d %d", &row, &col)
	return row - 1, col - 1
}

// Check if the move is valid
func isValidMove(board [][]string, row, col int) bool {
	if row < 0 || row > 2 || col < 0 || col > 2 {
		return false
	}
	if board[row][col] == "_" {
		return true
	}
	return false
}
// Make the move
func makeMove(board [][]string, row, col int, player string) [][]string {
	board[row][col] = player
	return board
}
// Check if the board is full
func isFull(board [][]string) bool {
	for i := 0; i < len(board); i++ {
		for j := 0; j < len(board); j++ {
			if board[i][j] == "_" {
				return false
			}
		}
	}
	return true
}
// Switch player
func switchPlayer(player string) string {
	if player == "X" {
		return "O"
	}
	return "X"
}

// Scan the board to see if a player has won
func checkWin(board [][]string, player string) bool {
	// Check rows
	for i := 0; i < len(board); i++ {
		if board[i][0] == player && board[i][1] == player && board[i][2] == player {
			return true
		}
	}
	// Check columns
	for i := 0; i < len(board); i++ {
		if board[0][i] == player && board[1][i] == player && board[2][i] == player {
			return true
		}
	}
	// Check diagonals
	if (board[0][0] == player && board[1][1] == player && board[2][2] == player) || (board[0][2] == player && board[1][1] == player && board[2][0] == player) {
		return true
	}
	return false
}

func main() {
	welcome()
	board := initBoard()
	player := "X"

	// Main game loop
	for {
		printBoard(board)
		row, col := getMove(board)

		if isValidMove(board, row, col) {
			board := makeMove(board, row, col, player)
			if checkWin(board, player) {
				fmt.Printf("\nCongratulations! Player %s wins!\n", player)
				break
			} else if isFull(board) {
				fmt.Println("It's a tie!")
				break
			}
		}

		player = switchPlayer(player)
	}
}