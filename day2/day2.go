package day2

import (
	"fmt"
	"strconv"
	"strings"
)

func RunPart1(inputLines []string) {
	xPos := 0
	yPos := 0
	for _, val := range inputLines {
		cmd := strings.Split(val, " ")
		moveVal, _ := strconv.Atoi(cmd[1])
		switch cmd[0] {
		case "forward":
			xPos += moveVal
		case "down":
			yPos += moveVal
		case "up":
			yPos -= moveVal
		}
	}
	fmt.Println(xPos * yPos)
}

func RunPart2(inputLines []string) {
	xPos := 0
	yPos := 0
	aim := 0
	for _, val := range inputLines {
		cmd := strings.Split(val, " ")
		moveVal, _ := strconv.Atoi(cmd[1])
		switch cmd[0] {
		case "forward":
			yPos += moveVal
			xPos += moveVal * aim
		case "down":
			aim += moveVal
		case "up":
			aim -= moveVal
		}
	}
	fmt.Println(xPos * yPos)
}
