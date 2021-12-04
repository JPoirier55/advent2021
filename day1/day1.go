package day1

import (
	"fmt"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func RunPart1(inputLines []string) {
	fmt.Println("running day1 part 1")
	last := 0
	count := 0
	for i, val := range inputLines {
		current, err := strconv.Atoi(val)
		check(err)
		if i > 0 {
			if current > last {
				count++
			}
		}
		last = current
	}
	fmt.Println(count)
}

func RunPart2(inputLines []string) {
	fmt.Println("running day1 part 2")
	lastSum := 0
	count := 0
	for i, _ := range inputLines {
		var currentSum int
		if i < len(inputLines)-2 {
			current, err := strconv.Atoi(inputLines[i])
			current1, err := strconv.Atoi(inputLines[i+1])
			current2, err := strconv.Atoi(inputLines[i+2])
			check(err)
			currentSum = current + current1 + current2
			if i > 0 && currentSum > lastSum {
				count++
			}
		}
		lastSum = currentSum
	}
	fmt.Println(count)
}
