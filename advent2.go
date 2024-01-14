package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
)

func strToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return i
}

func getCalCode(s string) int {
	s = strings.ReplaceAll(s, "one", "o1e")
	s = strings.ReplaceAll(s, "two", "t2o")
	s = strings.ReplaceAll(s, "three", "t3e")
	s = strings.ReplaceAll(s, "four", "f4r")
	s = strings.ReplaceAll(s, "five", "f5e")
	s = strings.ReplaceAll(s, "six", "s6x")
	s = strings.ReplaceAll(s, "seven", "s7n")
	s = strings.ReplaceAll(s, "eight", "e8t")
	s = strings.ReplaceAll(s, "nine", "n9e")

	pat := regexp.MustCompile(`(\d)`)
	matches := pat.FindAllStringSubmatch(s, -1)
	code := strToInt(matches[0][0] + matches[len(matches)-1][0])

	return code
}

func main() {
	var calSum int = 0
	// read codes text file for processing as arg[0]
	argsWithoutProg := os.Args[1:]
	f, _ := os.Open(argsWithoutProg[0])
	scanner := bufio.NewScanner(f)
	// process each line of input file
	lineNum := 0
	for scanner.Scan() {
		line := scanner.Text()
		calSum += getCalCode(line)
		lineNum += 1
	}
	fmt.Println(calSum)
}
