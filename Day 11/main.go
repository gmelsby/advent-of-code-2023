package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
)

type coordinates struct {
	x int
	y int
}

func main() {
	distance := solveBigGalaxyPuzzle("input.txt", 2)
	fmt.Println(distance)
	distance2 := solveBigGalaxyPuzzle("input.txt", 1000000)
	fmt.Println(distance2)
}

func solveGalaxyPuzzle(input string) int {
	readFile, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	grid := [][]rune{}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	for fileScanner.Scan() {
		line := fileScanner.Text()
		grid = append(grid, []rune(line))
		allEmpty := true
		for _, r := range grid[len(grid)-1] {
			if r != '.' {
				allEmpty = false
				break
			}
		}

		if allEmpty {
			grid = append(grid, []rune(line))
		}
	}

	readFile.Close()

	i := 0
	for i < len(grid[0]) {
		allEmpty := true
		for j := 0; j < len(grid); j++ {
			if grid[j][i] != '.' {
				allEmpty = false
				i++
				break
			}
		}
		if allEmpty {
			for j := 0; j < len(grid); j++ {
				grid[j] = append(append(grid[j][:i], '.'), grid[j][i:]...)
			}
			i += 2
		}
	}

	visited := []coordinates{}
	distance := 0
	for y := 0; y < len(grid); y++ {
		for x := 0; x < len(grid[0]); x++ {
			if grid[y][x] == '#' {
				for _, coord := range visited {
					distance += int(math.Abs(float64(x - coord.x)))
					distance += int(math.Abs(float64(y - coord.y)))
				}
				visited = append(visited, coordinates{x, y})
			}
		}
	}

	return distance
}

func solveBigGalaxyPuzzle(input string, scaleFactor int) int {
	readFile, err := os.Open(input)
	if err != nil {
		panic(err)
	}

	grid := [][]rune{}
	galaxies := []coordinates{}
	fileScanner := bufio.NewScanner(readFile)
	fileScanner.Split(bufio.ScanLines)
	y := 0
	for fileScanner.Scan() {
		line := fileScanner.Text()
		grid = append(grid, []rune(line))
		allEmpty := true
		for x, r := range grid[len(grid)-1] {
			if r != '.' {
				allEmpty = false
				galaxies = append(galaxies, coordinates{x, y})
			}
		}
		if allEmpty {
			y += scaleFactor
		} else {
			y++
		}
	}

	readFile.Close()

	for i := len(grid[0]) - 1; i >= 0; i-- {
		allEmpty := true
		for j := 0; j < len(grid); j++ {
			if grid[j][i] != '.' {
				allEmpty = false
			}
		}
		if allEmpty {
			for idx, galaxy := range galaxies {
				if galaxy.x > i {
					galaxies[idx].x += scaleFactor - 1
				}
			}
		}
	}

	distance := 0
	for idx, galaxy := range galaxies {
		for _, visitedGalaxy := range galaxies[:idx] {
			distance += int(math.Abs(float64(galaxy.x - visitedGalaxy.x)))
			distance += int(math.Abs(float64(galaxy.y - visitedGalaxy.y)))
		}
	}

	return distance
}
