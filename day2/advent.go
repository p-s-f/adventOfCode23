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
	for gameNumber, line := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		fmt.Println(line)
		possible = true
		for _, match := range re.FindAllStringSubmatch(line, -1) {
			score, _ := strconv.Atoi(match[1])
			switch match[2] {
			case "red":
				if score <= 12 {
					possible = false
				}
			case "green":
				if score <= 13 {
					possible = false
				}
			case "blue":
				if score <= 14 {
					possible = false
				}
			}
			// if !impossible {
			// 	gameCount += gameNumber
			// 	fmt.Printf("Game %d was possible! Count is now %d\n", gameNumber, gameCount)
			// }

		}
		fmt.Printf("Possible game %v\n", possible)
		if possible {
			gameCount += gameNumber
		}
	}
	fmt.Println(gameCount)
}
