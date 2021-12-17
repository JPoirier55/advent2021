package day9

import (
	"fmt"
	"strconv"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type Point struct {
	up *Point
	down *Point
	left *Point
	right *Point
	value int
	basin int
	visited bool
}

func RunPart1(inputLines []string) {
	pointMap := getPointMap(inputLines)
	sum := 0
	for _, pt := range pointMap {
		if (pt.right == nil || (pt.right.value > pt.value)) &&
			(pt.left == nil || (pt.left.value > pt.value)) &&
			(pt.up == nil || (pt.up.value > pt.value)) &&
			(pt.down == nil || (pt.down.value > pt.value)) {
			sum += pt.value + 1
		}
	}
	fmt.Println(sum)
}

func getPointMap(inputLines []string) []Point{
	lvlMap := make([][]*Point, len(inputLines))
	for i := range lvlMap {
		lvlMap[i] = make([]*Point, len(inputLines[0]))
	}
	for i, line := range inputLines {
		for j, point := range line {
			temp, err := strconv.Atoi(string(point))
			check(err)
			lvlMap[i][j] = &Point{value: temp, basin: -1, visited: false}
		}
	}
	points := make([]Point, 0)
	for y, line := range lvlMap {
		for x, point := range line {
			if y == 0 {
				// no up
				point.up = nil
			} else {
				point.up = lvlMap[y-1][x]
			}
			if y == len(lvlMap)-1 {
				// no down
				point.down = nil
			} else {
				point.down = lvlMap[y+1][x]
			}
			if x == 0 {
				// no left
				point.left = nil
			} else {
				point.left = lvlMap[y][x-1]
			}
			if x == len(line)-1 {
				// no right
				point.right = nil
			} else {
				point.right = lvlMap[y][x+1]
			}
			points = append(points, *point)
		}
	}
	return points
}

func printPoint(point Point) {
	if point.up != nil { fmt.Printf("up: %d ", point.up.value) } else { fmt.Printf("up: nil ") }
	if point.down != nil { fmt.Printf("down: %d ", point.down.value) } else { fmt.Printf("down: nil ") }
	if point.left != nil { fmt.Printf("left: %d ", point.left.value) } else { fmt.Printf("left: nil ") }
	if point.right != nil { fmt.Printf("right: %d ", point.right.value) } else { fmt.Printf("right: nil ") }
	fmt.Printf(" val: %d basin: %d", point.value, point.basin)
	fmt.Println()
}

func markBasin(point *Point, basin int) bool {
	debug := false
	point.visited = true
	if point.up != nil && point.up.value != 9 && point.up.basin == -1 && !point.up.visited{
		if debug { fmt.Println("going up") }
		markBasin(point.up, basin)
	} else {
		if debug { fmt.Println("no more up") }
		point.basin = basin
	}
	if point.down != nil && point.down.value != 9 && point.down.basin == -1 && !point.down.visited{
		if debug {fmt.Println("going down")}
		markBasin(point.down, basin)
	} else {
		if debug {fmt.Println("no more down")}
		point.basin = basin
	}
	if point.left != nil && point.left.value != 9 && point.left.basin == -1 && !point.left.visited{
		if debug {fmt.Println("going left")}
		markBasin(point.left, basin)
	} else {
		if debug {fmt.Println("no more left")}
		point.basin = basin
	}
	if point.right != nil && point.right.value != 9 && point.right.basin == -1 && !point.right.visited{
		if debug {fmt.Println("going right")}
		markBasin(point.right, basin)
	} else {
		if debug {fmt.Println("no more right")}
		point.basin = basin
	}
	if debug {
		fmt.Println(point)
		fmt.Println(point.basin)
	}
	return true
}

func RunPart2(inputLines []string) {
	pointMap := getPointMap(inputLines)
	currentBasin := 0
	basinMap := make(map[int]int)
	basins := make([]int, 0)
	for _, pt := range pointMap {
		pt2 := pt
		if pt2.value != 9 && pt2.basin == -1 && !pt2.visited{
			markBasin(&pt2, currentBasin)
		}
		if pt2.right != nil && pt2.right.basin == -1 && pt2.value == 9 && pt2.right.value != 9{
			currentBasin++
		}
		fmt.Println(pt2)
		basinMap[pt2.basin]++
		basins = append(basins, pt2.basin)
	}
	count := 0
	for _, y := range basins {
		if y == -1 {
			fmt.Print(9)
		} else {
			fmt.Print(y)
		}

		count++
		if count > 9 {
			fmt.Println()
			count = 0
		}
	}

	fmt.Println(basinMap)
	result := 1
	for i, t := range basinMap {
		if i != -1 {
			result *= t
		}
	}
	fmt.Println(result)


}
