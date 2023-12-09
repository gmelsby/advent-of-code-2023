package main

import (
	"bufio"
	"math"
	"os"
	"strconv"
	"strings"
)

func quadratic(a, b, c int) []float64 {
	discriminant := int(math.Pow(float64(b), 2)) - 4*a*c
	if discriminant < 0 {
		return []float64{}
	} else if discriminant == 0 {
		return []float64{float64(-b) / 2 * float64(a)}
	} else {
		discriminantSqrt := math.Sqrt(float64(discriminant))
		result := []float64{}
		result = append(result, (float64(-b)+discriminantSqrt)/2*float64(a))
		result = append(result, (float64(-b)-discriminantSqrt)/2*float64(a))
		return result
	}
}

func solveBoatPuzzle(input string) int {
	readFile, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	result := 1

	times := []int{}
	distances := []int{}
	for fileScanner.Scan() {
		line := fileScanner.Text()
		timesString, ok := strings.CutPrefix(line, "Time: ")
		if ok {
			times = sliceToInts(timesString)
		}
		distanceString, ok := strings.CutPrefix(line, "Distance: ")
		if ok {
			distances = sliceToInts(distanceString)
		}
	}

	readFile.Close()

	for idx, time := range times {
		// -x^2 + time * x - distance > 0
		quadraticResults := quadratic(-1, time, -1*distances[idx])
		// case where no better distance is possible
		if len(quadraticResults) != 2 {
			continue
		}
		lowerEnd := int(math.Ceil(quadraticResults[0]))
		if lowerEnd < 0 {
			lowerEnd = 0
		}
		roundingAdjustment := 0
		if float64(lowerEnd) == quadraticResults[0] {
			roundingAdjustment += 1
		}
		higherEnd := int(math.Floor(quadraticResults[1]))
		if higherEnd > int(distances[idx]) {
			higherEnd = int(distances[idx])
		}
		if float64(higherEnd) == quadraticResults[1] {
			roundingAdjustment += 1
		}
		result *= higherEnd - lowerEnd + 1 - roundingAdjustment
	}

	return result
}

func sliceToInts(stringToSlice string) []int {
	stringToSlice = strings.Trim(stringToSlice, " ")
	stringSlice := strings.Split(stringToSlice, " ")
	result := []int{}
	for _, numString := range stringSlice {
		num, err := strconv.Atoi(numString)
		if err != nil {
			continue
		}
		result = append(result, num)
	}
	return result
}
