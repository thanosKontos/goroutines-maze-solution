package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var testMaze = [][]string{
	{"x", " ", "x", "x", " "},
	{"S", " ", "x", " ", " "},
	{"x", " ", " ", " ", " "},
	{"x", " ", " ", "x", " "},
	{" ", " ", "x", "F", " "},
}

func TestPositionValues(t *testing.T) {
	cases := map[string]position{
		"S": position{1, 0},
		" ": position{3, 1},
		"x": position{4, 2},
		"F": position{4, 3},
	}

	for expectedVal, aPosition := range cases {
		assert.Equal(t, expectedVal, aPosition.getValue(testMaze))
	}
}

func TestPositionOutOfBounds(t *testing.T) {
	cases := [...]position{
		position{-1, 1},
		position{1, -2},
		position{1, 5},
		position{5, 1},
	}

	for _, aPosition := range cases {
		assert.True(t, aPosition.outOfBounds(testMaze))
	}
}

func TestEmptyValue(t *testing.T) {
	cases := [...]position{
		position{0, 1},
		position{4, 4},
	}

	for _, aPosition := range cases {
		assert.True(t, aPosition.isEmpty(testMaze))
	}
}

func TestWall(t *testing.T) {
	cases := [...]position{
		position{0, 0},
		position{1, 2},
	}

	for _, aPosition := range cases {
		assert.True(t, aPosition.isWall(testMaze))
	}
}

func TestPositionToString(t *testing.T) {
	cases := map[string]position{
		"1:0":  position{1, 0},
		"3:1":  position{3, 1},
		"4:3":  position{4, 3},
		"14:3": position{14, 3},
	}

	for expectedVal, aPosition := range cases {
		assert.Equal(t, expectedVal, aPosition.String())
	}
}
