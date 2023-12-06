package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

type Game struct {
	Color  string
	Amount int
}

func parse_each_game() {
	fileScanner := read_file("day_two_input.txt")
	fileScanner.Split(bufio.ScanLines)

	var red_limit = Game{
		Color:  "red",
		Amount: 12,
	}

	var green_limit = Game{
		Color:  "green",
		Amount: 13,
	}

	var blue_limit = Game{
		Color:  "blue",
		Amount: 14,
	}

	var allowed_to_play []int
	for fileScanner.Scan() {
		text := fileScanner.Text()

		game_id := strings.Split(strings.Split(text, ":")[0], " ")[1]
		game_id_number, _ := strconv.Atoi(game_id)

		split_by_semicolon := strings.Split(text, ":")
		join_back := strings.Join(split_by_semicolon[1:], ", ")
		games := strings.Split(join_back, ";")

		too_high := false
		for _, games := range games {

			// split on commas to get the amount per color
			each_color := strings.Split(games, ",")
			for _, color := range each_color {
				// create a struct for color and amount
				color_and_amount := strings.Split(color, " ")
				number, _ := strconv.Atoi(color_and_amount[1])

				var color_struct = Game{}
				color_struct.Color = color_and_amount[2]
				color_struct.Amount = number

				// fmt.Print(color_struct.Color, color_struct.Amount)
				if color_struct.Color == red_limit.Color {
					if color_struct.Amount > red_limit.Amount {
						too_high = true
					}
				} else if color_struct.Color == green_limit.Color {
					if color_struct.Amount > green_limit.Amount {
						too_high = true
					}
				} else {
					if color_struct.Amount > blue_limit.Amount {
						too_high = true
					}
				}
			}
		}
		if too_high == false {
			allowed_to_play = append(allowed_to_play, game_id_number)
		}

	}
	bigNum := 0
	for _, i := range allowed_to_play {
		bigNum = bigNum + i
	}
	// final answer
	fmt.Print(bigNum)
}
