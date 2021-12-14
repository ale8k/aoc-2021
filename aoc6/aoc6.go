package aoc6

import (
	"fmt"
	"io"
	"net/http"
	_ "net/http/pprof"
	"os"
	"strconv"
	"strings"
)

// represents a lanternfish, contains their state
// but additionally tracks the school
type lanternfish struct {
	// A fishy's time
	BirthTimer int
}

func AOC6P1(data *os.File) int {
	go func() {
		fmt.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	byteData, _ := io.ReadAll(data)
	k := strings.Split(string(byteData), ",")
	l := make([]int, len(k))
	for i, o := range k {
		l[i], _ = strconv.Atoi(o)
	}

	fishies := mapInputToLanternfish(l)

	daysLeft := 256

	for i := 0; i < daysLeft; i++ {
		fmt.Printf("Day %d len %d\n", i, len(fishies))
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
	fishies := make([]*lanternfish, len(input))
	for i, v := range input {
		fishies[i] = &lanternfish{v}
	}
	return fishies
}
