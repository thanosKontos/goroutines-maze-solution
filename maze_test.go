package main

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMazeCreation(t *testing.T) {
	maze := createMaze("mazes/maze1.txt")

	assert.Equal(t, "x", maze[0][3])
	assert.Equal(t, " ", maze[1][3])
}

func TestInvalidMazeCreationPanics(t *testing.T) {
	assert.Panics(t, func(){ createMaze("mazes/non_existing.txt") })
}

func TestFindMazeStart(t *testing.T) {
	maze := createMaze("mazes/maze1.txt")
	assert.Equal(t, position{1, 0}, maze.findStartPosition())
}

func TestMazeWithoutStartPanics(t *testing.T) {
	maze := createMaze("mazes/maze_no_start.txt")
	assert.Panics(t, func(){ maze.findStartPosition() })
}
