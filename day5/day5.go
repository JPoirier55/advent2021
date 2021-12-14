package day5

import (
	"fmt"
	"strconv"
	"strings"
)

type Vector struct {
	start []int
	end []int
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func maxes(){}

func Vectorize(line string) Vector {
	lineSplit := strings.Split(line," ")
	tempA := strings.Split(lineSplit[0], ",")
	tempB := strings.Split(lineSplit[2], ",")
	pointA := make([]int, 2)
	pointB := make([]int, 2)
	// set point a (start)
	tempA0, err := strconv.Atoi(tempA[0])
	check(err)
	tempA1, err := strconv.Atoi(tempA[1])
	check(err)
	pointA[0] = tempA0
	pointA[1] = tempA1
	// set point b (end)
	tempB0, err := strconv.Atoi(tempB[0])
	check(err)
	tempB1, err := strconv.Atoi(tempB[1])
	check(err)
	pointB[0] = tempB0
	pointB[1] = tempB1

	vector := Vector{start:pointA, end: pointB}

	return vector
}

func markPoints(vector Vector, boardMap map[string]int, part2 bool) {
	// ex 9,7 -> 7,9
	startX := vector.start[0]
	startY := vector.start[1]
	endX := vector.end[0]
	endY := vector.end[1]
	x1 := vector.start[0]
	y1 := vector.start[1]
	x2 := vector.end[0]
	y2 := vector.end[1]

	if !(startX == endX || startY == endY) {
		if !part2 {
			return
		}
		if x1 < x2 {
			if y1 < y2 {
				// x,y increase
				//fmt.Println("x,y increase")
				x := x1
				y := y1
				for x < x2 {
					addToBoard(x,y,boardMap)
					y++
					x++
				}
			} else {
				// x increases, y decreases
				//fmt.Println("x increases y decrease")
				x := x1
				y := y1
				for x < x2 {
					addToBoard(x,y,boardMap)
					y--
					x++
				}
			}
		} else {
			if y1 < y2 {
				// y increases, x decreases
				//fmt.Println("y increase, x decrease")
				x := x1
				y := y1
				for x > x2 {
					addToBoard(x,y,boardMap)
					y++
					x--
				}
			} else {
				//fmt.Println("both decrease")
				x := x1
				y := y1
				for x > x2 {
					addToBoard(x,y,boardMap)
					y--
					x--
				}
			}
		}
	} else {
		if vector.start[0] > vector.end[0] {
			startX = vector.end[0]
			endX = vector.start[0]
		}
		if vector.start[1] > vector.end[1] {
			startY = vector.end[1]
			endY = vector.start[1]
		}
		x := startX
		for x <= endX {
			y := startY
			for y <= endY {
				addToBoard(x,y,boardMap)
				y++
			}
			x++
		}
	}
}

func addToBoard(x int, y int, boardMap map[string]int) {
	//fmt.Print("x: ", x)
	//fmt.Print(" - y: ",y)
	//fmt.Println()
	boardMap[strconv.Itoa(x)+","+strconv.Itoa(y)] +=1
}

func sumBoard(boardMap map[string]int) int {
	sum := 0
	for _, point := range boardMap {
		if point > 1  {
			sum++
		}
	}
	return sum
}

func RunPart1(inputLines []string) {
	boardMap := make(map[string]int)
	for _, line := range inputLines {
		vector := Vectorize(line)
		markPoints(vector, boardMap, false)
	}
	fmt.Println(sumBoard(boardMap))
}

func RunPart2(inputLines []string) {
	boardMap := make(map[string]int)
	for _, line := range inputLines {
		vector := Vectorize(line)
		markPoints(vector, boardMap, true)
	}
	fmt.Println(sumBoard(boardMap))
}
