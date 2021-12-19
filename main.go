package main

import (
	"fmt"
	"github.com/jpoirier55/advent2021/day1"
	"github.com/jpoirier55/advent2021/day10"
	"github.com/jpoirier55/advent2021/day2"
	"github.com/jpoirier55/advent2021/day3"
	"github.com/jpoirier55/advent2021/day4"
	"github.com/jpoirier55/advent2021/day5"
	"github.com/jpoirier55/advent2021/day6"
	"github.com/jpoirier55/advent2021/day7"
	"github.com/jpoirier55/advent2021/day8"
	"github.com/jpoirier55/advent2021/day9"
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
	case "day4":
		day4.RunPart1(inputLines)
		day4.RunPart2(inputLines)
	case "day5":
		day5.RunPart1(inputLines)
		day5.RunPart2(inputLines)
	case "day6":
		day6.RunPart1(inputLines)
		day6.RunPart2(inputLines)
	case "day7":
		day7.RunPart1(inputLines)
		day7.RunPart2(inputLines)
	case "day8":
		day8.RunPart1(inputLines)
		day8.RunPart2(inputLines)
	case "day9":
		day9.RunPart1(inputLines)
		day9.RunPart2(inputLines)
	case "day10":
		day10.RunPart1(inputLines)
		day10.RunPart2(inputLines)
	}

}
