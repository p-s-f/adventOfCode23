package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func checkLine(numbers [][]int, lineToCheck string, partNumber int, validPartCount int) int {
	var left, right, partNumberToAdd int

	if numbers[partNumber][2] > 0 {
		left = numbers[partNumber][2] - 1
	}

	if numbers[partNumber][2]+numbers[partNumber][1] == (len(lineToCheck) - 1) {
		right = len(lineToCheck) - 1
	} else {
		right = numbers[partNumber][2] + (numbers[partNumber][1] + 1)
	}

	subString := lineToCheck[left:right]
	charCount := 0

	for charCount < len(subString) {
		charToCheck := subString[charCount : charCount+1]
		if _, err := strconv.Atoi(charToCheck); err != nil {
			if charToCheck != "." {
				partNumberToAdd = numbers[partNumber][0]
			}
		}
		charCount++
	}
	return validPartCount + partNumberToAdd
}

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
		foundPartNumbers := 0

		// Find all part numbers in this line adn add to numbers slice the value of the part number and the length of number in digits
		for i, match := range re.FindAllStringSubmatch(inputAsSlice[lineCount], -1) {
			partNumber, _ := strconv.Atoi(match[0])
			foundPartNumbers++
			numbers[i] = make([]int, 3)
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
		for i < foundPartNumbers {

			if lineCount > 0 {
				vaildPartCount = checkLine(numbers, inputAsSlice[lineCount-1], i, vaildPartCount)
			}

			vaildPartCount = checkLine(numbers, inputAsSlice[lineCount], i, vaildPartCount)

			if lineCount < numLines-1 {
				vaildPartCount = checkLine(numbers, inputAsSlice[lineCount+1], i, vaildPartCount)
			}

			i++
		}
		lineCount++
	}

	fmt.Println(vaildPartCount)
}
