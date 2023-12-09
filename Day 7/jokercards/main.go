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
	'J': 1,
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

var jokerUpgrades = map[int]map[handClass]handClass{
	1: {
		highCard:     onePair,
		onePair:      threeOfAKind,
		twoPair:      fullHouse,
		threeOfAKind: fourOfAKind,
		fullHouse:    fourOfAKind,
		fourOfAKind:  fiveOfAKind,
		fiveOfAKind:  fiveOfAKind,
	},
	2: {
		highCard:     threeOfAKind,
		onePair:      fourOfAKind,
		twoPair:      fourOfAKind,
		threeOfAKind: fiveOfAKind,
		fullHouse:    fiveOfAKind,
		fourOfAKind:  fiveOfAKind,
		fiveOfAKind:  fiveOfAKind,
	},
	3: {
		highCard:     fourOfAKind,
		onePair:      fiveOfAKind,
		twoPair:      fiveOfAKind,
		threeOfAKind: fiveOfAKind,
		fullHouse:    fiveOfAKind,
		fourOfAKind:  fiveOfAKind,
		fiveOfAKind:  fiveOfAKind,
	},
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
	fmt.Println(counter)
	jokerCount := counter[rune('J')]
	delete(counter, rune('J'))

	countCount := map[int]int{}

	// count number of counts
	for _, v := range counter {
		countCount[v] += 1
	}

	result := highCard

	if countCount[5] == 1 {
		result = fiveOfAKind
	} else if countCount[4] == 1 {
		result = fourOfAKind
	} else if countCount[3] == 1 {
		if countCount[2] == 1 {
			result = fullHouse
		} else {
			result = threeOfAKind
		}
	} else if countCount[2] == 2 {
		result = twoPair
	} else if countCount[2] == 1 {
		result = onePair
	}

	if jokerCount == 0 {
		return handClass(result)
	}
	if jokerCount > 3 {
		return fiveOfAKind
	}
	return jokerUpgrades[jokerCount][handClass(result)]
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
		fmt.Println(newHand)
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
