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
	totalScore := 0
	for _, card := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		thisCard := strings.TrimSpace(strings.Split(card, ":")[1])
		winningNumbers := strings.Split(thisCard, "|")[0]
		numbersYouHave := strings.Split(thisCard, "|")[1]

		winningNumbersSet := make(map[int]bool)
		numbersYouHaveSet := make(map[int]bool)
		re := regexp.MustCompile(`(\d+)`)

		for _, winningNumber := range re.FindAllStringSubmatch(winningNumbers, -1) {
			winningNumberInt, _ := strconv.Atoi(strings.TrimSpace(winningNumber[0]))
			winningNumbersSet[winningNumberInt] = true
		}

		for _, numberYouHave := range re.FindAllStringSubmatch(numbersYouHave, -1) {
			numberYouHaveInt, _ := strconv.Atoi(strings.TrimSpace(numberYouHave[0]))
			numbersYouHaveSet[numberYouHaveInt] = true
		}

		cardScore := 0
		for k := range numbersYouHaveSet {
			if winningNumbersSet[k] {
				if cardScore == 0 {
					cardScore = 1
				} else {
					cardScore = cardScore * 2
				}
			}
		}
		totalScore += cardScore
	}

	fmt.Println(totalScore)
}
