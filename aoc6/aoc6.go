package aoc6

import (
	"fmt"
	"io"
	_ "net/http/pprof"
	"os"
	"strconv"
	"strings"
)

type lanternfish struct {
	// A fishy's time
	BirthTimer int
}

func AOC6P1(data *os.File) int {
	byteData, _ := io.ReadAll(data)
	k := strings.Split(string(byteData), ",")
	l := make([]int, len(k))
	for i, o := range k {
		l[i], _ = strconv.Atoi(o)
	}

	fishies := mapInputToLanternfish(l)

	for i := 0; i < 80; i++ {
		for _, fishy := range fishies {
			if fishy.BirthTimer != 0 {
				fishy.BirthTimer--
			} else {
				fishy.BirthTimer = 6
				fishies = append(fishies, &lanternfish{8})
			}
		}
	}

	return len(fishies)
}

// Creates lanternfish from the initial input
func mapInputToLanternfish(input []int) []*lanternfish {
	fishies := make([]*lanternfish, len(input), 915015336)
	for i, v := range input {
		fishies[i] = &lanternfish{v}
	}
	return fishies
}

func AOC6P2(data *os.File) int {
	byteData, _ := io.ReadAll(data)
	k := strings.Split(string(byteData), ",")
	l := make([]int, len(k), 26984457539)
	fmt.Println("length of l: ", len(l), "cal of l: ", cap(l))
	for i, o := range k {
		l[i], _ = strconv.Atoi(o)
	}

	fishies := l

	totalFishies := len(calculateBirthAmount(fishies, 256))
	return totalFishies

}

// Calcs amount of fishes based on provided logic
func calculateBirthAmount(fishies []int, days int) []int {
	for i := 1; i <= days; i++ {
		fmt.Println(i)
		runtime := len(fishies)
		newlyBornFishies := make([]int, 0)
		for j := 0; j < runtime; j++ {
			if fishies[j] == 0 {
				newlyBornFishies = append(newlyBornFishies, 8)
				fishies[j] = 6
			} else {
				fishies[j] = fishies[j] - 1
			}
		}
		fmt.Println("newly born...", len(newlyBornFishies))
		fishies = append(fishies, newlyBornFishies...)
	}

	return fishies
}

// After  1 day:  5
// After  2 days: 6
// After  3 days: 7
// After  4 days: 9
// After  5 days: 10
// After  6 days: 10
// After  7 days: 10
// After  8 days: 10
// After  9 days: 11
// After 10 days: 12
// After 11 days: 15
// 18 = 26
