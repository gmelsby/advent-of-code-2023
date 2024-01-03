package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	result1 := solveHashPuzzle("input.txt")
	fmt.Println(result1)
}

func solveHashPuzzle(input string) int {
	readFile, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Scan()
	line := fileScanner.Text()

	steps := strings.Split(line, ",")

	result := 0
	for _, step := range steps {
		result += hashAlgorithm(step)
	}
	return result
}

func hashAlgorithm(input string) int {
	current := 0
	for _, char := range input {
		current += int(char)
		current *= 17
		current = current % 256
	}
	return current
}
