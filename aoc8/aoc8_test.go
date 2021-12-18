package aoc8

import (
	"fmt"
	"os"
	"testing"
)

func TestAOC8P1(t *testing.T) {
	data, err := os.OpenFile("./aoc8input.txt", os.O_RDONLY, 0777)
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
			"it passes day 8 part 1",
			args{data: data},
			532,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AOC8P1(tt.args.data); got != tt.want {
				t.Errorf("AOC8P1() = %v, want %v", got, tt.want)
			}
		})
	}
}
