package day8

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func RunPart1(inputLines []string) {
	digitMap := make(map[int]int)
	for _, inputLine := range inputLines {
		temp := strings.Split(inputLine, "|")
		digits := strings.Fields(temp[1])
		fmt.Println(digits)
		for _, d := range digits {
			digitMap[len(d)]++
		}
		fmt.Println(digitMap)
	}
	sum := digitMap[2] + digitMap[4] + digitMap[3] + digitMap[7]
	fmt.Println(sum)
}

func subset(second, first []byte) bool {
	set := make(map[byte]int)
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

func difference(a, b []byte) []byte {
	mb := make(map[byte]struct{}, len(b))
	for _, x := range b {
		mb[x] = struct{}{}
	}
	var diff []byte
	for _, x := range a {
		if _, found := mb[x]; !found {
			diff = append(diff, x)
		}
	}
	return diff
}

func RunPart2(inputLines []string) {
	totalVal := 0
	for _, inputLine := range inputLines {
		temp := strings.Split(inputLine, "|")
		scrambleDigits := strings.Fields(temp[0])
		outputDigits := strings.Fields(temp[1])
		sort.Slice(scrambleDigits, func(i, j int) bool {
			return len(scrambleDigits[i]) < len(scrambleDigits[j])
		})
		digitStrings := make(map[int][]byte)
		for _, digit := range scrambleDigits {
			digitBytes := []byte(digit)
			if len(digit) == 2 {
				digitStrings[1] = digitBytes
			}
			if len(digit) == 3 {
				digitStrings[7] = digitBytes
			}
			if len(digit) == 4 {
				digitStrings[4] = digitBytes
			}
			if len(digit) == 5 {
				if subset(digitBytes, digitStrings[1]) {
					digitStrings[3] = digitBytes
				} else {
					if subset(digitBytes, difference(digitStrings[4], digitStrings[1])) {
						digitStrings[5] = digitBytes
					} else {
						digitStrings[2] = digitBytes
					}
				}
			}
			if len(digit) == 6 {
				if !subset(digitBytes, digitStrings[1]) && subset(digitBytes, digitStrings[5]) {
					digitStrings[6] = digitBytes
				} else if subset(digitBytes, digitStrings[1]) && !subset(digitBytes, digitStrings[5]) {
					digitStrings[0] = digitBytes
				} else {
					digitStrings[9] = digitBytes
				}
			}
			if len(digit) == 7 { digitStrings[8] = digitBytes }
		}
		numMap := make(map[string]int)
		for num, numBytes := range digitStrings {
			sorted := SortString(string(numBytes))
			numMap[sorted] = num
		}
		outputValue := ""
		for _, digit := range outputDigits {
			sorted := SortString(digit)
			outputValue += strconv.Itoa(numMap[sorted])
		}
		curVal, err := strconv.Atoi(outputValue)
		check(err)
		totalVal += curVal
	}
	fmt.Println(totalVal)
}

type sortRunes []rune

func (s sortRunes) Less(i, j int) bool {
	return s[i] < s[j]
}

func (s sortRunes) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (s sortRunes) Len() int {
	return len(s)
}

func SortString(s string) string {
	r := []rune(s)
	sort.Sort(sortRunes(r))
	return string(r)
}
