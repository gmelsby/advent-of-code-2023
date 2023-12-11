package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strings"
)

type direction int

const (
	left = iota
	right
	up
	down
)

var dirDict = map[direction]coordinate{
	left:  {-1, 0},
	right: {1, 0},
	up:    {0, -1},
	down:  {0, 1},
}

var pipeDict = map[rune]map[direction]direction{
	'|': {
		up:   up,
		down: down,
	},
	'-': {
		left:  left,
		right: right,
	},
	'L': {
		down: right,
		left: up,
	},
	'J': {
		down:  left,
		right: up,
	},
	'7': {
		right: down,
		up:    left,
	},
	'F': {
		left: down,
		up:   right,
	},
}

func main() {
	distance, area := solveLoopPuzzle("input.txt")
	fmt.Printf("distance: %d, area: %d\n", distance, area)
}

type coordinate struct {
	x int
	y int
}

type mover struct {
	coords coordinate
	dir    direction
}

func solveLoopPuzzle(input string) (int, int) {
	readFile, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	grid := []string{}
	startCoordinate := coordinate{-1, -1}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	lineNumber := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		grid = append(grid, line)
		if startCoordinate.x == -1 {
			startCoordinate.x = strings.IndexRune(line, 'S')
			startCoordinate.y = lineNumber
		}
		lineNumber++
	}
	readFile.Close()

	pipeLocs := map[int][]int{}

	movers := []*mover{}
	for _, dir := range []direction{left, right, up, down} {
		newCoord := coordinate{startCoordinate.x + dirDict[dir].x, startCoordinate.y + dirDict[dir].y}
		if newCoord.x < 0 {
			continue
		}
		if dirMap, ok := pipeDict[rune(grid[newCoord.y][newCoord.x])]; ok {
			if _, ok := dirMap[dir]; ok {
				movers = append(movers, &mover{newCoord, dir})
			}
		}
	}

	sValue := "S"
	switch movers[0].dir {
	case up:
		switch movers[1].dir {
		case left:
			sValue = "J"
		case right:
			sValue = "L"
		case down:
			sValue = "|"
		}
	case down:
		switch movers[1].dir {
		case left:
			sValue = "7"
		case right:
			sValue = "F"
		case up:
			sValue = "|"
		}
	case left:
		switch movers[1].dir {
		case down:
			sValue = "7"
		case up:
			sValue = "J"
		case right:
			sValue = "-"
		}
	case right:
		switch movers[1].dir {
		case down:
			sValue = "F"
		case up:
			sValue = "L"
		case left:
			sValue = "-"
		}
	}

	grid[startCoordinate.y] = strings.Replace(grid[startCoordinate.y], "S", sValue, 1)
	if sValue != "-" {
		pipeLocs[startCoordinate.y] = append(pipeLocs[startCoordinate.y], startCoordinate.x)
	}

	distance := 1

	for movers[0].coords != movers[1].coords {
		for _, m := range movers {
			if grid[m.coords.y][m.coords.x] != '-' {
				pipeLocs[m.coords.y] = append(pipeLocs[m.coords.y], m.coords.x)
			}
			m.dir = pipeDict[rune(grid[m.coords.y][m.coords.x])][m.dir]
			m.coords.x += dirDict[m.dir].x
			m.coords.y += dirDict[m.dir].y
		}
		distance += 1
	}

	m := movers[0]
	pipeLocs[m.coords.y] = append(pipeLocs[m.coords.y], m.coords.x)

	area := 0
	for line, intersections := range pipeLocs {
		sort.Ints(intersections)
		inside := false
		lineStart := '-'
		for i := 0; i < len(intersections)-1; i++ {
			char := rune(grid[line][intersections[i]])

			if char == 'F' || char == 'L' {
				lineStart = char
				continue
			}

			if char == '|' {
				inside = !inside
			}

			if char == 'J' {
				if lineStart == 'F' {
					inside = !inside
				}
			} else if char == '7' {
				if lineStart == 'L' {
					inside = !inside
				}
			}

			if inside {
				area += intersections[i+1] - intersections[i] - 1
			}
		}
	}
	return distance, area
}
