package day11

import (
	"fmt"
	"strconv"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

type JellyFish struct {
	upLeft *JellyFish
	up *JellyFish
	upRight *JellyFish
	left *JellyFish
	right *JellyFish
	downLeft *JellyFish
	down *JellyFish
	downRight *JellyFish
	level int
}

func RunPart1(inputLines []string) {
	jellyFishMap := getJellyFishMap(inputLines)
	step := 0
	for step < 100 {
		for _, jellyFish := range jellyFishMap {
			jellyFish.level += 1
		}
		for _, jellyFish := range jellyFishMap {
			if jellyFish.level > 9 {
				//flash
				jellyFish.level = 0
			}
		}
		step++
	}
	for _, jellyFish := range jellyFishMap {
		fmt.Println(jellyFish)
	}
}

func markBasin(jellyFishIn *JellyFish) {
	debug := false
	jellyFish := jellyFishIn
	if debug { time.Sleep(1 * time.Second) }

	if jellyFish.up != nil && jellyFish.up.value != 9 && jellyFish.up.basin == -1 && !jellyFish.up.visited{
		if debug { fmt.Println("going up") }
		markBasin(jellyFish.up, basin)
	} else {
		if jellyFish.basin == -1 {
			jellyFish.basin = basin
		}
		if debug { fmt.Println("no more up") }
	}
	if jellyFish.down != nil && jellyFish.down.value != 9 && jellyFish.down.basin == -1 && !jellyFish.down.visited{
		if debug {fmt.Println("going down")}
		markBasin(jellyFish.down, basin)
	} else {
		if debug {fmt.Println("no more down")}
		if jellyFish.basin == -1 {
			jellyFish.basin = basin
		}
	}
	if jellyFish.left != nil && jellyFish.left.value != 9 && jellyFish.left.basin == -1 && !jellyFish.left.visited{
		if debug {fmt.Println("going left")}
		markBasin(jellyFish.left, basin)
	} else {
		if debug {fmt.Println("no more left")}
		if jellyFish.basin == -1 {
			jellyFish.basin = basin
		}
	}
	if jellyFish.right != nil && jellyFish.right.value != 9 && jellyFish.right.basin == -1 && !jellyFish.right.visited{
		if debug {fmt.Println("going right")}
		markBasin(jellyFish.right, basin)
	} else {
		if debug {fmt.Println("no more right")}
		if jellyFish.basin == -1 {
			jellyFish.basin = basin
		}
	}
	if debug {
		fmt.Println(jellyFish)
	}
}

func getJellyFishMap(inputLines []string) []*JellyFish{
	lvlMap := make([][]*JellyFish, len(inputLines))
	for i := range lvlMap {
		lvlMap[i] = make([]*JellyFish, len(inputLines[0]))
	}
	for i, line := range inputLines {
		for j, jellyFish := range line {
			temp, err := strconv.Atoi(string(jellyFish))
			check(err)
			lvlMap[i][j] = &JellyFish{level: temp}
		}
	}
	jellyFishes := make([]*JellyFish, 0)
	for y, line := range lvlMap {
		for x, jellyFish := range line {
			if y == 0 {
				// no up
				jellyFish.up = nil
				jellyFish.upLeft = nil
				jellyFish.upRight = nil
			} else {
				jellyFish.up = lvlMap[y-1][x]
				if x == 0 {
					jellyFish.upLeft = nil
				} else {
					jellyFish.upLeft = lvlMap[y-1][x-1]
				}
				if x == len(line) -1 {
					jellyFish.upRight = nil
				} else {
					jellyFish.upRight = lvlMap[y-1][x+1]
				}
			}
			if y == len(lvlMap)-1 {
				// no down
				jellyFish.down = nil
				jellyFish.downRight = nil
				jellyFish.downLeft = nil
			} else {
				jellyFish.down = lvlMap[y+1][x]
				if x == len(line)-1 {
					jellyFish.downRight = nil
				} else {
					jellyFish.downRight = lvlMap[y+1][x+1]
				}
				if x == 0 {
					jellyFish.downLeft = nil
				} else {
					jellyFish.downLeft = lvlMap[y+1][x-1]
				}
			}
			if x == 0 {
				// no left
				jellyFish.left = nil
			} else {
				jellyFish.left = lvlMap[y][x-1]
			}
			if x == len(line)-1 {
				// no right
				jellyFish.right = nil
			} else {
				jellyFish.right = lvlMap[y][x+1]
			}
			jellyFishes = append(jellyFishes, jellyFish)
		}
	}
	return jellyFishes
}

func RunPart2(inputLines []string) {

}