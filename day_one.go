package main

import (
	"bufio"
	"fmt"
	"sort"
	"strconv"
	"strings"
)

type Numbers struct {
	// an easy way to link int with their english name: 1 = one
	Integer int
	Name    string
}

// an array of Numbers
type Num []Numbers

type Placement struct {
	// Placement is a struct that contains the beginning index of a number written out and the word itself
	Index int
	Word  string
}

// an array of Placements
type Placements []Placement

// get the word, return it's int counterpart
func (r Num) ValueOf(symbol string) int {
	for _, s := range r {
		if s.Name == symbol {
			return s.Integer
		}
	}
	return 0
}

func find_digits_in_each_line(split []string) []int {
	res := []int{}
	// traverse forwards through list of string chars and stop upon first occurrence of an int
	for _, i := range split {

		num, err := strconv.Atoi(i)
		if err != nil {
			continue
		} else {
			res = append(res, num)
		}

		break
	}

	// traverse backwards thru list of string chars and stop upon last occurrence of an int
	for index := range split {
		reverse := len(split) - 1 - index
		num, err := strconv.Atoi(split[reverse])
		if err != nil {
			continue
		} else {
			res = append(res, num)
		}
		break
	}
	return res
}

func find_digits_in_lines() {
	fileScanner := read_file("day_one_input.txt")
	fileScanner.Split(bufio.ScanLines)

	bigNumber := 0
	for fileScanner.Scan() {
		split := strings.Split(fileScanner.Text(), "")

		res := find_digits_in_each_line(split)

		added := ""
		for _, number := range res {
			s := strconv.Itoa(number)
			added = added + s
		}

		back, _ := strconv.Atoi(added)
		bigNumber = back + bigNumber
	}

	// part one result
	fmt.Print(bigNumber)
}

func append_and_replace(text string, index int, name string) Placement {
	placement := Placement{}
	placement.Index = index
	placement.Word = name

	return placement
}

func find_anything_in_lines() {
	fileScanner := read_file("day_one_input.txt")
	fileScanner.Split(bufio.ScanLines)

	var allNums = Num{
		{1, "one"},
		{2, "two"},
		{3, "three"},
		{4, "four"},
		{5, "five"},
		{6, "six"},
		{7, "seven"},
		{8, "eight"},
		{9, "nine"},
	}
	bigNum := 0
	for fileScanner.Scan() {
		var placements = Placements{}
		text := fileScanner.Text()
		for _, str := range allNums {

			// numbers in word form
			if strings.Contains(text, str.Name) {
				index := strings.Index(text, str.Name)

				if index != strings.LastIndex(text, str.Name) {
					last_index := strings.LastIndex(text, str.Name)
					// append first occurrence
					placement1 := append_and_replace(text, index, str.Name)
					placements = append(placements, placement1)
					// append last occurrence, those are all that matter
					placement2 := append_and_replace(text, last_index, str.Name)
					placements = append(placements, placement2)
				} else {
					placement := append_and_replace(text, index, str.Name)
					placements = append(placements, placement)
				}
			}
		}

		split := strings.Split(fileScanner.Text(), "")
		for index, i := range split {

			num, err := strconv.Atoi(i)
			if err != nil {
				continue
			} else {
				placement := Placement{}
				placement.Index = index
				back := strconv.Itoa(num)
				placement.Word = back
				placements = append(placements, placement)
			}

			break
		}
		for index := range split {
			reverse := len(split) - 1 - index
			num, err := strconv.Atoi(split[reverse])
			if err != nil {
				continue
			} else {
				placement := Placement{}
				placement.Index = len(split) - 1 - index
				back := strconv.Itoa(num)
				placement.Word = back
				placements = append(placements, placement)
			}
			break
		}

		// compare indices to see whose first in the string!
		sort.Slice(placements, func(i, j int) bool {
			return placements[i].Index < placements[j].Index
		})
		final := Placements{placements[0], placements[len(placements)-1]}
		add := ""
		for _, values := range final {
			// deal with the string conversions
			num, err := strconv.Atoi(values.Word)
			if err != nil {
				value := allNums.ValueOf(values.Word)
				str_form := strconv.Itoa(value)
				add = add + str_form
			} else {
				str_form := strconv.Itoa(num)
				add = add + str_form
			}
		}
		back_to_int, _ := strconv.Atoi(add)
		bigNum = bigNum + back_to_int
	}
	// part two result
	fmt.Print(bigNum)
}
