package main

import (
	"fmt"
	"os"

	"github.com/ale8k/aoc2021/aoc6"
)

func main() {
	data, err := os.OpenFile("./aoc6/aoc6inputexample.txt", os.O_RDONLY, 0777)
	if err != nil {
		fmt.Println("didnt load file")
	}
	defer data.Close()
	fmt.Println(aoc6.AOC6P2(data))

}
