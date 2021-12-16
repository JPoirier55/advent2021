package day7

import (
	"fmt"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func RunPart1(inputLines []string) {
	lineSplit := strings.Split(inputLines[0],",")
	crabSubs := make([]int, len(lineSplit))
	for i, val := range lineSplit {
		temp, err := strconv.Atoi(val)
		check(err)
		crabSubs[i] = temp
	}
	fmt.Println(crabSubs)

	max := arrMax(crabSubs)
	i := 0
	cheapest := 999999
	for i < max {
		cost := 0
		for _, sub := range crabSubs {
			diff := sub - i
			abs := findAbs(diff)
			cost += abs
		}
		if cost < cheapest {
			cheapest = cost
		}
		i++
	}
	fmt.Println(cheapest)
}

func findAbs(val int) int {
	if val < 0 { return -val } else { return val }
}

func arrMax(arr []int) int {
	max := -99999
	for _, val := range arr {
		if val > max {
			max = val
		}
	}
	return max
}

func RunPart2(inputLines []string) {
	lineSplit := strings.Split(inputLines[0],",")
	crabSubs := make([]int, len(lineSplit))
	for i, val := range lineSplit {
		temp, err := strconv.Atoi(val)
		check(err)
		crabSubs[i] = temp
	}
	fmt.Println(crabSubs)

	max := arrMax(crabSubs)
	i := 0
	cheapest := 9999999999999999
	for i < max {
		cost := 0
		for _, sub := range crabSubs {
			diff := sub - i
			abs := findAbs(diff)
			cost += findFuelCost(abs)
		}
		if cost < cheapest {
			cheapest = cost
		}
		i++
	}
	fmt.Println(cheapest)
}

func findFuelCost(val int) int {
	i := 0
	cost := 0
	for i <= val {
		cost += i
		i++
	}
	return cost
}
