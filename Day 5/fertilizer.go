package main

import (
	"bufio"
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
	seeds := []int{}
	// to store maps
	mapList := [][][]int{}
	mapFlag := false

	for fileScanner.Scan() {
		line := fileScanner.Text()

		// loads seeds
		if len(seeds) == 0 {
			seedStrings := strings.Split(strings.TrimPrefix(line, "seeds: "), " ")
			for _, seedString := range seedStrings {
				seedValue, err := strconv.Atoi(seedString)
				if err != nil {
					panic(err)
				}
				seeds = append(seeds, seedValue)
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

func solveLowestLocation(seeds []int, mapList [][][]int) int {
	locations := []int{}
	for _, currentVal := range seeds {
		for _, currentMap := range mapList {
			for _, subMap := range currentMap {
				if subMap[1] <= currentVal && currentVal < subMap[1]+subMap[2] {
					currentVal = subMap[0] + currentVal - subMap[1]
					break
				}
			}

		}
		locations = append(locations, currentVal)
	}
	return slices.Min(locations)
}
