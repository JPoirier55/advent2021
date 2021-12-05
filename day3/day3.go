package day3

import (
	"fmt"
	"strconv"
	"strings"
)

func RunPart1(inputLines []string) {
	totalLines := 0
	m := make(map[int]int)
	for _, val := range inputLines {
		totalLines++
		indVals := strings.Split(val, "")

		for i, indVal := range indVals {
			intVal, _ := strconv.Atoi(indVal)
			m[i] += intVal
		}
		fmt.Println(m)
	}

	for i, val := range m {
		if val > totalLines/2 {
			fmt.Printf("%d : 1", i)
		}else {
			fmt.Printf("%d : 0", i)
		}
		fmt.Println()
	}
}

func RunPart2(inputLines []string) {

}
