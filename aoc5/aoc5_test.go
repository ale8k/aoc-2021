package aoc5

import (
	"fmt"
	"os"
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

func TestAOC5P2(t *testing.T) {
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
			name: "it passes the given input matrix plotting task for part 2",
			args: args{data: data},
			want: 19676,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AOC5P2(tt.args.data); got != tt.want {
				t.Errorf("AOC5P2() = %v, want %v", got, tt.want)
			}
		})
	}
}
