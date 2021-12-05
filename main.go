package main

import (
	"fmt"
	"github.com/jpoirier55/advent2021/day1"
	"github.com/jpoirier55/advent2021/day2"
	"github.com/jpoirier55/advent2021/day3"
	"github.com/jpoirier55/advent2021/file"
	"os"
)

var inputFile string

func main() {
	day := os.Args[1]
	inputFile = day + "/input.txt"
	if len(os.Args) > 2 {
		if os.Args[2] == "test" {
			inputFile = day + "/test_input.txt"
			fmt.Println("--------TEST--------")
		}
	}
	inputLines := file.ReadFile(inputFile)
	switch day {
	case "day1":
		day1.RunPart1(inputLines)
		day1.RunPart2(inputLines)
	case "day2":
		day2.RunPart1(inputLines)
		day2.RunPart2(inputLines)
	case "day3":
		day3.RunPart1(inputLines)
		day3.RunPart2(inputLines)
	}

}
