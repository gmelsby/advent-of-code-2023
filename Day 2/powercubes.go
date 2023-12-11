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

func getPower(drawList []string) int {
	maxMap := map[string]int{
		"red":   0,
		"green": 0,
		"blue":  0,
	}

	for _, numberColorString := range drawList {
		number, color, _ := strings.Cut(numberColorString, " ")
		cubeCount, _ := strconv.Atoi(number)
		if maxMap[color] < cubeCount {
			maxMap[color] = cubeCount
		}
	}

	power := 1

	for _, count := range maxMap {
		power *= count
	}

	return power
}

func main() {
	readFile, err := os.Open("input.txt")
	check(err)

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	total := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		_, cubeString, _ := strings.Cut(line, ": ")

		// replace ; with , for uniformity, then split into slice of number/color pairs
		cubeDraws := strings.Split(strings.ReplaceAll(cubeString, ";", ","), ", ")
		total += getPower(cubeDraws)

	}

	readFile.Close()
	fmt.Println(total)
}
