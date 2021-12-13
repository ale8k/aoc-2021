package aoc5

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// An entry like 1,1 -> 1,3 covers points 1,1, 1,2, and 1,3.
// An entry like 9,7 -> 7,7 covers points 9,7, 8,7, and 7,7.

func AOC5P1(data *os.File) {
	reader := bufio.NewReader(data)
	lines := make([][]int, 0)
	for {
		line, _, err := reader.ReadLine()
		if err == io.EOF {
			break
		}

		splitCoordinates := strings.Split(string(line), " -> ")
		fromXY := strings.Split(splitCoordinates[0], ",")
		toXY := strings.Split(splitCoordinates[1], ",")

		fromX, _ := strconv.Atoi(fromXY[0])
		toX, _ := strconv.Atoi(toXY[0])

		fromY, _ := strconv.Atoi(fromXY[1])
		toY, _ := strconv.Atoi(toXY[1])

		// Handle only horizontal for first part
		if fromX == toX {
			if fromY < toY {
				for fromY != toY {
					lines = append(lines, []int{fromX, fromY})
					fromY++
				}
			} else {
				for fromY != toY {
					lines = append(lines, []int{fromX, fromY})
					fromY--
				}
			}
		} else if fromY == toY {
			if fromX < toX {
				for fromX != toX {
					lines = append(lines, []int{fromX, fromY})
					fromX++
				}
			} else {
				for fromX != toX {
					lines = append(lines, []int{fromX, fromY})
					fromX--
				}

			}
		}

	}
	fmt.Println(lines)
}
