package main

import (
	"testing"
	"github.com/stretchr/testify/assert"
)

func TestNormalLeftMove(t *testing.T) {
	assert := assert.New(t)

	var curr_dial int = 19
	var direction rune = 'L'
	var distance int = 19

	Move(&curr_dial, direction, distance)

	assert.Equal(curr_dial, 0, "Dial should be set to 0")
}

func TestNormalRightMove(t *testing.T) {
	assert := assert.New(t)

	var curr_dial int = 11
	var direction rune = 'R'
	var distance int = 8

	Move(&curr_dial, direction, distance)

	assert.Equal(curr_dial, 19, "Dial should be set to 19")
}

func TestWrapLeft(t *testing.T) {
	assert := assert.New(t)

	curr_dial := 0
	direction := 'L'
	distance := 1

	Move(&curr_dial, direction, distance)

	assert.Equal(curr_dial, 99, "Dial should be set to 99")
}

func TestWrapRight(t *testing.T) {
	assert := assert.New(t)

	curr_dial := 99
	direction := 'R'
	distance := 1

	Move(&curr_dial, direction, distance)

	assert.Equal(curr_dial, 0, "Dial should be set to 0")
}

func TestRightRotation(t *testing.T) {
	assert := assert.New(t)

	curr_dial := 50
	direction := 'R'
	distance := 1000

	rotations := Move(&curr_dial, direction, distance)

	assert.Equal(rotations, 10, "Should have clicked past zero 10 times")
}

func TestLeftRotation(t *testing.T) {
	assert := assert.New(t)

	curr_dial := 50
	direction := 'L'
	distance := 1000

	rotations := Move(&curr_dial, direction, distance)

	assert.Equal(rotations, 10, "Should have clicked past zero 10 times")
}
