package main

import (
	"bufio"
	"bytes"
	"fmt"
	"io"
	"os"
	"regexp"
)

func lineCounter(r io.Reader) (int, error) {
	buf := make([]byte, 32*1024)
	count := 0
	lineSep := []byte{'\n'}

	for {
		c, err := r.Read(buf)
		count += bytes.Count(buf[:c], lineSep)

		switch {
		case err == io.EOF:
			return count, nil

		case err != nil:
			return count, err
		}
	}
}

func ReadLine(r io.Reader, lineNum int) (line string, lastLine int, err error) {
	sc := bufio.NewScanner(r)
	for sc.Scan() {
		lastLine++
		if lastLine == lineNum {
			// you can return sc.Bytes() if you need output in []bytes
			return sc.Text(), lastLine, sc.Err()
		}
	}
	return line, lastLine, io.EOF
}

func main() {
	// read input text file for processing as arg[0]
	argsWithoutProg := os.Args[1:]
	f, _ := os.Open(argsWithoutProg[0])
	numlines, _ := lineCounter(f)
	f, _ = os.Open(argsWithoutProg[0])
	line3, _, _ := ReadLine(f, 3)
	fmt.Printf("There are %d lines in input file\n", numlines)
	fmt.Printf("And line 3 is \"%s\"\n", line3)
	lineCount := 1
	for lineCount <= numlines {
		// f, _ = os.Open(argsWithoutProg[0])
		// currentLine, _, _ := ReadLine(f, lineCount)
		// fmt.Printf("Line %d: \"%s\"\n", lineCount, currentLine)
		// lineCount++

		f, _ = os.Open(argsWithoutProg[0])
		prevLine, _, _ := ReadLine(f, lineCount-1)
		fmt.Printf("Line %d: \"%s\"\n", lineCount-1, prevLine)

		f, _ = os.Open(argsWithoutProg[0])
		currentLine, _, _ := ReadLine(f, lineCount)
		fmt.Printf("Line %d: \"%s\"\n", lineCount, currentLine)

		f, _ = os.Open(argsWithoutProg[0])
		nextLine, _, _ := ReadLine(f, lineCount+1)
		fmt.Printf("Line %d: \"%s\"\n", lineCount+1, nextLine)

		re := regexp.MustCompile(`(\d+)`)
		var numbers [](string)
		numbers = make([]string, 2)
		for i, match := range re.FindAllStringSubmatch(nextLine, -1) {
			// fmt.Printf("Found match %d is %s\n", i+1, match[0])
			numbers[i] = match[0]
		}

		for i, match := range re.FindAllStringSubmatchIndex(nextLine, -1) {
			// fmt.Printf("Found match %d is %d\n", i+1, match[0])
			numbers[i] = fmt.Sprintf("Found %s has string index %d and is length %d\n", numbers[i], match[0], len(numbers[i]))
		}

		fmt.Println(numbers)

		fmt.Printf("\n******************\n")
		lineCount++
	}

}
