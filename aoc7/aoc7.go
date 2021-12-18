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

	_, highest := getHighest(positions)
	_, lowest := getLowest(positions)

	currentLowestFuelRate := 0
	for i := lowest; i < highest; i++ {
		current := 0

		for j := 0; j < len(positions); j++ {
			diff := 0
			if positions[j] > i {
				diff = calcFuel((positions[j] - i) + 1)
			} else if i > positions[j] {
				diff = calcFuel((i - positions[j]) + 1)
			}

			current += diff
		}

		if i == 0 {
			currentLowestFuelRate = current
		}
		if current < currentLowestFuelRate {
			currentLowestFuelRate = current

		}
	}

	return currentLowestFuelRate
}

// Based on the steps, calcs the fuel incrementally up to the step
func calcFuel(steps int) int {
	finalFuel := make([]int, 0)
	for i := 0; i < steps; i++ {
		finalFuel = append(finalFuel, i)
	}
	k := 0
	for _, v := range finalFuel {
		k += v
	}
	return k
}

func getHighest(slice []int) (int, int) {
	currentHighestIndex := 0
	currentHighest := 0
	for i, v := range slice {
		if v > currentHighest {
			currentHighest = v
			currentHighestIndex = i
		}
	}
	return currentHighestIndex, currentHighest
}

func getLowest(slice []int) (int, int) {
	currentLowestIndex := 0
	currentLowest := slice[0]
	for i, v := range slice {
		if v < currentLowest {
			currentLowest = v
			currentLowestIndex = i
		}
	}
	return currentLowestIndex, currentLowest
}
