package main

import (
	"fmt"
	"os"
)

func main() {
	switch os.Args[1] {
	case "day1":
		find_digits_in_lines()
		fmt.Print("\n")
		find_anything_in_lines()
	case "day2":
		fmt.Print("day2")
	}
}
