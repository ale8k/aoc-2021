package aoc2

import (
	"strconv"
	"strings"
)

// AOC Day 2 challenge part 1
func AOC2P1(data []string) int {
	const (
		FORWARD string = "forward"
		UP      string = "up"
		DOWN    string = "down"
	)
	horizontal := 0
	depth := 0

	for _, v := range data {
		split := strings.Split(v, " ")
		switch split[0] {
		case FORWARD:
			k, _ := strconv.Atoi(split[1])
			horizontal += k
		case UP:
			k, _ := strconv.Atoi(split[1])
			depth -= k
		case DOWN:
			k, _ := strconv.Atoi(split[1])
			depth += k
		}
	}
	return horizontal * depth
}
