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

	count_map := map[int]int{}
	current_card_number := 1

	for fileScanner.Scan() {
		line := fileScanner.Text()
		_, numbers, _ := strings.Cut(line, ": ")
		have_numbers, winning_numbers, _ := strings.Cut(numbers, " | ")

		have_set := numStrToSet(have_numbers)
		winning_set := numStrToSet(winning_numbers)

		// iterate over set to get intersection count
		win_count := 0
		for num := range have_set {
			if _, ok := winning_set[num]; ok {
				win_count += 1
			}
		}

		// check and increment card count with the one card we start with
		card_count, ok := count_map[current_card_number]
		if !ok {
			card_count = 0
		}
		count_map[current_card_number] = card_count + 1

		for card_no := current_card_number + 1; card_no <= current_card_number+win_count; card_no += 1 {
			card_no_count, ok := count_map[card_no]
			if !ok {
				card_no_count = 0
			}
			count_map[card_no] = card_no_count + count_map[current_card_number]
		}

		current_card_number += 1
	}
	readFile.Close()

	totalCount := 0
	for _, count := range count_map {
		totalCount += count
	}
	fmt.Println(totalCount)
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
