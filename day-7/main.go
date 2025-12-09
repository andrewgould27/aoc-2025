package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	file, err := os.Open("INPUT.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()

	var grid [][]rune
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}

	// Find starting position S
	var startCol int
	for col, ch := range grid[0] {
		if ch == 'S' {
			startCol = col
			break
		}
	}

	// Track active beams: column -> number of timelines at that position
	beams := map[int]int{startCol: 1}

	rows := len(grid)
	cols := len(grid[0])

	// Process row by row, starting from row 1 (below S)
	for row := 1; row < rows; row++ {
		nextBeams := make(map[int]int)

		for col, timelines := range beams {
			ch := grid[row][col]

			if ch == '^' {
				// Each timeline splits into two
				leftCol := col - 1
				rightCol := col + 1

				if leftCol >= 0 {
					nextBeams[leftCol] += timelines
				}
				if rightCol < cols {
					nextBeams[rightCol] += timelines
				}
			} else {
				// Timelines continue downward
				nextBeams[col] += timelines
			}
		}

		beams = nextBeams
	}

	// Sum up all timelines
	totalTimelines := 0
	for _, timelines := range beams {
		totalTimelines += timelines
	}

	fmt.Println(totalTimelines)
}
