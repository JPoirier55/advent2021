package main

import (
	"fmt"
	"github.com/jpoirier55/advent2021/day1"
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
	if day == "day1" {
		day1.RunPart1(inputLines)
		day1.RunPart2(inputLines)
	}
}
