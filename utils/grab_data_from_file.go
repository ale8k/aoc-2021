package utils

import (
	"io"
	"os"
	"strings"
)

func GrabDataFromFile(path string) []string {
	file, err := os.Open(path)
	handleErr(err)
	data, err := io.ReadAll(file)
	handleErr(err)
	lines := strings.Split(string(data), "\n")
	return lines
}
