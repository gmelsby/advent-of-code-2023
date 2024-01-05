package main

import (
	"bufio"
	"fmt"
	"hash/fnv"
	"os"
	"strings"
)

func main() {
	result1 := solveRockPuzzle("input.txt")
	fmt.Println(result1)
	result2 := solveSpinRockPuzzle("input.txt")
	fmt.Println(result2)
}

func solveRockPuzzle(input string) int {
	readFile, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer readFile.Close()

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

	result := 0
	for _, rockLine := range rockList {
		result += i - (rockLine - 1)
	}
	return result
}

func solveSpinRockPuzzle(input string) int {
	readFile, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	puzzle := []string{}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		puzzle = append(puzzle, line)
	}

	cycleStart, cycleLength := spinUntilCycle(puzzle)
	numCycles := (1000000000-cycleStart)%cycleLength + cycleStart
	for i := 0; i < numCycles; i++ {
		puzzle = cycle(puzzle)
	}

	return (calculateLoad(puzzle))
}

func calculateLoad(puzzle []string) int {
	result := 0
	for i, line := range puzzle {
		result += (len(puzzle) - i) * strings.Count(line, "O")
	}
	return result
}

// returns start of cycle, length of cycle
func spinUntilCycle(puzzle []string) (int, int) {
	cycleMap := map[int]int{}
	cycleMap[hashPuzzle(puzzle)] = 0

	step := 1
	for {
		puzzle = cycle(puzzle)
		puzzleHash := hashPuzzle(puzzle)
		if priorStep, ok := cycleMap[puzzleHash]; ok {
			return priorStep, step - priorStep
		}
		cycleMap[puzzleHash] = step
		step++
	}
}

func cycle(pattern []string) []string {
	for i := 0; i < 4; i++ {
		pattern = moveRocksNorth(pattern)
		pattern = rotatePattern(pattern)
	}
	return pattern
}

func hashPuzzle(puzzle []string) int {
	h := fnv.New32a()
	h.Write([]byte(strings.Join(puzzle, " ")))
	return int(h.Sum32())
}

func moveRocksNorth(pattern []string) []string {
	newPattern := []string{}
	ledger := make([]int, len(pattern[0]))
	for i, line := range pattern {
		newPattern = append(newPattern, pattern[i])
		for j, char := range line {
			switch char {
			case 'O':
				replacementLine := []rune(newPattern[i])
				replacementLine[j] = '.'
				newPattern[i] = string(replacementLine)

				newLine := []rune(newPattern[ledger[j]])
				newLine[j] = 'O'
				newPattern[ledger[j]] = string(newLine)
				ledger[j] += 1
			case '#':
				ledger[j] = i + 1
			}
		}
	}
	return newPattern
}

func rotatePattern(pattern []string) []string {
	rotatedPattern := []string{}
	for i := 0; i < len(pattern[0]); i++ {
		newLine := []byte{}
		for j := len(pattern) - 1; j >= 0; j-- {
			newLine = append(newLine, pattern[j][i])
		}
		rotatedPattern = append(rotatedPattern, string(newLine))
	}
	return rotatedPattern
}
