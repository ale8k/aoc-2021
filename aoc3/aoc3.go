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

func AOC3P2(data []string) int {
	return 10
}
