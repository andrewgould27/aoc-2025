package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"strings"
	"strconv"
)

func part_1() {
	file, err := os.Open("INPUT.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	var lines []string 
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	var rows [][]string
	for _, line := range lines {
		rows = append(rows, strings.Fields(line))
	}

	ops := rows[len(rows) - 1]
	numRows := rows[:len(rows) - 1]
	var results []int

	for col, op := range ops {
		var nums []int
		for _, row := range numRows {
			n, _ := strconv.Atoi(row[col])
			nums = append(nums, n)
		}

		result := nums[0]
		for _, n := range nums[1:] {
			switch op {
			case "+":
				result += n
			case "*":
				result *= n
			}
		}

		results = append(results, result)
	}

	sum := 0
	for _, r := range results {
		sum += r
	}
	fmt.Printf("Total Sum: %v\n", sum)
}

func part_2() {
    file, err := os.Open("INPUT.txt")
    if err != nil {
        log.Fatal(err)
    }
    defer file.Close()

    scanner := bufio.NewScanner(file)
    var lines []string
    for scanner.Scan() {
        lines = append(lines, scanner.Text())
    }

    // Find max length and pad all lines to same length
    maxLen := 0
    for _, line := range lines {
        if len(line) > maxLen {
            maxLen = len(line)
        }
    }
    for i := range lines {
        for len(lines[i]) < maxLen {
            lines[i] += " "
        }
    }

    opRow := lines[len(lines)-1]
    numRows := lines[:len(lines)-1]

    var total int

    // Find problems by locating operators and scanning left to space-only columns
    col := maxLen - 1
    for col >= 0 {
        isSpace := true
        for _, row := range numRows {
            if col < len(row) && row[col] != ' ' {
                isSpace = false
                break
            }
        }
        if opRow[col] == ' ' && isSpace {
            col--
            continue
        }

        op := ""
        opCol := col
        for opCol >= 0 && op == "" {
            if opRow[opCol] == '+' || opRow[opCol] == '*' {
                op = string(opRow[opCol])
            } else {
                opCol--
            }
        }

        leftCol := opCol
        for leftCol > 0 {
            allSpace := opRow[leftCol-1] == ' '
            for _, row := range numRows {
                if row[leftCol-1] != ' ' {
                    allSpace = false
                    break
                }
            }
            if allSpace {
                break
            }
            leftCol--
        }

        var nums []int
        for c := leftCol; c <= col; c++ {
            numStr := ""
            for _, row := range numRows {
                ch := row[c]
                if ch >= '0' && ch <= '9' {
                    numStr += string(ch)
                }
            }
            if numStr != "" {
                n, _ := strconv.Atoi(numStr)
                nums = append(nums, n)
            }
        }

        if len(nums) > 0 {
            result := nums[0]
            for _, n := range nums[1:] {
                switch op {
                case "+":
                    result += n
                case "*":
                    result *= n
                }
            }
            total += result
        }

        col = leftCol - 1
    }

    fmt.Printf("Total: %d\n", total)
}

func main() {
	part_1()
	part_2()
}
