package aoc3

import (
	"testing"

	"github.com/ale8k/aoc2021/utils"
)

func TestAOC3P1(t *testing.T) {
	type args struct {
		data []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			"it runs",
			args{data: func() []string {
				return utils.GrabDataFromFile("./aoc3input.txt")
			}()},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AOC3P1(tt.args.data)
		})
	}
}
