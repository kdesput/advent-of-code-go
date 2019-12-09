package main

import (
	"bufio"
	"fmt"
	"os"
)

func readNumbersFromLine(filepath string) []int {
	file, err := os.Open(filepath)

	if err != nil {
		fmt.Println("ERROR!")
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	return []int{1, 2}
}
