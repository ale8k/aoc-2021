package aoc7

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"
)

func AOC7P1(data *os.File) int {
	positionsString, _ := csv.NewReader(data).ReadAll()
	positions := make([]int, 0, len(positionsString[0]))
	for _, pos := range positionsString[0] {
		num, _ := strconv.Atoi(pos)
		positions = append(positions, num)
	}

	leastDiffIndex := 0
	countOfDiffs := 0
	for i, val := range positions {
		fmt.Println("checking index:", i)
		currentCheck := 0
		for j := 0; j < len(positions); j++ {
			if val > positions[j] {
				currentCheck += val - positions[j]
			} else if val < positions[j] {
				currentCheck += positions[j] - val
			}
			if i == j {
				continue
			}
		}
		if countOfDiffs == 0 {
			countOfDiffs = currentCheck
		}
		fmt.Printf("current check: %v \n \n", currentCheck)
		if countOfDiffs > currentCheck {
			countOfDiffs = currentCheck
			leastDiffIndex = i
		}
	}
	fmt.Println(leastDiffIndex, countOfDiffs)

	return countOfDiffs
}

func AOC7P2(data *os.File) int {
	positionsString, _ := csv.NewReader(data).ReadAll()
	positions := make([]int, 0, len(positionsString[0]))
	for _, pos := range positionsString[0] {
		num, _ := strconv.Atoi(pos)
		positions = append(positions, num)
	}

	leastDiffIndex := 0
	countOfDiffs := 0
	for i, val := range positions {
		fmt.Println("checking index:", i)
		currentCheck := 0
		for j := 0; j < len(positions); j++ {
			if val > positions[j] {
				currentCheck += val - positions[j]
			} else if val < positions[j] {
				currentCheck += positions[j] - val
			}
			if i == j {
				continue
			}
		}
		if countOfDiffs == 0 {
			countOfDiffs = currentCheck
		}
		fmt.Printf("current check: %v \n \n", currentCheck)
		if countOfDiffs > currentCheck {
			countOfDiffs = currentCheck
			leastDiffIndex = i
		}
	}
	fmt.Println(leastDiffIndex, countOfDiffs)

	return countOfDiffs
}
