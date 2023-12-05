package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
)

type Numbers struct {
	Integer int
	Name    string
}
type Num []Numbers

type Placement struct {
	Index int
	Word  string
}
type Placements []Placement

func (r Num) ValueOf(symbol string) int {
	for _, s := range r {
		if s.Name == symbol {
			return s.Integer
		}
	}

	return 0
}

func find_digits_in_lines() {
	readFile, err := os.Open("day_one_input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	bigNumber := 0
	for fileScanner.Scan() {
		split := strings.Split(fileScanner.Text(), "")
		res := []int{}

		for _, i := range split {

			num, err := strconv.Atoi(i)
			if err != nil {
				continue
			} else {
				res = append(res, num)
			}

			break
		}
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

		added := ""
		for _, number := range res {
			s := strconv.Itoa(number)
			added = added + s
		}

		back, _ := strconv.Atoi(added)
		bigNumber = back + bigNumber
	}

	fmt.Print(bigNumber)
	readFile.Close()
}

func append_and_replace(text string, index int, name string) Placement {
	placement := Placement{}
	placement.Index = index
	placement.Word = name

	return placement
}

func find_anything_in_lines() {
	readFile, err := os.Open("day_one_input.txt")

	if err != nil {
		fmt.Println(err)
	}
	fileScanner := bufio.NewScanner(readFile)
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

		// compare indices to see whose first!
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
	fmt.Print(bigNum)
	readFile.Close()
}
