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
	for cardNumber, card := range strings.Split(strings.TrimSpace(string(input)), "\n") {
		thisCard := strings.TrimSpace(strings.Split(card, ":")[1])
		winningNumbers := strings.Split(thisCard, "|")[0]
		numbersYouHave := strings.Split(thisCard, "|")[1]
		// fmt.Println(card)
		// fmt.Println(thisCard)
		// fmt.Println(winningNumbers)
		// fmt.Println(numbersYouHave)
		// fmt.Printf("Card %d has winning nums: %v and yourNums: %v\n", cardNumber, winningNumbers, numbersYouHave)

		winningNumbersSet := make(map[int]bool)
		numbersYouHaveSet := make(map[int]bool)
		re := regexp.MustCompile(`(\d+)`)
		for _, winningNumber := range re.FindAllStringSubmatch(winningNumbers, -1) {
			// fmt.Printf("%s", winningNumber[0])
			winningNumberInt, _ := strconv.Atoi(strings.TrimSpace(winningNumber[0]))
			// fmt.Println(winningNumberInt)
			winningNumbersSet[winningNumberInt] = true
		}

		for _, numberYouHave := range re.FindAllStringSubmatch(numbersYouHave, -1) {
			numberYouHaveInt, _ := strconv.Atoi(strings.TrimSpace(numberYouHave[0]))
			numbersYouHaveSet[numberYouHaveInt] = true
		}

		// fmt.Println(winningNumbersSet)
		// fmt.Println(numbersYouHaveSet)

		cardScore := 0
		for k := range numbersYouHaveSet {
			// fmt.Printf("Numbers you have %d\n", k)
			// fmt.Printf("Does %d exist in winners? %v\n", k, winningNumbersSet[k])
			if winningNumbersSet[k] {
				if cardScore == 0 {
					cardScore = 1
				} else {
					cardScore = cardScore * 2
				}
				// fmt.Printf("Winner! Cardscore is now %d\n", cardScore)
			}
		}
		fmt.Printf("For card %d Cardscore is %d\n", cardNumber+1, cardScore)

		totalScore += cardScore
	}

	fmt.Println(totalScore)
}
