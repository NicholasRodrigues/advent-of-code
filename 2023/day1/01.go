package main

import (
	"fmt"
	"github.com/NicholasRodrigues/advent-of-code/pkg/files"
	"github.com/NicholasRodrigues/advent-of-code/pkg/numbers"
	"strconv"
)

// extractNumbers extracts numbers from the parsed data using the provided map
func extractNumbers(lineArr [][]string, numbersMap map[string]int) []int {
	var numbers []int
	for _, line := range lineArr {
		var num string
		for _, char := range line {
			if _, ok := numbersMap[char]; ok {
				num += char
			}
		}
		if num != "" {
			numbers = append(numbers, processNumber(num))
		}
	}
	return numbers
}

// processNumber processes a string number according to the given rules
func processNumber(num string) int {
	var stringToSum string
	if len(num) == 1 {
		stringToSum = string(num[0]) + string(num[0])
	} else {
		stringToSum = string(num[0]) + string(num[len(num)-1])
	}
	numToSum, err := strconv.Atoi(stringToSum)
	if err != nil {
		panic(err)
	}
	return numToSum
}

func main() {
	data, err := files.ReadFile("2023/day1/input.txt")
	if err != nil {
		panic(err)
	}

	numbersMap := map[string]int{
		"0": 0,
		"1": 1,
		"2": 2,
		"3": 3,
		"4": 4,
		"5": 5,
		"6": 6,
		"7": 7,
		"8": 8,
		"9": 9,
	}

	lineArr := files.ParseData(data)
	extractedNumbers := extractNumbers(lineArr, numbersMap)
	sum := numbers.SumNumbers(extractedNumbers)

	fmt.Printf("Sum of all numbers: %d\n", sum)
}
