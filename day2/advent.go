package main

import (
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func main() {
	var gameCount, power int

	input, _ := os.ReadFile(os.Args[1:][0])
	re := regexp.MustCompile(`(\d+) (\w+)`)

	for gameNumber, line := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		var minRed, minGreen, minBlue int
		for _, match := range re.FindAllStringSubmatch(line, -1) {
			score, _ := strconv.Atoi(match[1])
			switch match[2] {
			case "red":
				if score > minRed {
					minRed = score
				}
			case "green":
				if score > minGreen {
					minGreen = score
				}
			case "blue":
				if score > minBlue {
					minBlue = score
				}
			}
		}
		if minRed <= 12 && minGreen <= 13 && minBlue <= 14 {
			gameCount += gameNumber + 1
		}
		power += minRed * minBlue * minGreen
	}
	fmt.Println(gameCount)
	fmt.Println(power)
}
