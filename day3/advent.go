package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	input, _ := os.ReadFile(os.Args[1:][0])

	inputAsSlice := strings.Split(strings.TrimSpace(string(input)), "\n")
	fmt.Println(inputAsSlice)

	lineCount := 0
	vaildPartCount := 0
	numLines := len(inputAsSlice)

	for lineCount < numLines {
		re := regexp.MustCompile(`(\d+)`)
		numbers := make([][]int, 20)
		partNumberToAdd := 0

		// Find all part numbers in this line adn add to numbers slice the value of the part number and the length of number in digits
		for i, match := range re.FindAllStringSubmatch(inputAsSlice[lineCount], -1) {
			partNumber, _ := strconv.Atoi(match[0])
			numbers[i][0] = partNumber
			numbers[i][1] = len(match[0])
		}

		// Find the string index position of each part number and add to numbers slice
		for i, match := range re.FindAllStringSubmatchIndex(inputAsSlice[lineCount], -1) {
			numbers[i][2] = match[0]
		}
		// numbers is now e.g
		/*
			    line = "...123....45...%.."
				numbers[0][0] = 123 // the part number itself
				numbers[0][1] = 3 // length of part number in digits
				numbers[0][2] = 3 // position of part number in line as string index
				numbers[1][0] = 456
				numbers[1][1] = 2
				numbers[1][2] = 10
		*/
		i := 0
		for i <= len(numbers) {
			adjacentSymbol := false

			if lineCount > 0 {
				for charToCheck := range inputAsSlice[lineCount-1][numbers[i][2]-1 : numbers[i][2]+numbers[i][1]] {
					if charToCheck != '.' {
						adjacentSymbol = true
						partNumberToAdd = numbers[i][0]
					}
				}
			}

			for charToCheck := range inputAsSlice[lineCount][numbers[i][2]-1 : numbers[i][2]+numbers[i][1]] {
				if charToCheck != '.' {
					adjacentSymbol = true
					partNumberToAdd = numbers[i][0]
				}
			}

			if lineCount < numLines-1 {
				for charToCheck := range inputAsSlice[lineCount+1][numbers[i][2]-1 : numbers[i][2]+numbers[i][1]] {
					if charToCheck != '.' {
						adjacentSymbol = true
						partNumberToAdd = numbers[i][0]
					}
				}
			}
			if adjacentSymbol {
				vaildPartCount += partNumberToAdd
			}
		}

		lineCount++
	}
	fmt.Println(vaildPartCount)
}
