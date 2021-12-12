package aoc4

import (
	"bufio"
	"context"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
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

func linearSearchString(param string, slice []string) int {
	for i, v := range slice {
		if v == param {
			return i
		}
	}
	return -1
}

func calculateIfPlayerWon(board []string) (bool, int) {
	calcNums := func(board []string) int {
		total := 0
		for _, v := range board {
			if v != "" {
				k, _ := strconv.Atoi(v)
				total += k
			}
		}
		return total
	}

	won := false
	total := 0
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
			total = calcNums(board)
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
			total = calcNums(board)
			break
		}
	}

	return won, total
}

type winner struct {
	id    int
	score int
}

var winnersTotal []winner = make([]winner, 0)

func AOC4P1(data *os.File) (winner, winner) {
	nums, boards := getBingoData(data)
	ctx, cancel := context.WithCancel(context.Background())

	boardPlayer := func(numbersRolledChan chan string, board []string, winnerId int, playerCtx context.Context) {
		for {
			select {
			case <-time.NewTimer(time.Second * 3).C:
				return
			case <-playerCtx.Done():
				return
			case newNumberRolled := <-numbersRolledChan:

				found := linearSearchString(newNumberRolled, board)
				if found != -1 {
					board[found] = ""
					won, total := calculateIfPlayerWon(board)
					if won {
						prevRolled, _ := strconv.Atoi(newNumberRolled)
						winnersTotal = append(winnersTotal, winner{winnerId, (prevRolled * total)})
						return
					}
				}
			}

		}
	}

	channelsList := make([]chan string, len(boards))
	// Create a channel for each of them
	fmt.Println("creating bingo channels")
	for i := range boards {
		channelsList[i] = make(chan string, len(nums))
	}

	fmt.Println("creating bingo boards go routines")
	for i, v := range boards {
		go boardPlayer(channelsList[i], v, i, ctx)
	}

	// Now we send the numbers to each player incrementally
	fmt.Println("sending nums to bingo channels")
	for _, num := range nums {
		for i := 0; i < len(boards)-10; i++ {
			channelsList[i] <- num
		}
		// Forces order, for purpose of test
		time.Sleep(time.Microsecond * 50)
	}
	cancel()
	fmt.Println("winner found")
	return winnersTotal[0], winnersTotal[len(winnersTotal)-1]
}
