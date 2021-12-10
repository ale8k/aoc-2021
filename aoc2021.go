package main

import (
	"fmt"

	"github.com/ale8k/aoc2021/aoc3"
	"github.com/ale8k/aoc2021/utils"
)

func main() {
	// 00100
	// 11110
	// 10110
	bob := []string{
		"00100",
		"11110",
		"10110",
	}
	fmt.Println(aoc3.AOC3P1(bob))

	fmt.Println(aoc3.AOC3P1(utils.GrabDataFromFile("./aoc3/aoc3input.txt")))

}
