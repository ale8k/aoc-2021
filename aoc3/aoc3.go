package aoc3

import (
	"errors"
	"strconv"
	"strings"
)

func parseBitMatrix(data []string) [][]string {
	bitMatrix := make([][]string, 0)
	for _, v := range data {
		bitLine := strings.Split(v, "")

		bitMatrix = append(bitMatrix, bitLine)
	}
	return bitMatrix
}

func getIterationCount(matrix [][]string) (int, error) {
	var bitLineIterations int
	if len(matrix[0]) > 0 {
		bitLineIterations = len(matrix[0])
	} else {
		return -1, errors.New("iteration count cannot be 0")
	}
	return bitLineIterations, nil
}

// Convenience to just flip the entire arr
func flipBitArr(bitArr []string) []string {
	copi := make([]string, len(bitArr))
	copy(copi, bitArr)

	for i, v := range copi {
		if v == "0" {
			copi[i] = "1"
		} else {
			copi[i] = "0"
		}
	}
	return copi
}

func AOC3P1(data []string) int {
	bitMatrix := parseBitMatrix(data)
	gammaMostCommon := make([]string, 0)
	iterationCount, err := getIterationCount(bitMatrix)

	if err != nil {
		return -1
	}

	for i := 0; i < iterationCount; i++ {
		row := make([]string, 0)

		for j := 0; j < len(bitMatrix); j++ {
			row = append(row, bitMatrix[j][i])
		}

		zeroes := 0
		ones := 0
		for _, val := range row {
			if val == "1" {
				ones += 1
			} else {
				zeroes += 1
			}
		}
		if ones > zeroes {
			gammaMostCommon = append(gammaMostCommon, "1")
		} else {
			gammaMostCommon = append(gammaMostCommon, "0")
		}

	}

	// Reverse did not add trailers in my original solution
	gam, _ := strconv.ParseInt(strings.Join(gammaMostCommon, ""), 2, 64)
	ep, _ := strconv.ParseInt(strings.Join(flipBitArr(gammaMostCommon), ""), 2, 64)
	return int(gam * ep)
}

// 00100
// 11110 -> 11110
// 10110 -> 10110 -> 10110 -> 10110 -> 10110
// 10111 -> 10111 -> 10111 -> 10111 -> 10111 -> 10111
// 10101 -> 10101 -> 10101 -> 10101
// 01111
// 00111
// 11100 -> 11100
// 10000 -> 10000 -> 10000
// 11001 -> 11001
// 00010
// 01010
// Go over each, find most/least common in row, filter rest, repeat
// In the fifth position,
// if there are an equal number of 0 bits and 1 bits (one each) keep 1 for most, 0 for least.
func AOC3P2(data []string) int {
	getOxyOrCo2 := func(which string, matrix [][]string, iterations int) int {
		for currentPosition := 0; currentPosition < iterations; currentPosition++ {
			zeroesList := make([][]string, 0)
			onesList := make([][]string, 0)

			for _, row := range matrix {
				if len(matrix) == 1 {
					break
				}
				// for least commin, use 0
				if row[currentPosition] == "1" {
					onesList = append(onesList, row)
				} else {
					zeroesList = append(zeroesList, row)
				}
			}

			if len(matrix) > 1 {
				if which == "co2" {
					if len(zeroesList) == len(onesList) {
						matrix = zeroesList
						// for least common, look for less than
					} else if len(zeroesList) < len(onesList) {
						matrix = zeroesList
					} else {
						matrix = onesList
					}
				} else {
					if len(zeroesList) == len(onesList) {
						matrix = onesList
						// for least common, look for less than
					} else if len(zeroesList) > len(onesList) {
						matrix = zeroesList
					} else {
						matrix = onesList
					}
				}
			}
		}
		k, _ := strconv.ParseInt(strings.Join(matrix[0], ""), 2, 64)
		return int(k)
	}

	oxyMatrix := parseBitMatrix(data)
	co2Matrix := make([][]string, len(oxyMatrix))
	copy(co2Matrix, oxyMatrix)
	iterationCount, err := getIterationCount(oxyMatrix)
	if err != nil {
		return -1
	}
	oxy := getOxyOrCo2("bob", oxyMatrix, iterationCount)
	co2 := getOxyOrCo2("co2", co2Matrix, iterationCount)
	return oxy * co2
}
