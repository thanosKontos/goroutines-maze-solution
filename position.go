package main

import (
	"fmt"
)

type position struct {
	x int
	y int
}

func (p *position) getValue(maze [][]string) string {
	return maze[p.x][p.y]
}

func (p *position) outOfBounds(maze [][]string) bool {
	return p.x < 0 || p.y < 0 || p.x > len(maze)-1 || p.y > len(maze[1])-1
}

func (p *position) isWall(maze [][]string) bool {
	return p.getValue(maze) == "x"
}

func (p *position) isEmpty(maze [][]string) bool {
	return p.getValue(maze) == " "
}

func (p *position) String() string {
	return fmt.Sprintf("%v:%v", p.x, p.y)
}