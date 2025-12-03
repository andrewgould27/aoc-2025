package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strconv"
)

func HandleLine(s string, n int) int {
	result := make([]byte, 0, n)

	for range n {
		max_idx := len(s) - (n - len(result))
		best_idx := 0

		for j := 1; j <= max_idx; j++ {
			if s[j] > s[best_idx] {
				best_idx = j
			}
		}

		result = append(result, s[best_idx])
		s = s[best_idx + 1:]
	}

	int_max, _ := strconv.Atoi(string(result))
	return int_max
}


func main() {
	file, err := os.Open("INPUT.txt")
	if err != nil {
		log.Fatal(err)
	}

	var p1_cumsum int = 0
	var p2_cumsum int = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		sum := HandleLine(line, 2)
		p1_cumsum += sum
		sum = HandleLine(line, 12)
		p2_cumsum += sum
	}

	fmt.Printf("Part 1 Cumsum: %v\n", p1_cumsum)
	fmt.Printf("Part 2 Cumsum: %v\n", p2_cumsum)
}
