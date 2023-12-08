package main

import (
	"bufio"
	"fmt"
	"os"
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
	mapList := []map[int]int{}
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
				mapList = append(mapList, map[int]int{})
			}

			if mapFlag {
				addLineValuesToMap(line, mapList[len(mapList)-1])
			}
		}
	}

	fmt.Println(seeds)
	fmt.Println(mapList)

	readFile.Close()
}

func addLineValuesToMap(line string, valueMap map[int]int) {
	// convert values to integers
	numStrings := strings.Split(line, " ")
	nums := []int{}
	for _, numString := range numStrings {
		num, _ := strconv.Atoi(numString)
		nums = append(nums, num)
	}
	// add values to map
	for idx := 0; idx < nums[2]; idx++ {
		valueMap[nums[1]+idx] = nums[0] + idx
	}
}
