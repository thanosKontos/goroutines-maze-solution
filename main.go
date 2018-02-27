package main

import (
	"fmt"
	"strconv"
	"time"
	"errors"
	"sync"
)

func main() {
	position := position{1, 0}

	sendScoutsToAllDirections(position)

	time.Sleep(time.Millisecond * 1) 
}

var mutex sync.Mutex

const (
	up = iota
	right
	down
	left
)

type position struct {
	x int
	y int
}

var maze = [][]string{
	{"x"," ","x","x"," "},
	{"S"," ","x"," "," "},
	{"x"," "," "," "," "},
	{"x"," "," ","x"," "},
	{" "," ","x","F"," "},
}

var visited = make(map[string]bool)

func sendScoutsToAllDirections(from position) {
	if newPosition, err := canMove(up, from); err == nil {
		go move(up, newPosition)
	}

	if newPosition, err := canMove(right, from); err == nil {
		go move(right, newPosition)
	}

	if newPosition, err := canMove(down, from); err == nil {
		go move(down, newPosition)
	}

	if newPosition, err := canMove(left, from); err == nil {
		go move(left, newPosition)
	}
}

func canMove(direction int, current position) (position, error) {
	var newPosition position;

	switch direction {
	case up:
		newPosition = position{
			current.x+1,
			current.y,
		}
	case right:
		newPosition = position{
			current.x,
			current.y+1,
		}
	case down:
		newPosition = position{
			current.x-1,
			current.y,
		}
	case left:
		newPosition = position{
			current.x,
			current.y-1,
		}
	}

	if (newPosition.outOfBounds()) {
		return position{}, errors.New("OutOfBounds")
	}

	if (newPosition.isWall()) {
		return position{}, errors.New("Wall")
	}

	if (newPosition.isVisited()) {
		return position{}, errors.New("Visited")
	}

	return newPosition, nil
}

func move(direction int, from position) {
	key := strconv.Itoa(from.x) + " " + strconv.Itoa(from.y)
	
	mutex.Lock()
	visited[key] = true
	mutex.Unlock()

	fmt.Println("Visited " + key)

	if (from.getValue() == "F") {
		fmt.Println("Finish found!")
		return
		// Send to done channel to notify everyone to stop
	} else {
		sendScoutsToAllDirections(from)
	}
}

func (p *position) getValue() string {
	return maze[p.x][p.y]
}

func (p *position) outOfBounds() bool {
	return p.x < 0 || p.y < 0 || p.x > 4 || p.y > 4
}

func (p *position) isWall() bool {
	return maze[p.x][p.y] == "x"
}

func (p *position) isVisited() bool {
	key := strconv.Itoa(p.x) + " " + strconv.Itoa(p.y)
	mutex.Lock()
	_, ok := visited[key];
	mutex.Unlock()

	return ok;
}
