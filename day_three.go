package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// a symbol and its index
type Symbol struct {
	Index     int
	Symbol    string
	Adjacents []Number
}

type Number struct {
	StartingIndex int
	Number        string
	Length        int
	LineNumber    int
}

// an array of chars and their indices
type Symbols []Symbol
type Parts []Number

type allSymbols []Symbols
type allParts []Parts

type EverySymbolAndPart struct {
	AllSymbols allSymbols
	AllParts   allParts
}

// this function will be run before a part is detected to be adjacent
// to a symbol to ensure no duplicates
func isStructInParts(a Number, b Parts) bool {
	for _, b := range b {
		if b.StartingIndex == a.StartingIndex && b.Number == a.Number && b.LineNumber == a.LineNumber {
			return true
		}
	}
	return false
}

func isStructInSymbolArray(a Number, b Symbol) bool {
	for _, i := range b.Adjacents {
		if i.StartingIndex == a.StartingIndex && i.Number == a.Number && i.LineNumber == a.LineNumber {
			return true
		}
	}

	return false
}

// use indices here, if its an index that is -1:+1 on top or below a number, it counts, as well as right next to it on the same line
func determine_part_and_symbol_placements(lines []string) EverySymbolAndPart {

	var all_symbols = allSymbols{}
	var all_parts = allParts{}

	for i, line := range lines {

		each_char := strings.Split(line, "")

		var symbols = Symbols{}
		var numbers = Parts{}

		for index, chars := range each_char {

			// first gotta check if we have captured a number previously
			// because we don't want to repeat digits here
			// example of the bad: actual number=418 -> and then capturing 18 and then 8
			if index >= 1 && len(numbers) != 0 {
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
						Index:     index,
						Symbol:    chars,
						Adjacents: []Number{},
					}
					symbols = append(symbols, symbol_struct)
				}
			} else {
				// gotta iterate until the number is complete
				// construct a number here
				number := ""
				for i := index; i < len(each_char); i++ {
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
					LineNumber:    i,
				}
				numbers = append(numbers, number_struct)
			}
		}
		all_symbols = append(all_symbols, symbols)
		all_parts = append(all_parts, numbers)
	}
	var symbols_and_parts_all_lines = EverySymbolAndPart{
		AllSymbols: all_symbols,
		AllParts:   all_parts,
	}
	return symbols_and_parts_all_lines
}

func get_above_or_below_adjacents(symbols Symbols, numbers Number, final_confirmed_parts Parts) Parts {
	for i, outer := range symbols {
		potential_adjacents_left := float64(numbers.StartingIndex - outer.Index)
		potential_adjacents_right := float64((numbers.StartingIndex + numbers.Length - 1) - outer.Index)
		// checks for diagonals or aligned with the first or last digit
		if math.Abs(potential_adjacents_left) == 1 ||
			math.Abs(potential_adjacents_right) == 1 ||
			outer.Index == numbers.StartingIndex ||
			outer.Index == (numbers.StartingIndex+(numbers.Length-1)) {
			if !isStructInSymbolArray(numbers, outer) {
				symbols[i].Adjacents = append(outer.Adjacents, numbers)
			}
			if !isStructInParts(numbers, final_confirmed_parts) {
				final_confirmed_parts = append(final_confirmed_parts, numbers)
			}
		}
	}
	return final_confirmed_parts
}

func determine_if_a_real_part(all_items EverySymbolAndPart) {
	var final_confirmed_parts = Parts{}

	// big outer loop
	for index, outer := range all_items.AllParts {
		var above_symbols = Symbols{}
		var below_symbols = Symbols{}
		same_line_symbols := all_items.AllSymbols[index]
		if index > 0 {
			// after first line
			above_symbols = all_items.AllSymbols[index-1]
		}
		if index < len(all_items.AllParts)-1 {
			// before last line
			below_symbols = all_items.AllSymbols[index+1]
		}

		// check for directly adjacent symbols in the same line
		for _, numbers := range outer {
			for i, symbol := range same_line_symbols {
				potential_adjacents_left := float64(symbol.Index - numbers.StartingIndex)
				potential_adjacents_right := float64((numbers.StartingIndex + numbers.Length - 1) - symbol.Index)

				if math.Abs(potential_adjacents_left) == 1 || math.Abs(potential_adjacents_right) == 1 {
					if !isStructInSymbolArray(numbers, symbol) {
						same_line_symbols[i].Adjacents = append(symbol.Adjacents, numbers)
					}
					if !isStructInParts(numbers, final_confirmed_parts) {
						final_confirmed_parts = append(final_confirmed_parts, numbers)
					}
				}
			}
		}

		// now i'll check for diagonal, above, or below adjacents
		for _, numbers := range outer {
			final_confirmed_parts = get_above_or_below_adjacents(above_symbols, numbers, final_confirmed_parts)
			final_confirmed_parts = get_above_or_below_adjacents(below_symbols, numbers, final_confirmed_parts)
		}

	}
	// part 2
	sum_of_gear_ratios := 0
	for _, i := range all_items.AllSymbols {
		for _, j := range i {
			if j.Symbol == "*" {
				if len(j.Adjacents) == 2 {
					convert1, _ := strconv.Atoi(j.Adjacents[0].Number)
					convert2, _ := strconv.Atoi(j.Adjacents[1].Number)
					gear_ratio := convert1 * convert2
					sum_of_gear_ratios = sum_of_gear_ratios + gear_ratio
				} else {
					continue
				}
			} else {
				continue
			}
		}
	}
	// fmt.Print(final_confirmed_parts, "\n")
	bigNum := 0
	for _, n := range final_confirmed_parts {
		convert, _ := strconv.Atoi(n.Number)
		bigNum = bigNum + convert
	}
	fmt.Print("part1 ", bigNum, "\n")
	fmt.Print("part2 ", sum_of_gear_ratios)
}
