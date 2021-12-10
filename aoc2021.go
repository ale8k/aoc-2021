package main

import (
	"fmt"
	"strconv"

	"github.com/ale8k/aoc2021/aoc1"
	"github.com/ale8k/aoc2021/aoc2"
	"github.com/ale8k/aoc2021/utils"
)

func main() {
	runAoc1p1()
	runAoc1p2()

	runAoc2p1()

}

func runAoc2p1() {
	lines := utils.GrabDataFromFile("./aoc2/aoc2input.txt")

	v := aoc2.AOC2P1(lines)
	fmt.Println(v)
}

func runAoc1p2() {
	lines := utils.GrabDataFromFile("./aoc1/aoc1input.txt")

	sonarData := make([]int, len(lines))
	for i, v := range lines {
		sonarData[i], _ = strconv.Atoi(v)
	}

	v := aoc1.AOC1P2(sonarData)
	fmt.Println(v)
}

func runAoc1p1() {
	lines := utils.GrabDataFromFile("./aoc1/aoc1input.txt")
	sonarData := make([]int, len(lines))

	for i, v := range lines {
		sonarData[i], _ = strconv.Atoi(v)
	}

	v := aoc1.AOC1P1(sonarData)
	fmt.Println(v)
}
