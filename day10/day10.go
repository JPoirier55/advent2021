package day10

import (
	"fmt"
	"sort"
)

type Stack []string

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// IsEmpty: check if stack is empty
func (s *Stack) IsEmpty() bool {
	return len(*s) == 0
}

// Push a new value onto the stack
func (s *Stack) Push(str string) {
	*s = append(*s, str) // Simply append the new value to the end of the stack
}

// Remove and return top element of stack. Return false if stack is empty.
func (s *Stack) Pop() (string, bool) {
	if s.IsEmpty() {
		return "", true
	} else {
		index := len(*s) - 1 // Get the index of the top most element.
		element := (*s)[index] // Index into the slice and obtain the element.
		*s = (*s)[:index] // Remove it from the stack by slicing it off.
		return element, false
	}
}

func findOppChar(char string) string {
	switch char {
	case "[":
		return "]"
	case "{":
		return "}"
	case "(":
		return ")"
	case "<":
		return ">"
	}
	return ""
}

func RunPart1(inputLines []string) {
	fails := make([]string, 0)
	for _, line := range inputLines {

		var stack Stack // create a stack variable of type Stack
		for _, c := range line {
			//time.Sleep(1 * time.Second)
			if string(c) == "}" {
				retVal, empty := stack.Pop()
				if retVal != "{" || empty {
					fails = append(fails, string(c))
				}
			} else if string(c) == "]" {
				retVal, empty := stack.Pop()
				if retVal != "[" || empty {
					fails = append(fails, string(c))
				}
			} else if string(c) == ")" {
				retVal, empty := stack.Pop()
				if string(retVal) != "(" || empty {
					fails = append(fails, string(c))
				}
			} else if string(c) == ">" {
				retVal, empty := stack.Pop()
				if retVal != "<" || empty {
					fails = append(fails, string(c))
				}
			} else {

				stack.Push(string(c))
			}
		}
	}
	fmt.Println(fails)
	returnSum := 0
	for _, err := range fails {
		switch err {
		case ")":
			returnSum += 3
		case "]":
			returnSum += 57
		case "}":
			returnSum += 1197
		case ">":
			returnSum += 25137
		}
	}
	fmt.Println(returnSum)
}

func RunPart2(inputLines []string) {
	scores := make([]int, 0)
	for _, line := range inputLines {
		failed := false
		var stack Stack // create a stack variable of type Stack
		for _, c := range line {
			if string(c) == "}" {
				retVal, empty := stack.Pop()
				if retVal != "{" || empty {
					failed = true
					break
				}
			} else if string(c) == "]" {
				retVal, empty := stack.Pop()
				if retVal != "[" || empty {
					failed = true
					break
				}
			} else if string(c) == ")" {
				retVal, empty := stack.Pop()
				if string(retVal) != "(" || empty {
					failed = true
					break
				}
			} else if string(c) == ">" {
				retVal, empty := stack.Pop()
				if retVal != "<" || empty {
					failed = true
					break
				}
			} else {
				stack.Push(string(c))
			}
		}
		if !failed {
			additions := make([]string, 0)
			empty := false
			for !empty {
				retVal, emp := stack.Pop()
				empty = emp
				if !empty {
					opp := findOppChar(retVal)
					additions = append(additions, opp)
				}
			}
			score := 0
			for _, addition := range additions {
				score *= 5
				switch addition {
				case ")":
					score += 1
				case "]":
					score += 2
				case "}":
					score += 3
				case ">":
					score += 4
				}
				fmt.Println(score)
			}
			fmt.Println(score)
			scores = append(scores, score)
		}
	}
	sort.Ints(scores)
	fmt.Println(scores)
	fmt.Println(scores[len(scores)/2])
}