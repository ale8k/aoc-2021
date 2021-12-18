package aoc7

import (
	"fmt"
	"os"
	"testing"
)

func TestAOC7P1(t *testing.T) {
	data, err := os.OpenFile("./aoc7input.txt", os.O_RDONLY, 0777)
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
			name: "it passes p1",
			args: args{data: data},
			want: 349812,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AOC7P1(tt.args.data); got != tt.want {
				t.Errorf("AOC7P1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAOC7P2(t *testing.T) {
	data, err := os.OpenFile("./aoc7input.txt", os.O_RDONLY, 0777)
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
			name: "it passes p2",
			args: args{data: data},
			want: 99763899,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AOC7P2(tt.args.data); got != tt.want {
				t.Errorf("AOC7P2() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_calcFuel(t *testing.T) {
	type args struct {
		steps int
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			name: "n",
			args: args{steps: 1},
			want: 0,
		},
		{
			name: "n",
			args: args{steps: 2},
			want: 1,
		},
		{
			name: "n",
			args: args{steps: 5},
			want: 10,
		},
		{
			name: "n",
			args: args{steps: 10},
			want: 45,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calcFuel(tt.args.steps); got != tt.want {
				t.Errorf("calcFuel() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getHighest(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name:  "n",
			args:  args{slice: []int{5, 8, 2, 10, 1}},
			want:  3,
			want1: 10,
		},
		{
			name:  "n",
			args:  args{slice: []int{50, 8, 2, 10, 1}},
			want:  0,
			want1: 50,
		},
		{
			name:  "n",
			args:  args{slice: []int{5, 8, 2, 10, 1000}},
			want:  4,
			want1: 1000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getHighest(tt.args.slice)
			if got != tt.want {
				t.Errorf("getHighest() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getHighest() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func Test_getLowest(t *testing.T) {
	type args struct {
		slice []int
	}
	tests := []struct {
		name  string
		args  args
		want  int
		want1 int
	}{
		{
			name:  "n",
			args:  args{slice: []int{5, 8, 2, 10, 10}},
			want:  2,
			want1: 2,
		},
		{
			name:  "n",
			args:  args{slice: []int{1, 8, 2, 10, 1}},
			want:  0,
			want1: 1,
		},
		{
			name:  "n",
			args:  args{slice: []int{5, 8, 2, 10, 1}},
			want:  4,
			want1: 1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getLowest(tt.args.slice)
			if got != tt.want {
				t.Errorf("getLowest() got = %v, want %v", got, tt.want)
			}
			if got1 != tt.want1 {
				t.Errorf("getLowest() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}
