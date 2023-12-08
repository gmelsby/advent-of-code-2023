package main

import (
	"bufio"
	"cmp"
	"fmt"
	"os"
	"slices"
	"strconv"
	"strings"
)

func main() {
	readFile, err := os.Open("input.txt")
	if err != nil {
		panic(err)
	}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	// to store seeds
	seeds := [][]int{}
	// to store maps
	mapList := [][][]int{}
	mapFlag := false

	for fileScanner.Scan() {
		line := fileScanner.Text()

		// loads seeds
		if len(seeds) == 0 {
			seedStrings := strings.Split(strings.TrimPrefix(line, "seeds: "), " ")
			for i := 0; i < len(seedStrings); i += 2 {
				seedValue, _ := strconv.Atoi(seedStrings[i])
				rangeValue, _ := strconv.Atoi(seedStrings[i+1])

				seeds = append(seeds, []int{seedValue, seedValue + rangeValue - 1})
			}
		} else {
			// blank line or label
			if len(line) == 0 || line == "\n" || strings.HasSuffix(line, ":") {
				mapFlag = false
				// first line in new map
			} else if mapFlag == false {
				mapFlag = true
				mapList = append(mapList, [][]int{})
			}

			if mapFlag {
				mapList[len(mapList)-1] = addLineValuesToMap(line, mapList[len(mapList)-1])
			}
		}
	}

	fmt.Println(solveLowestLocation(seeds, mapList))
	readFile.Close()
}

func addLineValuesToMap(line string, valueMap [][]int) [][]int {
	// convert values to integers
	numStrings := strings.Split(line, " ")
	nums := []int{}
	for _, numString := range numStrings {
		num, _ := strconv.Atoi(numString)
		nums = append(nums, num)
	}
	valueMap = append(valueMap, nums)
	return valueMap
}

func solveLowestLocation(seeds [][]int, mapList [][][]int) int {
	currentRanges := seeds
	for _, currentMap := range mapList {
		nextRanges := [][]int{}
		for _, subMap := range currentMap {
			newCurrentRanges := [][]int{}
			for _, currentRange := range currentRanges {
				bottomWithinUpper := subMap[1] < currentRange[1]
				topWithinLower := subMap[1]+subMap[2]-1 >= currentRange[0]
				topWithinUpper := subMap[1]+subMap[2]-1 < currentRange[1]
				bottomWithinLower := subMap[1] > currentRange[0]
				// not in bounds
				if !bottomWithinUpper || !topWithinLower {
					newCurrentRanges = append(newCurrentRanges, currentRange)
					continue
				}
				if topWithinUpper && bottomWithinLower {
					newCurrentRanges = append(newCurrentRanges, []int{currentRange[0], subMap[1] - 1})
					newCurrentRanges = append(newCurrentRanges, []int{subMap[1] + subMap[2] - 1, currentRange[1]})
					nextRanges = append(nextRanges, []int{subMap[0], subMap[0] + subMap[2] - 1})
				} else if topWithinUpper {
					newCurrentRanges = append(newCurrentRanges, []int{subMap[1] + subMap[2], currentRange[1]})
					nextRanges = append(nextRanges, []int{subMap[0] + currentRange[0] - subMap[1], subMap[0] + subMap[2] - 1})
				} else if bottomWithinLower {
					newCurrentRanges = append(newCurrentRanges, []int{currentRange[0], subMap[1] - 1})
					nextRanges = append(nextRanges, []int{subMap[0], subMap[0] + currentRange[1] - subMap[1]})
				} else {
					nextRanges = append(nextRanges, []int{subMap[0] + currentRange[0] - subMap[1], subMap[0] + currentRange[1] - subMap[1]})
				}
			}
			currentRanges = newCurrentRanges
		}
		currentRanges = append(currentRanges, nextRanges...)
	}

	return slices.MinFunc(currentRanges, func(a, b []int) int {
		return cmp.Compare(a[0], b[0])
	})[0]
}
