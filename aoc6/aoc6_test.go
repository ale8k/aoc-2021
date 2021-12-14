package aoc6

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func TestAOC6P1(t *testing.T) {
	data, err := os.OpenFile("./aoc6input.txt", os.O_RDONLY, 0777)
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
			"it passes aoc6 p1 test",
			args{data: data},
			374994,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := AOC6P1(tt.args.data); got != tt.want {
				t.Errorf("AOC6P1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_mapInputToLanternfish(t *testing.T) {
	type args struct {
		input []int
	}
	tests := []struct {
		name string
		args args
		want []*lanternfish
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := mapInputToLanternfish(tt.args.input); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mapInputToLanternfish() = %v, want %v", got, tt.want)
			}
		})
	}
}
