package aoc4

import (
	"bufio"
	"io"
	"os"
	"regexp"
	"strings"
)

func getBingoData(data *os.File) ([]string, [][]string) {
	reader := bufio.NewReader(data)
	line, _, _ := reader.ReadLine()

	bingoNums := strings.Split(string(line), ",")

	bingoBoards := make([][]string, 0)

	for {
		var err error
		board := make([]string, 0)
		for {
			line, _, err = reader.ReadLine()

			if len(line) == 0 && len(board) > 0 {
				bingoBoards = append(bingoBoards, board)
				break
			}
			replaceSpacesRegex := regexp.MustCompile(`(\s\s+|\s+)`)
			includesEmpties := strings.Split(replaceSpacesRegex.ReplaceAllString(string(line), " "), " ")
			noEmpties := make([]string, 0, len(includesEmpties))
			for _, item := range includesEmpties {
				if item != "" {
					noEmpties = append(noEmpties, item)
				}
			}
			board = append(board, noEmpties...)

		}
		if err == io.EOF {
			break
		}
	}

	return bingoNums, bingoBoards
}

// This challenge fits go PERFECTLY!
func AOC4P1(data *os.File) {
	_, boards := getBingoData(data)

	bingoChannel := make(chan string, len(boards))
	winnerChannel := make(chan string)

	boardPlayer := func(numbersRolledChan chan string, winnerChan chan string) {

	}

}

func calculateIfPlayerWon(board []string) bool {
	won := false
	// Run vertical check
	for i := 0; i < 5; i++ {
		winCounter := 0
		for j := 0; j <= 20; j += 5 {
			if board[i+j] == "" {
				winCounter++
			}
		}
		if winCounter == 5 {
			won = true
			break
		}
	}

	// Run horizontal check
	for i := 0; i <= 20; i += 5 {
		winCounter := 0
		for j := 0; j < 5; j++ {
			if board[i+j] == "" {
				winCounter++
			}
		}
		if winCounter == 5 {
			won = true
			break
		}
	}

	return won
}
