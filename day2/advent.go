package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func strToInt(s string) int {
	if s != "" {
		i, err := strconv.Atoi(s)
		if err != nil {
			panic(err)
		}
		return i
	}
	return 0
}

func blueCheck(count int) bool {
	return count < 15
}

func redCheck(count int) bool {
	return count < 13
}

func greenCheck(count int) bool {
	return count < 14
}

func main() {
	// read codes text file for processing as arg[0]
	// argsWithoutProg := os.Args[1:]
	// f, _ := os.Open(argsWithoutProg[0])
	// scanner := bufio.NewScanner(f)
	// // process each line of input file
	// lineNum := 0
	// for scanner.Scan() {
	// 	line := scanner.Text()
	// 	calSum += getCalCode(line)
	// 	lineNum += 1
	// }
	games := strings.Split("Game 3: 6 green, 20 red, 25 green; 5 blue, 4 red, 13 green; 5 green, 1 red", ";")
	// fmt.Println(games)
	pat := regexp.MustCompile(`((\d+) blue)|((\d+) red)|((\d+) green)`)
	matches := pat.FindAllStringSubmatch(games[0], -1)
	// fmt.Println(matches[1][2])
	// fmt.Println(len(matches[1]))
	justDigit := regexp.MustCompile(`^\d+$`)
	colourMatcher := regexp.MustCompile(`\d+ (\w+)`)
	var currentColour string
	var cubeCount int
	for _, colours := range matches {
		for _, colour := range colours {
			if colourMatcher.MatchString(colour) {
				current := colourMatcher.FindAllStringSubmatch(colour, -1)
				currentColour = current[0][1]
			}
			if justDigit.MatchString(colour) {
				cubeCount = strToInt(colour)
				if currentColour == "blue" {
					if !blueCheck(cubeCount) {
						fmt.Println(currentColour)
						fmt.Println("IMPOSSIBLE")
					}
				}
				if currentColour == "red" {
					if !redCheck(cubeCount) {
						fmt.Println(currentColour)
						fmt.Println("IMPOSSIBLE")
					}
				}
				if currentColour == "green" {
					if !greenCheck(cubeCount) {
						fmt.Println(currentColour)
						fmt.Println("IMPOSSIBLE")
					}
				}
				fmt.Println(currentColour)
				fmt.Println(cubeCount)
			}
		}
	}
	blue := strToInt(matches[0][2])
	if blue < 15 {

	}

}
