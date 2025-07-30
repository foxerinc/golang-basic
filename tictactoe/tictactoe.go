package main

import (
	"fmt"
	"strings"
)

var winningMoves = [][][]int{
	{{0, 0}, {0, 1}, {0, 2}},
	{{1, 0}, {1, 1}, {1, 2}},
	{{2, 0}, {2, 1}, {2, 2}},
	{{0, 0}, {1, 0}, {2, 0}},
	{{0, 1}, {1, 1}, {2, 1}},
	{{0, 2}, {1, 2}, {2, 2}},
	{{0, 0}, {1, 1}, {2, 2}},
	{{0, 2}, {1, 1}, {2, 0}},
}

func validateWin(playerMap [][]rune, player rune) bool {
	for _, line := range winningMoves {
		if playerMap[line[0][0]][line[0][1]] == player && playerMap[line[1][0]][line[1][1]] == player && playerMap[line[2][0]][line[2][1]] == player {
			return true
		}

	}

	return false
}

func validateInput(inp string) (string, int) {

	if len(inp) != 9 {
		return "Invalid Input: (Please input 9 character which represent X,O or -)", -1
	}
	countX := 0
	countO := 0
	countBlank := 0
	for _, ch := range inp {
		if ch != 'O' && ch != 'X' && ch != '-' {
			return "Invalid Input. Please input only X,O,or - character", -1
		}

		switch ch {
		case 'O':
			countO++
		case 'X':
			countX++
		default:
			countBlank++
		}
	}

	if countX-countO > 1 || countO-countX > 1 {
		return "Invalid Game. Please input fair move", -1
	}

	return "", countBlank

}

func main() {

	moveInput := "XXXOOOXXO"
	moveInput = strings.ToUpper(moveInput)

	err, countBlank := validateInput(moveInput)
	if err != "" {
		fmt.Println(err)
		return
	}

	moveSlice := [][]rune{
		[]rune(moveInput[0:3]),
		[]rune(moveInput[3:6]),
		[]rune(moveInput[6:9]),
	}

	isStillOnProcess := countBlank > 0

	if isStillOnProcess {
		fmt.Println("Game still in progress!")
	} else {
		x := validateWin(moveSlice, 'X')
		o := validateWin(moveSlice, 'O')

		if x && o {
			fmt.Println("Invalid game board")
		} else if x {
			fmt.Println("X wins!")
		} else if o {
			fmt.Println("O wins!")
		} else {
			fmt.Println("Its a draw!")
		}
	}

	fmt.Println("Current Board:")
	for _, row := range moveSlice {
		fmt.Printf(" %c | %c | %c\n", row[0], row[1], row[2])
	}
}