package main

import (
	"bufio"
	"fmt"
	"math"
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

	total := 0

	for fileScanner.Scan() {
		line := fileScanner.Text()
		_, numbers, _ := strings.Cut(line, ": ")
		have_numbers, winning_numbers, _ := strings.Cut(numbers, " | ")

		have_set := numStrToSet(have_numbers)
		winning_set := numStrToSet(winning_numbers)

		// iterate over set to get intersection count
		count := 0
		for num := range have_set {
			if _, ok := winning_set[num]; ok {
				count += 1
			}
		}
		total += int(math.Pow(2, float64(count-1)))
	}

	readFile.Close()
	fmt.Println(total)
}

func numStrToSet(numberString string) map[int]struct{} {
	result := map[int]struct{}{}
	for _, num := range strings.Split(numberString, " ") {
		number, err := strconv.Atoi(num)
		if err != nil {
			continue
		}
		result[number] = struct{}{}
	}
	return result
}
