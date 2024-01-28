package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

type part struct {
	partNumber       int
	partNumberLength int
	partLineIndex    int
}

func checkLine(numbers []part, lineToCheck string, partNumber int, validPartCount int) int {
	var left, right, partNumberToAdd int

	if numbers[partNumber].partLineIndex > 0 {
		left = numbers[partNumber].partLineIndex - 1
	}

	if numbers[partNumber].partLineIndex+numbers[partNumber].partNumberLength == (len(lineToCheck) - 1) {
		right = len(lineToCheck) - 1
	} else {
		right = numbers[partNumber].partLineIndex + (numbers[partNumber].partNumberLength + 1)
	}

	subString := lineToCheck[left:right]
	charCount := 0

	for charCount < len(subString) {
		charToCheck := subString[charCount : charCount+1]
		if _, err := strconv.Atoi(charToCheck); err != nil {
			if charToCheck != "." {
				partNumberToAdd = numbers[partNumber].partNumber
			}
		}
		charCount++
	}
	return validPartCount + partNumberToAdd
}

func main() {
	input, _ := os.ReadFile(os.Args[1:][0])

	inputAsSlice := strings.Split(strings.TrimSpace(string(input)), "\n")

	lineCount := 0
	vaildPartCount := 0
	numLines := len(inputAsSlice)

	for lineCount < numLines {
		re := regexp.MustCompile(`(\d+)`)
		numbers := make([]part, 20)
		foundPartNumbers := 0

		// Find all part numbers in this line adn add to numbers slice the value of the part number and the length of number in digits
		for i, match := range re.FindAllStringSubmatch(inputAsSlice[lineCount], -1) {
			partNumber, _ := strconv.Atoi(match[0])
			foundPartNumbers++
			numbers[i].partNumber = partNumber
			numbers[i].partNumberLength = len(match[0])
		}

		// Find the string index position of each part number and add to numbers slice
		for i, match := range re.FindAllStringSubmatchIndex(inputAsSlice[lineCount], -1) {
			numbers[i].partLineIndex = match[0]
		}

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
