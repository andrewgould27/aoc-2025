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

func mod(a, b int) int {
	return ((a % b) + b) % b
}


func ParseInstruction(input string) (rune, int) {
	direction := rune(input[0])
	distance, err := strconv.Atoi(input[1:])
	if err != nil {
		log.Fatal(err)
	}

	return direction, distance
}

func Move(curr_dial *int, dir rune, distance int) {
	if dir == 'L' {
		distance *= -1
	}

	*curr_dial = mod(*curr_dial + distance, 100)

	fmt.Printf("New pos: %v\n", *curr_dial)
}

func main() {
	file, err := os.Open("INPUT.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var dial_value int = 50
	var number_of_zero int = 0

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		direction, distance := ParseInstruction(line)
		Move(&dial_value, direction, distance)

		if (dial_value == 0) {
			number_of_zero++
		}
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Printf("Number of times the dial was zero: %v\n", number_of_zero)
}
