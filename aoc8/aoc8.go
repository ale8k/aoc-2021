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
	fmt.Println("total count is: ", totalCount)

	return totalCount

}
