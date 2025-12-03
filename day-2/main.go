package main

import (
	"fmt"
	"strings"
	"os"
	"log"
	"bufio"
	"strconv"
)

func HandleRange(start, end int) int {
	var sum int = 0

	for i := start; i <= end; i++ {
		s := strconv.Itoa(i)
		if len(s) % 2 != 0 {
			continue
		}

		half := len(s) / 2
		
		if s[:half] == s[half:] {
			sum += i
		}
	}
	return sum
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

		raw_ranges := strings.Split(line, ",")

		for _, raw_range := range raw_ranges {
			range_arr := strings.SplitN(raw_range, "-", 2)
			start, err := strconv.Atoi(range_arr[0])
			if err != nil {
				log.Fatal(err)
			}
			end, err := strconv.Atoi(range_arr[1])
			if err != nil {
				log.Fatal(err)
			}

			fmt.Printf("(%v, %v)\n", start, end)

			sum := HandleRange(start, end)
			cumsum += sum
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Cumsum: %v\n", cumsum)
	fmt.Println("Done.")
}
