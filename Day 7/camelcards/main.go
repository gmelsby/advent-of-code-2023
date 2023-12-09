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
	fmt.Println(solveCamelCards("../input.txt"))
}

type hand struct {
	cards string
	bid   int
	class handClass
}

// enum-like handling of hand class
type handClass int

const (
	highCard = iota
	onePair
	twoPair
	threeOfAKind
	fullHouse
	fourOfAKind
	fiveOfAKind
)

var cardValues = map[byte]int{
	'A': 14,
	'K': 13,
	'Q': 12,
	'J': 11,
	'T': 10,
	'9': 9,
	'8': 8,
	'7': 7,
	'6': 6,
	'5': 5,
	'4': 4,
	'3': 3,
	'2': 2,
}

func compareHands(hand1, hand2 hand) int {
	if hand1.class > hand2.class {
		return 1
	} else if hand1.class < hand2.class {
		return -1
	}
	for idx, card := range hand1.cards {
		comparison := cmp.Compare(cardValues[byte(card)], cardValues[hand2.cards[idx]])
		if comparison != 0 {
			return comparison
		}
	}
	return 0
}

func classifyCards(cards string) handClass {
	counter := map[rune]int{}
	for _, card := range cards {
		counter[card] += 1
	}
	switch uniqueCards := len(counter); uniqueCards {
	case 5:
		return highCard
	case 4:
		return onePair
	case 3:
		maxCount := 0
		for _, value := range counter {
			if value > maxCount {
				maxCount = value
			}
		}
		if maxCount == 3 {
			return threeOfAKind
		}
		return twoPair
	case 2:
		for _, value := range counter {
			if value == 2 || value == 3 {
				return fullHouse
			}
			return fourOfAKind
		}
	case 1:
		return fiveOfAKind
	}
	return highCard
}

func solveCamelCards(input string) int {
	readFile, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	hands := []hand{}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		handString, valueString, _ := strings.Cut(line, " ")
		handClass := classifyCards(handString)
		handValue, _ := strconv.Atoi(valueString)
		newHand := hand{handString, handValue, handClass}
		hands = append(hands, newHand)
	}
	readFile.Close()

	slices.SortFunc(hands, compareHands)
	result := 0
	for i, currentHand := range hands {
		result += (i + 1) * currentHand.bid
	}

	return result
}
