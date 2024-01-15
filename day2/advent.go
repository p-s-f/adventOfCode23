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
	re := regexp.MustCompile(`(\d+) (\w+)`)
	gameCount := 0
	possible := true
	power := 0

	for gameNumber, line := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		possible = true
		minRed := 0
		minBlue := 0
		minGreen := 0
		for _, match := range re.FindAllStringSubmatch(line, -1) {
			score, _ := strconv.Atoi(match[1])
			switch match[2] {
			case "red":
				if score > minRed {
					minRed = score
				}
				if score > 12 {
					possible = false
				}
			case "green":
				if score > minGreen {
					minGreen = score
				}
				if score > 13 {
					possible = false
				}
			case "blue":
				if score > minBlue {
					minBlue = score
				}
				if score > 14 {
					possible = false
				}
			}
		}
		power += minRed * minBlue * minGreen
		if possible {
			gameCount += gameNumber + 1
		}
	}
	fmt.Println(gameCount)
	fmt.Println(power)
}
