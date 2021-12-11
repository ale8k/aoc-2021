package aoc4

import (
	"os"
	"reflect"
	"testing"
)

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
	type args struct {
		data *os.File
	}
	tests := []struct {
		name string
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			AOC4P1(tt.args.data)
		})
	}
}

func Test_calculateIfPlayerWon(t *testing.T) {
	type args struct {
		board []string
	}
	tests := []struct {
		name string
		args args
		want bool
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
			want: true,
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
			want: true,
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
			want: false,
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
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := calculateIfPlayerWon(tt.args.board); got != tt.want {
				t.Errorf("calculateIfPlayerWon() = %v, want %v", got, tt.want)
			}
		})
	}
}
