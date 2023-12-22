package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	result1, result2 := solveMirrorPuzzle("input.txt")
	fmt.Println(result1)
	fmt.Println(result2)
}

func solveMirrorPuzzle(input string) (int, int) {
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

	result1 := 0
	result2 := 0
	for _, pattern := range grids {
		result1 += symmetryValue(pattern)
		result2 += smudgedSymmetryValue(pattern)
	}
	return result1, result2
}

func rotatePattern(pattern []string) []string {
	rotatedPattern := []string{}
	for i := 0; i < len(pattern[0]); i++ {
		newLine := []byte{}
		for j := 0; j < len(pattern); j++ {
			newLine = append(newLine, pattern[j][i])
		}
		rotatedPattern = append(rotatedPattern, string(newLine))
	}
	return rotatedPattern
}

func symmetryValue(pattern []string) int {
	horizontalResult := horizontalSymmetry(pattern)
	if horizontalResult != -1 {
		return horizontalResult * 100
	}

	return horizontalSymmetry(rotatePattern(pattern))
}

func smudgedSymmetryValue(pattern []string) int {
	horizontalResult := smudgedHorizontalSymmetry(pattern)
	if horizontalResult != -1 {
		return horizontalResult * 100
	}

	return smudgedHorizontalSymmetry(rotatePattern(pattern))
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

func smudgedHorizontalSymmetry(pattern []string) int {
	for i := 0; i < len(pattern); i++ {
		if smudgedSymmetryExists(pattern[:i], pattern[i:]) {
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

func smudgedSymmetryExists(processed, toGo []string) bool {
	if len(processed) == 0 || len(toGo) == 0 {
		return false
	}
	smudgeCount := 0
	for i := 0; i < min(len(processed), len(toGo)); i++ {
		smudgeCount += stringDifferences(processed[len(processed)-1-i], toGo[i])
		if smudgeCount > 2 {
			return false
		}
	}

	if smudgeCount == 1 {
		return true
	}

	return false
}

// returns the number of differences between two strings of the same length
func stringDifferences(str1, str2 string) int {
	count := 0
	for i, char := range str1 {
		if byte(char) != str2[i] {
			count += 1
		}
	}
	return count
}
