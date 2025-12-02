package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"fmt"
)

// You have a circular combination dial, with the numbers 0 to 99
// The file contains instructions on how to open the safe
// by giving you a direction: L (to lower numbers), R (to higher numbers)
// and a distance to turn it in that direction.
// e.g. if the dial is at 11, and the directino is R8, the number is 19

// The password is the number of times the dial is pointing at 0 after any
// rotation in the sequence.

// The dial starts by pointing at 50

// Part two requires me to keep track of how many times it clicks past zero
// not just when it lands on zero. 

func mod(a, b int) int {
	return ((a % b) + b) % b
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}


func ParseInstruction(input string) (rune, int) {
	direction := rune(input[0])
	distance, err := strconv.Atoi(input[1:])
	if err != nil {
		log.Fatal(err)
	}

	return direction, distance
}

func Move(curr_dial *int, dir rune, distance int) int {
	if dir == 'L' {
		distance *= -1
	}

	start := *curr_dial
	end := start + distance
	*curr_dial = mod(end, 100)

	var rotations int
	if distance >= 0 {
		rotations = end / 100
	} else {
		if end < 0 {
			rotations = (start - end) / 100
		} else {
			rotations = 0
		}
	}

	return rotations
}

func main() {
	file, err := os.Open("INPUT.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var dial_value int = 50
	var number_of_zero int = 0
	var clicked_past_zero int = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		direction, distance := ParseInstruction(line)
		rotations := Move(&dial_value, direction, distance)

		clicked_past_zero += rotations

		if (dial_value == 0) {
			number_of_zero++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Number of time clicked past zero: %v\n", clicked_past_zero)
	fmt.Printf("Number of times the dial was zero: %v\n", number_of_zero)
	fmt.Printf("0x434C49434B Password: %v\n", clicked_past_zero + number_of_zero)
}
