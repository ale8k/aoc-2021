package aoc1

import (
	"strconv"
	"testing"

	"github.com/ale8k/aoc2021/utils"
)

func TestAOC1P1(t *testing.T) {
	type args struct {
		data []int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"it passes aoc test",
			args{data: func() []int {
				lines := utils.GrabDataFromFile("./aoc1input.txt")

				sonarData := make([]int, len(lines))
				for i, v := range lines {
					sonarData[i], _ = strconv.Atoi(v)
				}
				return sonarData
			}()},
			1529,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AOC1P1(tt.args.data); got != tt.want {
				t.Errorf("AOC1P1() = %v, want %v", got, tt.want)
			}
		})
	}
}
