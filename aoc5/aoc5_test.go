package aoc5

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestAOC5P1(t *testing.T) {
	data, err := os.OpenFile("./aoc5input.txt", os.O_RDONLY, 0777)
	if err != nil {
		fmt.Println("didnt load file")
	}
	defer data.Close()
	type args struct {
		data *os.File
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "it passes the given input matrix plotting task",
			args: args{data: data},
			want: 7414,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AOC5P1(tt.args.data); got != tt.want {
				t.Errorf("AOC5P1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_makeMatrixAndLines(t *testing.T) {
	type args struct {
		data *os.File
	}
	tests := []struct {
		name  string
		args  args
		want  [][]int
		want1 [][]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := makeMatrixAndLines(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("makeMatrixAndLines() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("makeMatrixAndLines() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_fillMatrix(t *testing.T) {
	type args struct {
		matrix [][]int
		x      int
		y      int
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fillMatrix(tt.args.matrix, tt.args.x, tt.args.y)
		})
	}
}

func Test_getHorizontalLines(t *testing.T) {
	type args struct {
		data *os.File
	}
	tests := []struct {
		name string
		args args
		want [][]int
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := getHorizontalLines(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getHorizontalLines() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getLargestXY(t *testing.T) {
	type args struct {
		slice [][]int
		xOrY  int
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
			if got := getLargestXY(tt.args.slice, tt.args.xOrY); got != tt.want {
				t.Errorf("getLargestXY() = %v, want %v", got, tt.want)
			}
		})
	}
}
