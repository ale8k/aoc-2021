package aoc8

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// 0:      1:      2:      3:      4:
// aaaa    ....    aaaa    aaaa    ....
// b    c  .    c  .    c  .    c  b    c
// b    c  .    c  .    c  .    c  b    c
// ....    ....    dddd    dddd    dddd
// e    f  .    f  e    .  .    f  .    f
// e    f  .    f  e    .  .    f  .    f
// gggg    ....    gggg    gggg    ....

//  5:      6:      7:      8:      9:
// aaaa    aaaa    aaaa    aaaa    aaaa
// b    .  b    .  .    c  b    c  b    c
// b    .  b    .  .    c  b    c  b    c
// dddd    dddd    ....    dddd    dddd
// .    f  e    f  .    f  e    f  .    f
// .    f  e    f  .    f  e    f  .    f
// gggg    gggg    ....    gggg    gggg

func AOC8P1(data *os.File) int {
	// Break input example into *outputs* only:
	reader := bufio.NewReader(data)
	segmentsStrings := make([]string, 0)
	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		segmentsStrings = append(segmentsStrings, strings.Split(string(line), "|")[1])
	}

	// 1, 4, 7, or 8
	const (
		ONE   int = 2
		FOUR  int = 4
		SEVEN int = 3
		EIGHT int = 7
	)

	totalCount := 0

	for _, v := range segmentsStrings {
		// For every output, break it up
		separated := strings.Split(v, " ")
		// Check length of each, see if it matches 1/4/7/8
		for _, v2 := range separated {
			switch len(v2) {
			case ONE:
				fallthrough
			case FOUR:
				fallthrough
			case SEVEN:
				fallthrough
			case EIGHT:
				totalCount++
			}
		}
	}
	return totalCount

}

// So, the unique signal patterns would correspond to the following digits:
// acedgfb: 8
// cdfbe: 5
// gcdfa: 2
// fbcad: 3
// dab: 7
// cefabd: 9
// cdfgeb: 6
// eafb: 4
// cagedb: 0
// ab: 1

// Then, the four digits of the output value can be decoded:
// cdfeb: 5
// fcadb: 3
// cdfeb: 5
// cdbaf: 3

func AOC8P2(data *os.File) int {
	reader := bufio.NewReader(data)
	segmentsStrings := make([][]string, 0)
	for {
		line, _, err := reader.ReadLine()

		if err == io.EOF {
			break
		}

		segmentsStrings = append(segmentsStrings, strings.Split(string(line), "|"))
	}

	// Calculate 1/4/7/8 examples of each in each line
	// we then know the 'segments' for these (or at least the possible combinations)
	// so, calc segments that are used in 1/4/7/8
	// So basically we need a mapper from segment -> segment, then a parser for the mapped segments

	// func calcSegments(...) string[]string
	const (
		ONE_SIZE   int = 2 // unique C F
		SEVEN_SIZE int = 3 // unique A C F
		FOUR_SIZE  int = 4 // unique B C D F
		EIGHT_SIZE int = 7 // unique A B C D E F G - bit harder this one
	)

	one := ""
	seven := ""
	four := ""
	eight := ""

	k := segmentsStrings[0][0]
	stuff := strings.Split(k, " ")

	for _, j := range stuff {
		sizeOfNum := len(j)
		switch sizeOfNum {
		case ONE_SIZE:
			one = j
		case SEVEN_SIZE:
			seven = j
		case FOUR_SIZE:
			four = j
		case EIGHT_SIZE:
			eight = j
		}
	}
	// inner func, detect patterns between 1,7,4,8
	// so we compare to original a-g segments and give them integer position
	// i prefer explicit over iota, sorry
	const (
		A int = 0
		B int = 1
		C int = 2
		D int = 3
		E int = 4
		F int = 5
		G int = 6
	)
	// We also know original positions of 1,7,4,8 by 'index' in the provided crypto
	// One, 0 = C, 1 = F
	// Seven, 0 = A, 1 = C, 2 = F
	// Four, 0 = B, 1 = C, 2 = D, 3 = F
	// Eight 0-6 all in order
	// ONE := []int{C, F}
	// SEVEN := []int{A, C, F}
	// FOUR := []int{B, C, D, F}
	// EIGHT := []int{A, B, C, D, E, F, G}
	segmentMap1 := make(map[string]int)
	segmentMap2 := make(map[string]int)
	segmentMap3 := make(map[string]int)
	segmentMap4 := make(map[string]int)

	setSegmentsForOne := func(n string, segMap map[string]int) {
		segMap[string(n[0])] = C
		segMap[string(n[1])] = F
	}

	setSegmentsForSeven := func(n string, segMap map[string]int) {
		segMap[string(n[0])] = A
		segMap[string(n[1])] = C
		segMap[string(n[2])] = F
	}

	setSegmentsForFour := func(n string, segMap map[string]int) {
		segMap[string(n[0])] = B
		segMap[string(n[1])] = C
		segMap[string(n[2])] = D
		segMap[string(n[3])] = F
	}

	// If this works... then it basically solves all of them lol
	setSegmentsForEight := func(n string, segMap map[string]int) {
		segMap[string(n[0])] = A
		segMap[string(n[1])] = B
		segMap[string(n[2])] = C
		segMap[string(n[3])] = D
		segMap[string(n[4])] = E
		segMap[string(n[5])] = F
		segMap[string(n[6])] = G
	}
	fmt.Println(one, seven, four, eight)
	setSegmentsForOne(one, segmentMap1)
	setSegmentsForSeven(seven, segmentMap2)
	setSegmentsForFour(four, segmentMap3)
	setSegmentsForEight(eight, segmentMap4)
	fmt.Println(segmentMap1)
	fmt.Println(segmentMap2)
	fmt.Println(segmentMap3)
	fmt.Println(segmentMap4)

	return 'A'
}
