package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Println(solveWasteland("input.txt"))
	fmt.Println(solveGhostPuzzle("ghost-test-input.txt"))
}

type node struct {
	value     string
	leftNode  *node
	rightNode *node
}

func loadInput(input string) (string, map[string]*node) {
	readFile, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	nodeDict := map[string]*node{}

	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)

	fileScanner.Scan()
	directions := fileScanner.Text()
	fileScanner.Scan()

	for fileScanner.Scan() {
		line := fileScanner.Text()
		currentNodeString, subNodesString, _ := strings.Cut(line, " = ")
		currentNode := getNode(nodeDict, currentNodeString)
		subNodesString = subNodesString[1 : len(subNodesString)-1]
		leftNodeString, rightNodeString, _ := strings.Cut(subNodesString, ", ")
		currentNode.leftNode = getNode(nodeDict, leftNodeString)
		currentNode.rightNode = getNode(nodeDict, rightNodeString)
	}
	readFile.Close()
	return directions, nodeDict
}

func solveWasteland(input string) int {
	directions, nodeDict := loadInput(input)

	moveCount := 0
	presentLocation := *nodeDict["AAA"]

	for presentLocation.value != "ZZZ" {
		nextDirection := directions[moveCount%len(directions)]
		if nextDirection == 'R' {
			presentLocation = *presentLocation.rightNode
		} else {
			presentLocation = *presentLocation.leftNode
		}
		moveCount += 1
	}

	return moveCount
}

func solveGhostPuzzle(input string) int {
	directions, nodeDict := loadInput(input)

	moveCount := 0
	presentLocations := []node{}
	for k, v := range nodeDict {
		if k[len(k)-1] == 'A' {
			presentLocations = append(presentLocations, *v)
		}
	}
	fmt.Println(presentLocations)

	for !allEndsInZ(presentLocations) {
		nextDirection := directions[moveCount%len(directions)]
		if nextDirection == 'R' {
			presentLocations = updateAllRight(presentLocations)
		} else {
			presentLocations = updateAllLeft(presentLocations)
		}
		fmt.Println(presentLocations)
		moveCount += 1
	}

	return moveCount
}

func updateAllLeft(presentLocations []node) []node {
	newLocations := []node{}
	for _, location := range presentLocations {
		newLocations = append(newLocations, *location.leftNode)
	}
	return newLocations
}

func updateAllRight(presentLocations []node) []node {
	newLocations := []node{}
	for _, location := range presentLocations {
		newLocations = append(newLocations, *location.rightNode)
	}
	return newLocations
}

func allEndsInZ(presentLocations []node) bool {
	for _, location := range presentLocations {
		if location.value[len(location.value)-1] != 'Z' {
			return false
		}
	}
	return true
}

// adds node to map if not already present, returns node
func getNode(nodeMap map[string]*node, nodeString string) *node {
	existingNode, ok := nodeMap[nodeString]
	if ok {
		return existingNode
	}
	newNode := node{nodeString, nil, nil}
	nodeMap[nodeString] = &newNode
	return &newNode
}
