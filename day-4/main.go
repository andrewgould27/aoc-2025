package main

import (
	"fmt"
	"bufio"
	"os"
	"log"
)

var directions = [][2]int {
	{-1, -1}, {-1, 0}, {-1, 1},
    {0, -1},           {0, 1},
	{1, -1},  {1, 0},  {1, 1},
}

func CountNeighbors(grid [][]rune, row, col int) (ats int) {
	rows := len(grid)
	cols := len(grid[0])

	for _, d := range directions {
		newRow, newCol := row+d[0], col+d[1]
		if newRow >= 0 && newRow < rows && newCol >= 0 && newCol < cols {
			switch grid[newRow][newCol] {
			case '@':
				ats++
			case 'x':
				ats++
			}
		}
	}

	return ats
}

func SetXs(grid[][] rune) {
	for i := range len(grid) {
		for j := range len(grid[0]) {
			if grid[i][j] == 'x' {
				grid[i][j] = '.'
			}
		}
	}
}

func main() {
	file, err := os.Open("INPUT.txt")
	if err != nil {
		log.Fatal("Error reading file", err)
	}
	defer file.Close()

	var grid [][]rune

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		grid = append(grid, []rune(scanner.Text()))
	}
	
	var cumsum int = 0
	var first_pass int = 0
	for true {
		total_reachable := 0
		for i := range len(grid) {
			for j := range len(grid[0]) {
				if grid[i][j] == '@' {
					ats := CountNeighbors(grid, i, j)

					if ats < 4 {
						total_reachable++
						grid[i][j] = 'x'
					}
				}
			}
		}

		SetXs(grid)

		if total_reachable == 0 {
			break
		}

		if first_pass == 0 {
			first_pass = total_reachable
		}

		cumsum += total_reachable
	}

	fmt.Printf("First Pass Reachable: %d\n", first_pass)
	fmt.Printf("Total Reachable: %d\n", cumsum)
}
