package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func solveSingleBoatPuzzle(input string) int {
	readFile, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	result := 1

	time := 0
	distance := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		timesString, ok := strings.CutPrefix(line, "Time: ")
		if ok {
			time = sliceToSingleInt(timesString)
		}
		distanceString, ok := strings.CutPrefix(line, "Distance: ")
		if ok {
			distance = sliceToSingleInt(distanceString)
		}
	}

	readFile.Close()

	// -x^2 + time * x - distance > 0
	quadraticResults := quadratic(-1, time, -1*distance)
	lowerEnd := int(math.Ceil(quadraticResults[0]))
	roundingAdjustment := 0
	if float64(lowerEnd) == quadraticResults[0] {
		roundingAdjustment += 1
	}
	higherEnd := int(math.Floor(quadraticResults[1]))
	if float64(higherEnd) == quadraticResults[1] {
		roundingAdjustment += 1
	}
	result *= higherEnd - lowerEnd + 1 - roundingAdjustment

	return result
}

func sliceToSingleInt(stringToSlice string) int {
	stringToSlice = strings.Trim(stringToSlice, " ")
	stringSlice := strings.Split(stringToSlice, " ")
	resultStrings := []string{}
	for _, numString := range stringSlice {
		if numString != " " {
			resultStrings = append(resultStrings, numString)
		}
	}
	result, _ := strconv.Atoi(strings.Join(resultStrings, ""))
	return result
}
