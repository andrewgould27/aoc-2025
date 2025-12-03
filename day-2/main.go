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
		length := len(s)

		for size := 1; size <= length/2; size++ {
			if length % size == 0 {
				pattern := s[:size]
				matched := true
				for i := size; i < length; i += size {
					if s[i:i+size] != pattern {
						matched = false
						break
					}
				}

				if matched {
					sum += i
					break
				}
			}
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

		raw_ranges := strings.SplitSeq(line, ",")

		for raw_range := range raw_ranges {
			range_arr := strings.SplitN(raw_range, "-", 2)
			start, err := strconv.Atoi(range_arr[0])
			if err != nil {
				log.Fatal(err)
			}
			end, err := strconv.Atoi(range_arr[1])
			if err != nil {
				log.Fatal(err)
			}

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
