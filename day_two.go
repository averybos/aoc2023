package main

import (
	"bufio"
	"fmt"
	"strings"
)

type GameLimit struct {
	Color  string
	Amount int
}

func parse_each_game() {
	fileScanner := read_file("day_two_input.txt")
	fileScanner.Split(bufio.ScanLines)

	var red_limit = GameLimit{
		Color:  "red",
		Amount: 12,
	}

	var green_limit = GameLimit{
		Color:  "green",
		Amount: 13,
	}

	var blue_limit = GameLimit{
		Color:  "blue",
		Amount: 14,
	}

	fmt.Print(red_limit, green_limit, blue_limit)
	for fileScanner.Scan() {
		text := fileScanner.Text()

		game_id := strings.Split(strings.Split(text, ":")[0], " ")[1]
		// convert to an int ?

		fmt.Print(game_id, "\n")

		split_by_semicolon := strings.Split(text, ":")
		join_back := strings.Join(split_by_semicolon[1:], ", ")
		games := strings.Split(join_back, ";")

		for _, games := range games {
			fmt.Print(games, "\n")
		}

	}
}
