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
		determine_players()
		fmt.Print("\n")
		determine_least_amount_of_cubes()
	case "day3":
		determine_part_numbers()
	}
}
