package day4

import (
	"fmt"
	"strconv"
	"strings"
)

func convertStrArr(arr []string) []int {
	var inputNumbers = []int{}
	for _, i := range arr {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		inputNumbers = append(inputNumbers, j)
	}
	return inputNumbers
}

func RunPart1(inputLines []string) {
	inputNumbers := convertStrArr(strings.Split(inputLines[0], ","))

	// create board
	boardMap := make(map[int][][]int)
	boardNum := 0
	for _, line := range inputLines[2:] {
		if line != "" {
			row := convertStrArr(strings.Fields(line))
			boardMap[boardNum] = append(boardMap[boardNum], row)
		} else {
			boardNum++
		}
	}
	// add winners
	for key, board := range boardMap {
		boardCopy := make([][]int, len(board))
		copy(boardCopy, board)
		i := 0
		for i < 5 {
			col := make([]int, 0)
			for _, row := range board {
				col = append(col, row[i])
			}
			i++
			boardCopy = append(boardCopy, col)
		}
		boardMap[key] = boardCopy
	}
	fmt.Println(playBingo(boardMap, inputNumbers))
}

func playBingo(boardMap map[int][][]int, inputNumbers []int) int {
	calledList := make([]int, len(inputNumbers))
	for i, pickedNum := range inputNumbers {
		calledList[i] = pickedNum
		if i >= 4 {
			for _, board := range boardMap {
				for _, winner := range board {
					if subset(winner, calledList) {
						fmt.Println("WINNER")
						fmt.Println(pickedNum)
						fmt.Println(winner)
						sum := sumArray(calledList, board[:5]) // Remove the vertical slices from the sum
						fmt.Println(sum)
						return pickedNum*sum
					}
				}
			}
		}
	}
	return 0
}

func stringInSlice(a int, list []int) bool {
	for _, b := range list {
		if b == a {
			return true
		}
	}
	return false
}

func sumArray(arr []int, board [][]int) int {
	sum := 0
	for _, boardX := range board {
		for _, boardY := range boardX {
			if !stringInSlice(boardY, arr) {
				sum += boardY
			}
		}
	}
	return sum
}

func subset(first, second []int) bool {
	set := make(map[int]int)
	for _, value := range second {
		set[value] += 1
	}

	for _, value := range first {
		if count, found := set[value]; !found {
			return false
		} else if count < 1 {
			return false
		} else {
			set[value] = count - 1
		}
	}

	return true
}

