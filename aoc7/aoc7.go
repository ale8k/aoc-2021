package aoc7

import (
	"encoding/csv"
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

	return 10
}
