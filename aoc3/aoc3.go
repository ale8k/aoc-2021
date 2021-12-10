package aoc3

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func parseBitMatrix(data []string) [][]int {
	bitMatrix := make([][]int, 0)
	for _, v := range data {
		bitLine := strings.Split(v, "")
		bitLineAtoI := make([]int, 0, 10)
		for _, v := range bitLine {
			k, _ := strconv.Atoi(v)
			bitLineAtoI = append(bitLineAtoI, k)
		}
		bitMatrix = append(bitMatrix, bitLineAtoI)
	}
	return bitMatrix
}

func getIterationCount(matrix [][]int) int {
	var bitLineIterations int
	if len(matrix[0]) > 0 {
		bitLineIterations = len(matrix[0]) - 1
	} else {
		return -1
	}
	return bitLineIterations
}

// Parses bit array to int (ignores idea of whether it is short, etc, whatnot)
// just runs against each bits positional value
func parseBitArrayToInt(bitArr []int) int {
	finalInt := 0
	length := len(bitArr) - 1
	iterator := 0
	for i := length; i > 0; i-- {
		valueToAdd := math.Pow(2, float64(iterator))
		if bitArr[i] == 1 {
			finalInt += int(valueToAdd)
		}
		iterator++

	}
	return finalInt
}

// Convenience to just flip the entire arr
func flipBitArr(bitArr []int) []int {
	copi := make([]int, len(bitArr))
	copy(copi, bitArr)

	for i, v := range copi {
		if v == 0 {
			copi[i] = 1
		} else {
			copi[i] = 0
		}
	}
	return copi
}

func AOC3P1(data []string) int {
	bitMatrix := parseBitMatrix(data)
	gammaMostCommon := make([]int, 0)

	for i := 0; i < len(bitMatrix[0]); i++ {
		row := make([]int, 0)

		for j := 0; j < len(bitMatrix); j++ {
			row = append(row, bitMatrix[j][i])
		}

		zeroes := 0
		ones := 0
		for _, val := range row {
			if val == 1 {
				ones += 1
			} else {
				zeroes += 1
			}
		}

		if zeroes < ones {
			gammaMostCommon = append(gammaMostCommon, 1)
		} else {
			gammaMostCommon = append(gammaMostCommon, 0)
		}
	}
	gam := parseBitArrayToInt(gammaMostCommon) >> 1
	ep := parseBitArrayToInt((flipBitArr(gammaMostCommon))) >> 1
	fmt.Println(gam, ep)
	return 69
}
