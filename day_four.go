package main

import (
	"fmt"
	"strconv"
	"strings"
)

func split_by_spaces(string_to_split string) []int {
	var numbers = []int{}
	split_up_winning_numbers := strings.Split(string_to_split, " ")
	for _, i := range split_up_winning_numbers {
		str_to_num, err := strconv.Atoi(i)
		if err == nil {
			numbers = append(numbers, str_to_num)
		}
	}
	return numbers
}

func find_the_winners(lines []string) {
	final_number := 0
	for _, line := range lines {
		get_card_number := strings.Split(line, ":")
		card_number := get_card_number[0]

		split_up := strings.Split(line, "|")

		winning_numbers := split_by_spaces(split_up[0])
		playing_numbers := split_by_spaces(split_up[1])
		fmt.Print(card_number, "\n")
		points := 0
		for _, win := range winning_numbers {
			for _, play := range playing_numbers {
				if win == play {
					if points != 0 {
						points = points * 2
					} else {
						points = 1
					}

				}
			}
		}
		final_number = final_number + points
	}
	fmt.Print("part1 ", final_number)
}
