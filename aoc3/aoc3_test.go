package aoc3

import (
	"reflect"
	"testing"

	"github.com/ale8k/aoc2021/utils"
)

func Test_parseBitMatrix(t *testing.T) {
	type args struct {
		data []string
	}
	tests := []struct {
		name string
		args args
		want [][]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := parseBitMatrix(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("parseBitMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getIterationCount(t *testing.T) {
	type args struct {
		matrix [][]string
	}
	tests := []struct {
		name    string
		args    args
		want    int
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := getIterationCount(tt.args.matrix)
			if (err != nil) != tt.wantErr {
				t.Errorf("getIterationCount() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("getIterationCount() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_flipBitArr(t *testing.T) {
	type args struct {
		bitArr []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := flipBitArr(tt.args.bitArr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flipBitArr() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAOC3P1(t *testing.T) {
	type args struct {
		data []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"it passes aoc3p1",
			args{data: func() []string {
				return utils.GrabDataFromFile("./aoc3input.txt")
			}()},
			2250414,
		},
		{
			"it passes aoc3p1 with examplar values",
			args{data: []string{
				"00100",
				"11110",
				"10110",
				"10111",
				"10101",
				"01111",
				"00111",
				"11100",
				"10000",
				"11001",
				"00010",
				"01010",
			}},
			198,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AOC3P1(tt.args.data); got != tt.want {
				t.Errorf("AOC3P2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAOC3P2(t *testing.T) {
	type args struct {
		data []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AOC3P2(tt.args.data); got != tt.want {
				t.Errorf("AOC3P2() = %v, want %v", got, tt.want)
			}
		})
	}
}
