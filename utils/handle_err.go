package utils

import "fmt"

func handleErr(err error) {
	if err != nil {
		fmt.Printf("error: %v", err)
	}
}
