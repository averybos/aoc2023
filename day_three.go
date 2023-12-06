package main

import (
	"bufio"
	"math"
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
type Parts []Number

// use indices here, if its an index that is -1:+1 on top or below a number, it counts, as well as right next to it on the same line
func determine_part_numbers(file *bufio.Scanner) {

	var final_confirmed_parts = Parts{}

	file.Split(bufio.ScanLines)

	// var symbols_to_compare = Symbols{}
	// var numbers_to_compare = Numbers{}

	for file.Scan() {

		text := file.Text()
		each_char := strings.Split(text, "")

		var symbols = Symbols{}
		var numbers = Parts{}

		for index, chars := range each_char {

			// first gotta check if we have captured a number previously
			// because we don't want to repeat digits here
			// example of the bad: actual number=418 -> and then capturing 18 and then 8
			if index > 1 && len(numbers) != 0 {
				_, err1 := strconv.Atoi(chars)
				_, err2 := strconv.Atoi(each_char[index-1])

				if err2 == nil {
					// check for the case where a symbol is directly after a number
					if err1 != nil {
						if each_char[index] == "." {
							continue
						}

					} else {
						continue
					}
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

		// fmt.Print(symbols, numbers)
		// fmt.Print("\n")

		// are there any symbols directly to the left or right of any number?
		for _, symbol := range symbols {
			for _, num := range numbers {
				potential_adjacents_left := float64(symbol.Index - num.StartingIndex)
				potential_adjacents_right := float64((num.StartingIndex + num.Length - 1) - symbol.Index)

				if math.Abs(potential_adjacents_left) == 1 || math.Abs(potential_adjacents_right) == 1 {
					final_confirmed_parts = append(final_confirmed_parts, num)
				}
			}
		}
		// begin comparing previous symbol indices with the current number index range
	}

}
