package main

import (
	"errors"
	"fmt"
	"sync"
	"os"
)

const (
	up = iota
	right
	down
	left
)
var posChan = make(chan position, 500)
var doneChan = make(chan []string)
var mutex sync.Mutex
var visited = make(map[string][]string)

func main() {
	filename := os.Args[1]
	maze := createMaze(filename)
	startPosition := maze.findStartPosition()

	visited[startPosition.String()] = append(visited[startPosition.String()], startPosition.String())
	sendScoutsToAllDirections(startPosition, maze)

	for {
		select {
		case position := <-posChan:
			if position.isEmpty(maze) {
				sendScoutsToAllDirections(position, maze)
			}
		case finalRoute := <-doneChan:
			fmt.Println(finalRoute)
			return
		}
	}
}

func sendScoutsToAllDirections(from position, maze mazeType) {
	if to, err := canMove(up, from, maze); err == nil {
		go move(from, to, maze)
	}

	if to, err := canMove(right, from, maze); err == nil {
		go move(from, to, maze)
	}

	if to, err := canMove(down, from, maze); err == nil {
		go move(from, to, maze)
	}

	if to, err := canMove(left, from, maze); err == nil {
		go move(from, to, maze)
	}
}

func move(from, to position, maze mazeType) {
	mutex.Lock()
	visited[to.String()] = append(visited[from.String()], to.String())
	mutex.Unlock()

	if to.getValue(maze) == mazeFinish {
		doneChan <- visited[to.String()]
	} else {
		sendScoutsToAllDirections(to, maze)
	}
}

func canMove(direction int, current position, maze mazeType) (position, error) {
	var newX, newY int

	switch direction {
	case up:
		newX, newY = current.x+1, current.y
	case right:
		newX, newY = current.x, current.y+1
	case down:
		newX, newY = current.x-1, current.y
	case left:
		newX, newY = current.x, current.y-1
	}

	newPosition := position{
		newX,
		newY,
	}

	if newPosition.outOfBounds(maze) {
		return position{}, errors.New("OutOfBounds")
	}

	if newPosition.isWall(maze) {
		return position{}, errors.New("Wall")
	}

	if newPosition.isVisited() {
		return position{}, errors.New("Visited")
	}

	return newPosition, nil
}

func (p *position) isVisited() bool {
	mutex.Lock()
	_, ok := visited[p.String()]
	mutex.Unlock()

	return ok
}
