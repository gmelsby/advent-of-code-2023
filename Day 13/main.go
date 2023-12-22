package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	result := solveMirrorPuzzle("input.txt")
	fmt.Println(result)
}

func solveMirrorPuzzle(input string) int {
	readFile, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	grids := [][]string{}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	grid := []string{}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		if len(line) == 0 {
			if len(grid) > 0 {
				grids = append(grids, grid)
			}
			grid = []string{}
			continue
		}
		grid = append(grid, line)
	}
	readFile.Close()

	if len(grid) != 0 {
		grids = append(grids, grid)
	}

	result := 0
	for _, pattern := range grids {
		result += symmetryValue(pattern)
	}
	return result
}

func symmetryValue(pattern []string) int {
	horizontalResult := horizontalSymmetry(pattern)
	if horizontalResult != -1 {
		return horizontalResult * 100
	}

	rotatedPattern := []string{}
	for i := 0; i < len(pattern[0]); i++ {
		newLine := []byte{}
		for j := 0; j < len(pattern); j++ {
			newLine = append(newLine, pattern[j][i])
		}
		rotatedPattern = append(rotatedPattern, string(newLine))
	}
	return horizontalSymmetry(rotatedPattern)
}

// returns line number of line before symmetry, -1 if no symmetry
func horizontalSymmetry(pattern []string) int {
	for i := 0; i < len(pattern); i++ {
		if symmetryExists(pattern[:i], pattern[i:]) {
			return i
		}
	}
	return -1
}

func symmetryExists(processed, toGo []string) bool {
	if len(processed) == 0 || len(toGo) == 0 {
		return false
	}
	for i := 0; i < min(len(processed), len(toGo)); i++ {
		if processed[len(processed)-1-i] != toGo[i] {
			return false
		}
	}
	return true
}
