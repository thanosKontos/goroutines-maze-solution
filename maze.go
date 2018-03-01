package main

import (
	"bufio"
	"os"
	"strings"
)

const (
	mazeStart = "S"
	mazeFinish = "F"
)

type mazeType [][]string

func createMaze(filename string) mazeType {
	var maze mazeType

	file, err := os.Open(filename)
    if err != nil {
        panic("Cannot find maze file")
    }
    defer file.Close()

	scanner := bufio.NewScanner(file)
	i := 0
    for scanner.Scan() {
		maze = append(maze, strings.Split(scanner.Text(), ""))
		i++
    }

    if err := scanner.Err(); err != nil {
        panic("Cannot read maze file")
    }

	return maze
}

func (maze *mazeType) findStartPosition() position {
	for i, row := range *maze {
        for j, val := range row {
			if val == mazeStart {
				return position{i, j}
			}
		}
	}

	panic("Cannot find start")
}