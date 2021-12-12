package aoc4

import (
	"fmt"
	"os"
	"reflect"
	"testing"
)

func Test_calculateIfPlayerWon(t *testing.T) {
	type args struct {
		board []string
	}
	tests := []struct {
		name  string
		args  args
		want  bool
		want2 int
	}{
		{
			name: "it calculates the player has won correctly vertically",
			args: args{board: []string{
				"0", "", "2", "3", "4",
				"5", "", "7", "8", "9",
				"10", "", "12", "13", "14",
				"15", "", "17", "18", "19",
				"20", "", "22", "24", "25",
			}},
			want:  true,
			want2: 10,
		},
		{
			name: "it calculates the player has won correctly horizontally",
			args: args{board: []string{
				"0", "1", "2", "3", "4",
				"", "", "", "", "",
				"10", "11", "12", "13", "14",
				"15", "16", "17", "18", "19",
				"20", "17", "22", "24", "25",
			}},
			want:  true,
			want2: 10,
		},
		{
			name: "it calculates the player has not won by 1 value",
			args: args{board: []string{
				"0", "1", "2", "3", "4",
				"", "20", "", "", "",
				"10", "11", "12", "13", "14",
				"15", "16", "17", "18", "19",
				"20", "17", "22", "24", "25",
			}},
			want:  false,
			want2: 10,
		},
		{
			name: "it calculates the player has not won by all values",
			args: args{board: []string{
				"0", "1", "2", "3", "4",
				"99", "99", "99", "99", "99",
				"10", "11", "12", "13", "14",
				"15", "16", "17", "18", "19",
				"20", "17", "22", "24", "25",
			}},
			want:  false,
			want2: 10,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got, got2 := calculateIfPlayerWon(tt.args.board); got != tt.want && got2 != tt.want2 {
				t.Errorf("calculateIfPlayerWon() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_linearSearchString(t *testing.T) {
	type args struct {
		param string
		slice []string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"it searches for the param correctly",
			args{
				param: "10",
				slice: []string{"1", "2", "3", "10", "20", "30"},
			},
			3,
		},
		{
			"it searches for the param correctly and it does not exist, so returns -1",
			args{
				param: "100",
				slice: []string{"1", "2", "3", "10", "20", "30"},
			},
			-1,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := linearSearchString(tt.args.param, tt.args.slice); got != tt.want {
				t.Errorf("linearSearchString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_getBingoData(t *testing.T) {
	type args struct {
		data *os.File
	}
	tests := []struct {
		name  string
		args  args
		want  []string
		want1 [][]string
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, got1 := getBingoData(tt.args.data)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getBingoData() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("getBingoData() got1 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestAOC4P1(t *testing.T) {

	data, err := os.OpenFile("./aoc4input.txt", os.O_RDONLY, 0777)
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
		want winner
	}{
		{
			name: "it gets the correct initial bingo winner",
			args: args{data: data},
			want: winner{id: 40, score: 87456},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println("test running fine i think ")
			if got, _ := AOC4P1(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AOC4P1() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAOC4P2(t *testing.T) {

	data, err := os.OpenFile("./aoc4input.txt", os.O_RDONLY, 0777)
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
		want winner
	}{
		{
			name: "it gets the correct initial bingo winner",
			args: args{data: data},
			want: winner{id: 84, score: 15561},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			fmt.Println("test running fine i think ")
			if _, got := AOC4P1(tt.args.data); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("AOC4P1() = %v, want %v", got, tt.want)
			}
		})
	}
}
