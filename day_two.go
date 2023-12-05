package main

import (
	"bufio"
	"fmt"
	"os"
)

func read_file() {
	readFile, err := os.Open("day_two_input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
}

// func main() {
// 	read_file()
// }
