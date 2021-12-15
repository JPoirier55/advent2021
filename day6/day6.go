package day6

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
	lanternFish := make([]int, len(lineSplit))
	for i, val := range lineSplit {
		temp, err := strconv.Atoi(val)
		check(err)
		lanternFish[i] = temp
	}

	//fmt.Println(lanternFish)
	day := 0
	for day < 150 {
		for i, fish := range lanternFish {
			if fish == 0 {
				lanternFish = append(lanternFish, 8)
				lanternFish[i] = 6
			} else {
				lanternFish[i]--
			}
		}
		day++
		//fmt.Println(lanternFish)
	}
	fmt.Println(len(lanternFish))

}

func RunPart2(inputLines []string) {
	lineSplit := strings.Split(inputLines[0],",")
	lanternFishRaw := make([]int64, 9)
	for _, val := range lineSplit {
		temp, err := strconv.Atoi(val)
		check(err)
		lanternFishRaw[temp] += 1
	}
	day := 0
	//fmt.Println(lanternFishRaw)
	for day < 256 {
		zeroNum := lanternFishRaw[0]
		lanternFishRaw[0] = 0
		for i, val := range lanternFishRaw {
			if i != 0 {
				lanternFishRaw[i-1] = val
				lanternFishRaw[i] = 0
			}
		}
		lanternFishRaw[8] += zeroNum
		lanternFishRaw[6] += zeroNum
		day++
		//fmt.Println(lanternFishRaw)
	}
	fmt.Println(sumArray(lanternFishRaw))



}

func sumArray(arr []int64) int64 {
	var sum int64
	for _, val := range arr {
		sum += int64(val)
	}
	return sum
}

func RunPart22(inputLines []string) {
	lineSplit := strings.Split(inputLines[0],",")
	lanternFishRaw := make([]int, len(lineSplit))
	for i, val := range lineSplit {
		temp, err := strconv.Atoi(val)
		check(err)
		lanternFishRaw[i] = temp
	}

	splitNum := 60
	totalFish := make([][]int, splitNum)
	index := 0
	for index < splitNum {
		length := len(lanternFishRaw)/splitNum
		firstSliceNum := index*length
		secondSliceNum := (index+1)*length
		totalFish[index] = lanternFishRaw[firstSliceNum:secondSliceNum]
		index++
	}

	fishNum := 0
	for _, lanternFish := range totalFish {
		fmt.Println(lanternFish)
		newFish := make([]int, len(lanternFish))
		copy(newFish, lanternFish)
		fmt.Println(len(newFish))
		day := 0
		for day < 200 {
			for i, val := range newFish {
				if val == 0 {
					newFish = append(newFish, 8)
					newFish[i] = 6
				} else {
					newFish[i]--
				}
			}
			day++
		}
		fmt.Println(len(newFish))
		fishNum += len(newFish)
		fmt.Println(fishNum)
	}
	// 12633835679
	//26984457539

}
