package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"sort"
	"strconv"
	"strings"
)

func inRange(value int, ranges [][2]int) bool {
	lo, hi := 0, len(ranges) - 1

	for lo <= hi {
		mid := (lo + hi) / 2
		if value < ranges[mid][0] {
			hi = mid - 1
		} else if value > ranges[mid][1] {
			lo = mid + 1
		} else {
			return true
		}
	}

	return false
}

func main() {
	file, err := os.Open("INPUT.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var ranges [][2]int
	var is_ranges bool = true
	var num_fresh int = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		if line == "" {
			is_ranges = false

			sort.Slice(ranges, func(i, j int) bool {
				return ranges[i][0] < ranges[j][0]
			})

			fmt.Printf("Original Ranges Length: %d\n", len(ranges))

			// Merge ranges
			merged := [][2]int{ranges[0]}
			for _, r := range ranges[1:] {
				last := &merged[len(merged)-1]
				if r[0] <= last[1] {
					if r[1] > last[1] {
						last[1] = r[1]
					}
				} else {
					merged = append(merged, r)
				}
			}

			ranges = merged

			fmt.Printf("Merged Ranges Length: %d\n", len(ranges))

		} else if is_ranges {
			parts := strings.SplitN(line, "-", 2)
			start, _ := strconv.Atoi(parts[0])
			end, _ := strconv.Atoi(parts[1])
			ranges = append(ranges, [2]int{start, end})
		} else {
			val, _ := strconv.Atoi(line)
			
			if inRange(val, ranges) {
				num_fresh++
			}
		}
	}

	total := 0
	for _, r := range ranges {
		total += r[1] - r[0] + 1
	}

	fmt.Printf("Total Fresh: %v\n", num_fresh)
	fmt.Printf("Total Fresh IDs: %v\n", total)

}
