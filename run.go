package main

import (
	"bufio"
	"fmt"
	"os"
)

func read_file(txt_file_name string) *bufio.Scanner {
	readFile, err := os.Open(txt_file_name)

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	return fileScanner
}

func main() {
	switch os.Args[1] {
	case "day1":
		find_digits_in_lines()
		fmt.Print("\n")
		find_anything_in_lines()
	case "day2":
		parse_each_game()
	}
}
