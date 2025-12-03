package main

import (
	"fmt"
	"os"
	"bufio"
	"log"
	"strconv"
)

func HandleLine(s string) int {
	result := make([]byte, 0, 12)

	for i := 0; i < 12; i++ {
		max_idx := len(s) - (12 - len(result))
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

	var cumsum int = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		sum := HandleLine(line)
		cumsum += sum
	}

	fmt.Printf("Cumsum: %v\n", cumsum)
}
