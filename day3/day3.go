package day3

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func decimalFromBinary(binaryStr string) int64 {
	if i, err := strconv.ParseInt(binaryStr, 2, 64); err != nil {
		fmt.Println(err)
		return 0
	} else {
		return i
	}
}

func RunPart1(inputLines []string) {
	totalLines := 0
	binMap := make(map[int]int)
	for _, val := range inputLines {
		totalLines++
		indVals := strings.Split(val, "")

		for i, indVal := range indVals {
			intVal, _ := strconv.Atoi(indVal)
			binMap[i] += intVal
		}
		fmt.Println(binMap)
	}
	keys := make([]int, 0, len(binMap))

	for k, _ := range binMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	gamma := ""
	epsilon := ""
	for _, k := range keys {
		if binMap[k] > totalLines/2 {
			fmt.Printf("%d : 1", k)
			gamma += "1"
			epsilon += "0"
		}else {
			fmt.Printf("%d : 0", k)
			gamma += "0"
			epsilon += "1"
		}
		fmt.Println()
	}
	fmt.Println(gamma)
	fmt.Println(epsilon)
	gammaDec := decimalFromBinary(gamma)
	epsilonDec := decimalFromBinary(epsilon)
	fmt.Println(gammaDec)
	fmt.Println(epsilonDec)
	fmt.Println(gammaDec*epsilonDec)
}

func RunPart2(inputLines []string) {
	totalLines := 0
	zeros := make([]string, 1)
	ones := make([]string, 1)
	for i, _ := range inputLines[0] {
		fmt.Println(string(inputLines[0][i]))
		for _, val := range inputLines {
			fmt.Println(val)
		}
	}



}
