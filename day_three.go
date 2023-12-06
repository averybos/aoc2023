package main

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"
)

// a symbol and its index
type Symbol struct {
	Index  int
	Symbol string
}

type Number struct {
	StartingIndex int
	Number        string
	Length        int
}

// an array of chars and their indices
type Symbols []Symbol
type PotentialParts []Number

// use indices here, if its an index that is -1:+1 on top or below a number, it counts, as well as right next to it on the same line
func determine_part_numbers() {
	fileScanner := read_file("day_three_input.txt")
	fileScanner.Split(bufio.ScanLines)

	// var symbols_to_compare = Symbols{}
	// var numbers_to_compare = NumberOrSymbols{}

	for fileScanner.Scan() {

		text := fileScanner.Text()
		each_char := strings.Split(text, "")

		var symbols = Symbols{}
		var numbers = PotentialParts{}

		for index, chars := range each_char {

			// first gotta check if we have captured a number previously
			// because we don't want to repeat digits here
			// ex: 418 and then capturing 18
			if index > 0 && len(numbers) != 0 {
				_, err := strconv.Atoi(each_char[index-1])
				if err == nil {
					continue
				}
			}

			_, err := strconv.Atoi(chars)

			if err != nil {
				if chars == "." {
					continue
				} else {
					var symbol_struct = Symbol{
						Index:  index,
						Symbol: chars,
					}
					symbols = append(symbols, symbol_struct)
				}
			} else {
				// gotta iterate until the number is complete
				// construct a number here
				number := ""
				for i := index; i < len(each_char)-1; i++ {
					_, err := strconv.Atoi(each_char[i])
					if err != nil {
						break
					} else {
						number = number + each_char[i]
					}
				}
				var number_struct = Number{
					StartingIndex: index,
					Number:        number,
					Length:        len(number),
				}
				numbers = append(numbers, number_struct)
			}
		}

		fmt.Print(symbols, numbers)
		fmt.Print("\n")
	}

}
