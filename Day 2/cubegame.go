package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func isValidDrawList(maxMap map[string]int, drawList []string) bool {
	for _, numberColorString := range drawList {
		number, color, _ := strings.Cut(numberColorString, " ")
		cubeCount, _ := strconv.Atoi(number)
		if maxMap[color] < cubeCount {
			return false
		}
	}
	return true
}

func main() {
	var maxMap = map[string]int{
		"red":   12,
		"green": 13,
		"blue":  14,
	}

	readFile, err := os.Open("input.txt")
	check(err)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	total := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		before, after, _ := strings.Cut(line, ": ")
		// isolate game number
		gameNumber, err := strconv.Atoi(strings.TrimPrefix(before, "Game "))
		check(err)

		// replace ; with , for uniformity, then split into slice of number/color pairs
		cubeDraws := strings.Split(strings.ReplaceAll(after, ";", ","), ", ")
		if isValidDrawList(maxMap, cubeDraws) {
			total += gameNumber
		}

	}

	readFile.Close()
	fmt.Println(total)
}
