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

func place_text_into_array(file *bufio.Scanner) []string {
	file.Split(bufio.ScanLines)
	result := []string{}
	for file.Scan() {
		text := file.Text()
		result = append(result, text)
	}
	return result
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
		fileScanner := read_file("day_three_input.txt")
		text := place_text_into_array(fileScanner)
		everything := determine_part_and_symbol_placements(text)
		determine_if_a_real_part(everything)
	case "day4":
		fileScanner := read_file("day_four_input.txt")
		text := place_text_into_array(fileScanner)
		find_the_winners(text)
	}
}
