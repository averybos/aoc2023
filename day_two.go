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

func parse_each_game(text string) (int, []string) {
	game_id := strings.Split(strings.Split(text, ":")[0], " ")[1]
	game_id_number, _ := strconv.Atoi(game_id)

	split_by_semicolon := strings.Split(text, ":")
	join_back := strings.Join(split_by_semicolon[1:], ", ")
	games := strings.Split(join_back, ";")

	return game_id_number, games
}

func make_the_comparisons(red Game, green Game, blue Game, color string, too_high bool) (bool, Game, Game, Game) {
	color_and_amount := strings.Split(color, " ")
	number, _ := strconv.Atoi(color_and_amount[1])

	var color_struct = Game{}
	color_struct.Color = color_and_amount[2]
	color_struct.Amount = number

	if color_struct.Color == red.Color {
		if color_struct.Amount > red.Amount {
			red.Amount = color_struct.Amount
			too_high = true
		}
	} else if color_struct.Color == green.Color {
		if color_struct.Amount > green.Amount {
			green.Amount = color_struct.Amount
			too_high = true
		}
	} else {
		if color_struct.Amount > blue.Amount {
			blue.Amount = color_struct.Amount
			too_high = true
		}
	}
	return too_high, red, green, blue
}

func determine_players() {
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

		game_id_number, games := parse_each_game(text)

		too_high := false
		for _, games := range games {

			// split on commas to get the amount per color
			each_color := strings.Split(games, ",")
			for _, color := range each_color {
				too_high, _, _, _ = make_the_comparisons(red_limit, green_limit, blue_limit, color, too_high)
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

func determine_least_amount_of_cubes() {
	fileScanner := read_file("day_two_input.txt")
	fileScanner.Split(bufio.ScanLines)

	var powers []int
	for fileScanner.Scan() {
		text := fileScanner.Text()
		_, games := parse_each_game(text)
		var red_highest_amount = Game{
			Color:  "red",
			Amount: 0,
		}

		var green_highest_amount = Game{
			Color:  "green",
			Amount: 0,
		}

		var blue_highest_amount = Game{
			Color:  "blue",
			Amount: 0,
		}
		for _, games := range games {

			// split on commas to get the amount per color
			each_color := strings.Split(games, ",")
			for _, color := range each_color {
				_, red_highest_amount, green_highest_amount, blue_highest_amount = make_the_comparisons(red_highest_amount, green_highest_amount, blue_highest_amount, color, false)
			}
		}
		power_of_set_of_cubes := red_highest_amount.Amount * green_highest_amount.Amount * blue_highest_amount.Amount
		powers = append(powers, power_of_set_of_cubes)
	}
	bigNum := 0
	for _, i := range powers {
		bigNum = bigNum + i
	}
	//  final answer
	fmt.Print(bigNum)
}
