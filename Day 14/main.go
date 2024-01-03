package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	result1 := solveRockPuzzle("input.txt")
	fmt.Println(result1)
}

func solveRockPuzzle(input string) int {
	readFile, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	ledger := []int{}
	rockList := []int{}
	i := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if i == 0 {
			ledger = make([]int, len(line))
		}
		for j, char := range line {
			switch char {
			case 'O':
				ledger[j] += 1
				rockList = append(rockList, ledger[j])
			case '#':
				ledger[j] = i + 1
			}
		}
		i++
	}
	readFile.Close()

	result := 0
	for _, rockLine := range rockList {
		result += i - (rockLine - 1)
	}
	return result
}
