package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	result1 := solveHashPuzzle("input.txt")
	fmt.Println(result1)
	result2 := solveHashMapPuzzle("input.txt")
	fmt.Println(result2)
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

type reflector = struct {
	label       string
	focalLength *int
}

func solveHashMapPuzzle(input string) int {
	readFile, err := os.Open(input)
	if err != nil {
		panic(err)
	}
	defer readFile.Close()

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Scan()
	line := fileScanner.Text()

	steps := strings.Split(line, ",")

	hashMap := map[int]*[]reflector{}
	for _, step := range steps {
		if strings.HasSuffix(step, "-") {
			label := strings.TrimSuffix(step, "-")
			boxAddr, ok := hashMap[hashAlgorithm(label)]
			if !ok {
				continue
			}
			box := *boxAddr
			for i, reflector := range box {
				if reflector.label == label {
					newBox := append(box[:i], box[i+1:]...)
					hashMap[hashAlgorithm(label)] = &newBox
					break
				}
			}
		} else {
			label, focalLengthStr, _ := strings.Cut(step, "=")
			focalLength, _ := strconv.Atoi(focalLengthStr)
			boxAddr, ok := hashMap[hashAlgorithm(label)]
			if !ok {
				boxAddr = &[]reflector{}
				hashMap[hashAlgorithm(label)] = boxAddr
			}
			box := *boxAddr
			foundFlag := false
			for _, reflector := range box {
				if reflector.label == label {
					*(reflector.focalLength) = focalLength
					foundFlag = true
					break
				}
			}
			if !foundFlag {
				newBox := append(box, reflector{label, &focalLength})
				hashMap[hashAlgorithm(label)] = &newBox
			}
		}
	}
	result := 0
	for box, contents := range hashMap {
		for i, lens := range *contents {
			result += (box + 1) * (i + 1) * (*lens.focalLength)
		}
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
