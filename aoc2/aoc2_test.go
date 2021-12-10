package aoc2

import (
	"testing"

	"github.com/ale8k/aoc2021/utils"
)

func TestAOC2P1(t *testing.T) {
	type args struct {
		data []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"it passes day 2p1",
			args{data: func() []string {
				return utils.GrabDataFromFile("./aoc2input.txt")
			}(),
			},
			2215080,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AOC2P1(tt.args.data); got != tt.want {
				t.Errorf("AOC2P1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAOC2P2(t *testing.T) {
	type args struct {
		data []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"it passes day 2p2",
			args{data: func() []string {
				return utils.GrabDataFromFile("./aoc2input.txt")
			}(),
			},
			1864715580,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AOC2P2(tt.args.data); got != tt.want {
				t.Errorf("AOC2P2() = %v, want %v", got, tt.want)
			}
		})
	}
}
