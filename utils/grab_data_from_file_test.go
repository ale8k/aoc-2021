package utils

import (
	"reflect"
	"testing"
)

func TestGrabDataFromFile(t *testing.T) {
	type args struct {
		path string
	}
	tests := []struct {
		name string
		args args
		want int
	}{
		{
			"it grabs file correctly",
			args{path: "../aoc1/aoc1input.txt"},
			2000,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := GrabDataFromFile(tt.args.path); !reflect.DeepEqual(len(got), tt.want) {
				t.Errorf("GrabDataFromFile() = %v, want %v", len(got), tt.want)
			}
		})
	}
}
