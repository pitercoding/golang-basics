package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	board := [9]string{
		"1", "2", "3",
		"4", "5", "6",
		"7", "8", "9",
	}

	currentPlayer := "X"

	scanner := bufio.NewScanner(os.Stdin)

	for {
		printBoard(board)

		fmt.Printf("\nPlayer %s - choose a position (1-9): ", currentPlayer)

		scanner.Scan()

		input := strings.TrimSpace(scanner.Text())

		position, err := strconv.Atoi(input)

		if err != nil || position < 1 || position > 9 {
			fmt.Println("Invalid position.")
			continue
		}

		index := position - 1

		if board[index] == "X" || board[index] == "O" {
			fmt.Println("Position already occupied.")
			continue
		}

		board[index] = currentPlayer

		if checkWinner(board) {
			printBoard(board)
			fmt.Printf("\nPlayer %s wins!\n", currentPlayer)
			break
		}

		if isDraw(board) {
			printBoard(board)
			fmt.Println("\nDraw!")
			break
		}

		if currentPlayer == "X" {
			currentPlayer = "O"
		} else {
			currentPlayer = "X"
		}
	}
}

func printBoard(board [9]string) {
	fmt.Println()

	fmt.Printf(" %s | %s | %s\n", board[0], board[1], board[2])
	fmt.Println("---+---+---")
	fmt.Printf(" %s | %s | %s\n", board[3], board[4], board[5])
	fmt.Println("---+---+---")
	fmt.Printf(" %s | %s | %s\n", board[6], board[7], board[8])
}

func checkWinner(board [9]string) bool {
	combinations := [][]int{
		{0, 1, 2},
		{3, 4, 5},
		{6, 7, 8},
		{0, 3, 6},
		{1, 4, 7},
		{2, 5, 8},
		{0, 4, 8},
		{2, 4, 6},
	}

	for _, combo := range combinations {
		a, b, c := combo[0], combo[1], combo[2]

		if board[a] == board[b] &&
			board[b] == board[c] {
			return true
		}
	}

	return false
}

func isDraw(board [9]string) bool {
	for _, value := range board {
		if value != "X" && value != "O" {
			return false
		}
	}

	return true
}
