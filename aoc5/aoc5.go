package aoc5

import (
	"bufio"
	"io"
	"os"
	"strconv"
	"strings"
)

// Example:
// .......1..
// ..1....1..
// ..1....1..
// .......1..
// .112111211
// ..........
// ..........
// ..........
// ..........
// 222111....

// Plots a bunch of coordinates onto a matrix
func AOC5P1(data *os.File) int {
	matrix, lineCoords := makeMatrixAndLines(data)

	for _, coord := range lineCoords {
		x := coord[0]
		y := coord[1]
		matrix[y][x]++
	}

	crossedOverLines := 0

	// Gets all co-ordinates of value 2 or greater
	for _, yAxis := range matrix {
		for _, coordinate := range yAxis {
			if coordinate > 1 {
				crossedOverLines++
			}
		}
	}

	return crossedOverLines
}

func makeMatrixAndLines(data *os.File) ([][]int, [][]int) {
	lines := getHorizontalLines(data)
	largestX := getLargestXY(lines, 0)
	largestY := getLargestXY(lines, 1)

	matrix := make([][]int, largestY+1)
	fillMatrix(matrix, largestX+1, largestY+1)
	return matrix, lines
}

// Filles a matrix of ints to default values of x/y size (presumes y constructed)
func fillMatrix(matrix [][]int, x int, y int) {
	for i := 0; i < y; i++ {
		k := make([]int, x)
		matrix[i] = k
	}
}

// Gets all co-ordinatates to plot the lines on a matrix
func getHorizontalLines(data *os.File) [][]int {
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
				for fromY <= toY {
					lines = append(lines, []int{fromX, fromY})
					fromY++
				}
			} else {
				for fromY >= toY {
					lines = append(lines, []int{fromX, fromY})
					fromY--
				}
			}
		} else if fromY == toY {
			if fromX < toX {
				for fromX <= toX {
					lines = append(lines, []int{fromX, fromY})
					fromX++
				}
			} else {
				for fromX >= toX {
					lines = append(lines, []int{fromX, fromY})
					fromX--
				}
			}
		}
	}
	return lines
}

// Gets argest x or y co-ord, 0 = x, 1 = y
func getLargestXY(slice [][]int, xOrY int) int {
	currentLargest := 0
	for _, v := range slice {
		if currentLargest < v[xOrY] {
			currentLargest = v[xOrY]
		}
	}
	return currentLargest
}

// Gets all co-ordinatates to plot the lines on a matrix
func getAllOverlappingLines(data *os.File) [][]int {
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

		// Think diagonal first
		// Else just treat it as a horizontal line, I think we can do this by a few conditions
		if fromX != toX && fromY != toY {
			// Now we do a check for if its a < > or > <, and -+ or +-
		}

		if fromX == toX {
			if fromY < toY {
				for fromY <= toY {
					lines = append(lines, []int{fromX, fromY})
					fromY++
				}
			} else {
				for fromY >= toY {
					lines = append(lines, []int{fromX, fromY})
					fromY--
				}
			}
		} else if fromY == toY {
			if fromX < toX {
				for fromX <= toX {
					lines = append(lines, []int{fromX, fromY})
					fromX++
				}
			} else {
				for fromX >= toX {
					lines = append(lines, []int{fromX, fromY})
					fromX--
				}
			}
		}
	}
	return lines
}
