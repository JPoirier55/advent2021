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
			gamma += "1"
			epsilon += "0"
		}else {
			gamma += "0"
			epsilon += "1"
		}
	}
	gammaDec := decimalFromBinary(gamma)
	epsilonDec := decimalFromBinary(epsilon)
	fmt.Println(gammaDec*epsilonDec)
}

func RunPart2(inputLines []string) {
	fmt.Println("O2")
	oxygen := findOxygen(inputLines)
	fmt.Println("CO2")
	co2 := findC02(inputLines)
	fmt.Println(oxygen)
	fmt.Println(co2)
	fmt.Println(oxygen*co2)
}

func findC02(inputLines []string) int {
	current := make([]string, len(inputLines))
	copy(current, inputLines)
	for i := 0; i <= len(inputLines[0])-1; i++ {
		zeros := make([]string, 0)
		ones := make([]string, 0)
		//fmt.Println("current")
		//fmt.Println(current)
		for j := 0; j <= len(current)-1; j++ {
			if current[j][i] == '0' {
				zeros = append(zeros, current[j])
			} else {
				ones = append(ones, current[j])
			}
		}
		if len(ones) == len(zeros) {
			fmt.Println(i)
			fmt.Println(current)
		}
		if len(ones) <= len(zeros) {
			current = make([]string, len(ones))
			copy(current, ones)
		} else if len(ones) == len(zeros) && len(ones) == 1 && len(zeros) == 1 {
			for _, val := range current {
				//fmt.Println(current)
				if val[i] == '0' {
					//fmt.Println(val)
					co2 := decimalFromBinary(val)
					return int(co2)
				}
			}
		} else {
			current = make([]string, len(zeros))
			copy(current, zeros)
		}
		//if len(current) < 5 {
		//	fmt.Println(current)
		//}
		if len(current) == 1 {
			co2 := decimalFromBinary(current[0])
			return int(co2)
		}
	}
	return 0
}

func findOxygen(inputLines []string) int {
	current := make([]string, len(inputLines))
	copy(current, inputLines)

	for i := 0; i <= len(inputLines[0])-1; i++ {
		zeros := make([]string, 0)
		ones := make([]string, 0)
		fmt.Println("current")
		fmt.Println(current)
		for j := 0; j <= len(current)-1; j++ {
			if current[j][i] == '0' {
				zeros = append(zeros, current[j])
			} else {
				ones = append(ones, current[j])
			}
		}
		if len(ones) >= len(zeros) {
			current = make([]string, len(ones))
			copy(current, ones)
		} else if len(ones) == len(zeros) && len(ones) == 1 && len(zeros) == 1 {
			for _, val := range current {
				if val[i] == '1' {
					oxygen := decimalFromBinary(val)
					return int(oxygen)
				}
			}
		} else {
			current = make([]string, len(zeros))
			copy(current, zeros)
		}

		//if len(current) < 5 {
		//	fmt.Println(current)
		//}
		if len(current) == 1 {
			oxygen := decimalFromBinary(current[0])
			return int(oxygen)
		}
	}
	return 0
}
