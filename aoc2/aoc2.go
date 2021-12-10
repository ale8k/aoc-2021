package aoc2

import (
	"strconv"
	"strings"
)

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

func AOC2P2(data []string) int {
	const (
		FORWARD string = "forward"
		UP      string = "up"
		DOWN    string = "down"
	)
	horizontal := 0
	depth := 0
	aim := 0

	for _, v := range data {
		split := strings.Split(v, " ")
		switch split[0] {
		case FORWARD:
			k, _ := strconv.Atoi(split[1])
			horizontal += k
			depth += aim * k
		case UP:
			k, _ := strconv.Atoi(split[1])
			aim -= k
		case DOWN:
			k, _ := strconv.Atoi(split[1])
			aim += k
		}
	}
	return horizontal * depth
}
