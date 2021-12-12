package aoc4

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"regexp"
	"strings"
	"time"
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

func AOC4P1(data *os.File) {
	nums, boards := getBingoData(data)
	winnerChannel := make(chan int)
	ctx, cancel := context.WithCancel(context.Background())

	makeChannelWithNums := func(nums []string) chan string {
		bingoChannel := make(chan string, len(nums))
		for _, v := range nums {
			bingoChannel <- v
		}
		return bingoChannel
	}

	boardPlayer := func(numbersRolledChan chan string, winnerChan chan int, board []string, winnerId int, endContext context.Context) {
		for {
			select {
			case <-endContext.Done(): // if cancel() execute
				return
			case newNumberRolled := <-numbersRolledChan:
				found := linearSearchString(newNumberRolled, board)
				if found != -1 {
					board[found] = ""
					won := calculateIfPlayerWon(board)
					if won {
						// TODO: Bug, context doesnt take precedence, we actually wanna check it
						// then default to new num rolled, not do an or as it'll never come in
						winnerChan <- winnerId
						// Now calc their score, we got num rolled when they won
						// and we know everything that isn't a ""
					}
				}
			}
		}
	}

	for i, v := range boards {

		time.Sleep(time.Microsecond)
		go boardPlayer(makeChannelWithNums(nums), winnerChannel, v, i, ctx)
	}

	// Wait for a winner
	winner := <-winnerChannel
	cancel()
	fmt.Println("Winner is: ", winner)
	fmt.Println("Winners score is: ", <-winnerChannel)

}

func linearSearchString(param string, slice []string) int {
	for i, v := range slice {
		if v == param {
			return i
		}
	}
	return -1
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
