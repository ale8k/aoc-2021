package aoc9

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func AOC9P1(data *os.File) int {
	reader := bufio.NewReader(data)

	mapped := make([][]int, 0)
	for {
		points, err := reader.ReadString('\n')
		pointsToAppend := make([]int, 0)
		for _, v := range strings.Split(points, "") {
			k, err := strconv.Atoi(v)
			if err == nil {
				pointsToAppend = append(pointsToAppend, k)
			}
		}
		mapped = append(mapped, pointsToAppend)
		if err == io.EOF {
			break
		}
	}

	// 2199943210
	// 3987894921
	// 9856789892
	// 8767896789
	// 9899965678

	fmt.Println(mapped)
	riskLevel := make([]int, 0)
	// Go over every row
	for rowIndex, row := range mapped {
		// Now we loop over a row's element and compare
		for i, v := range row {
			if i == 0 {
				// dont try -1 check
			}

			if i == len(row)-1 {
				// dont try +1 check
			}

			if rowIndex == 0 {
				// dont try -1 check
			}

			if rowIndex == len(mapped)-1 {
				// dont try +1 check
			}

			if v < row[i-1] && v < row[i+1] && v < mapped[rowIndex-1][i] && v < mapped[rowIndex+1][i] {
				riskLevel = append(riskLevel, v)
			}

		}
	}
	fmt.Println(riskLevel)

	return 10
}
