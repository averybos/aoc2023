package main

import (
	"fmt"
	"testing"
)

func test() {
	var n = Number{
		StartingIndex: 2,
		Number:        "234",
		Length:        3,
	}
	var p = Parts{
		Number{
			StartingIndex: 2,
			Number:        "234",
			Length:        3,
		},
		Number{
			StartingIndex: 8,
			Number:        "242",
			Length:        3,
		},
	}

	answer := isStructInParts(n, p)
	fmt.Print(answer)
}
func TestDeterminingParts(t *testing.T) {

	fileScanner := read_file("day_three_input_test.txt")
	text := place_text_into_array(fileScanner)
	everything := determine_part_and_symbol_placements(text)
	determine_if_a_real_part(everything)
	test()
}

func TestFindWinningNumbers(t *testing.T) {

	fileScanner := read_file("day_four_input_test.txt")
	text := place_text_into_array(fileScanner)
	find_the_winners(text)
}
