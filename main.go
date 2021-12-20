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
	fmt.Scanf("%d %d\r", &row, &col)
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

// Computer move
func computerMove(board [][] string) (int, int) {
	chkBoard := board
	// Check for winning or blocking move
	for i := 0; i < len(chkBoard); i++ {
		for j := 0; j < len(chkBoard[i]); j++ {
			// If the space is empty, check if it's a winning move
			if chkBoard[i][j] == "_" {
				// Assign the space to the computer
				chkBoard[i][j] = "O"
				// Check if the computer has won
				if checkWin(chkBoard, "O") {
					// If the computer has won, return the move
					return i, j
				}
				// If the computer has not won, assign the space to the player
				chkBoard[i][j] = "X"
				// Check if the player wins
				if checkWin(chkBoard, "X") {
					// If the wins, return the move to block the player
					return i, j
				}
				// If the player does not win, remove the space
				chkBoard[i][j] = "_"
			}
		}
	}
	// Check for corner move
	if chkBoard[0][0] == "_" {
		return 0, 0
	}
	if chkBoard[0][2] == "_" {
		return 0, 2
	}
	if chkBoard[2][0] == "_" {
		return 2, 0
	}
	if chkBoard[2][2] == "_" {
		return 2, 2
	}
	// Check for center move
	if chkBoard[1][1] == "_" {
		return 1, 1
	}
	// Return first possible move other than the corners and the center
	for i := 0; i < len(chkBoard); i++ {
		for j := 0; j < len(chkBoard[i]); j++ {
			if chkBoard[i][j] == "_" {
				return i, j
			}
		}
	}
	
	// If no possible move (shouldn't get to this point, but compiler complains) return an invalid move (-1, -1)
	return -1, -1
}

// Main loop for Player vs Player
func pvp() {
	board := initBoard()
	player := "X"
	for {
		printBoard(board)
		row, col := getMove(board)
		if isValidMove(board, row, col) {
			board = makeMove(board, row, col, player)
		} else {
			fmt.Println("Invalid move, try again.")
			continue
		}
		if checkWin(board, player) {
			printBoard(board)
			fmt.Printf("%s wins!\n", player)
			break
		} else if isFull(board) {
			printBoard(board)
			fmt.Println("It's a tie!")
			break
		}
		player = switchPlayer(player)
	}
	printBoard(board)
}

// Game loop for Player vs Computer
func pvc() {
	board := initBoard()
	player := "X"
	turn := player

	for {
		printBoard(board)
		if turn == player {
			row, col := getMove(board)
			if isValidMove(board, row, col) {
				board = makeMove(board, row, col, player)
			} else {
				fmt.Println("Invalid move, try again.")
				continue
			}
		} else {
			row, col := computerMove(board)
			board = makeMove(board, row, col, player)
		}
		if checkWin(board, player) {
			printBoard(board)
			fmt.Printf("%s wins!\n", player)
			break
		} else if isFull(board) {
			printBoard(board)
			fmt.Println("It's a tie!")
			break
		}
		player = switchPlayer(player)
	}
	printBoard(board)
}

func main() {
	welcome()

	fmt.Println("1. Player vs Player")
	fmt.Println("2. Player vs Computer")
	fmt.Println("3. Quit")
	fmt.Print("Enter your choice: ")
	var choice int
	fmt.Scanf("%d\r", &choice)

	switch choice {
		case 1:
			pvp()
		case 2:
			pvc()
		case 3:
			fmt.Println("Goodbye!")
		default:
			fmt.Println("Invalid choice, try again.")
	}
}