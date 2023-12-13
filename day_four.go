package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Match struct {
	game_id           int
	amount_of_matches int
	extra_copies      int
}
type Matches []Match

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
	all_matches := Matches{}
	for _, line := range lines {
		get_card_number := strings.Split(line, ":")
		card_number_arr := strings.Split(get_card_number[0], " ")
		card_number := card_number_arr[len(card_number_arr)-1]
		game_id, _ := strconv.Atoi(card_number)

		split_up := strings.Split(line, "|")

		winning_numbers := split_by_spaces(split_up[0])
		playing_numbers := split_by_spaces(split_up[1])

		points := 0
		amount_of_matches := 0
		for _, win := range winning_numbers {
			for _, play := range playing_numbers {
				if win == play {
					amount_of_matches += 1
					if points != 0 {
						points = points * 2
					} else {
						points = 1
					}

				}
			}
		}
		line_match := Match{
			amount_of_matches: amount_of_matches,
			game_id:           game_id,
			extra_copies:      0,
		}
		all_matches = append(all_matches, line_match)
		final_number = final_number + points
	}

	// per the # of matches in each Match struct, those need to be added to each consecutive game id's match
	// each struct
	for index := 1; index <= len(all_matches); index++ {
		if index == len(all_matches) {
			break
		}

		if all_matches[index-1].game_id != 1 {
			for i := 0; i < all_matches[index-1].extra_copies+1; i++ {
				for num := 1; num <= all_matches[index-1].amount_of_matches; num++ {
					if (index-1)+num >= 196 {
						continue
					}
					all_matches[(index-1)+num].extra_copies = all_matches[(index-1)+num].extra_copies + 1
				}

			}
		} else {
			for num := 1; num <= all_matches[index-1].amount_of_matches; num++ {
				if (index-1)+num >= 196 {
					continue
				}

				all_matches[(index-1)+num].extra_copies = all_matches[(index-1)+num].extra_copies + 1
			}
		}

	}
	all_copies := 0
	fmt.Print("part1 ", final_number, "\n")
	fmt.Print("part2 ", all_copies)
}
