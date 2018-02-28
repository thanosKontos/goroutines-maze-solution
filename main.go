package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"sync"
)

func main() {
	names := []string{"1 0"}
	visited["1 0"] = append(visited["1 0"], "1 0")
	sendScoutsToAllDirections(position{1, 0, false, names})

	for {
		select {
		case position := <-posChan:
			if position.isEmpty() {
				sendScoutsToAllDirections(position)
			}
		case <-doneChan:
			fmt.Println("quit")
			return
		}
	}
}

var posChan = make(chan position, 500)
var doneChan = make(chan struct{})

var mutex sync.Mutex

const (
	up = iota
	right
	down
	left
)

type position struct {
	x       int
	y       int
	visited bool
	route   []string
}

var maze = [][]string{
	{"x", " ", "x", "x", " "},
	{"S", " ", "x", " ", " "},
	{"x", " ", " ", " ", " "},
	{"x", " ", " ", "x", " "},
	{" ", " ", "x", "F", " "},
}

var visited = make(map[string][]string)

func sendScoutsToAllDirections(from position) {
	if to, err := canMove(up, from); err == nil {
		go move(from, to)
	}

	if to, err := canMove(right, from); err == nil {
		go move(from, to)
	}

	if to, err := canMove(down, from); err == nil {
		go move(from, to)
	}

	if to, err := canMove(left, from); err == nil {
		go move(from, to)
	}
}

func move(from, to position) {
	keyFrom := coordinatesToString(from.x, from.y)
	keyTo := coordinatesToString(to.x, to.y)

	mutex.Lock()
	visited[keyTo] = append(visited[keyFrom], keyTo)
	mutex.Unlock()

	//fmt.Println("Visited " + keyTo)

	if to.getValue() == "F" {
		fmt.Println(strings.Join(visited[keyTo], "\n"))
		doneChan <- struct{}{}
	} else {
		sendScoutsToAllDirections(to)
	}
}

func canMove(direction int, current position) (position, error) {
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
		false,
		current.route,
	}

	if newPosition.outOfBounds() {
		return position{}, errors.New("OutOfBounds")
	}

	if newPosition.isWall() {
		return position{}, errors.New("Wall")
	}

	if newPosition.isVisited() {
		return position{}, errors.New("Visited")
	}

	newPosition.route = append(current.route, coordinatesToString(newX, newY))

	return newPosition, nil
}

func (p *position) getValue() string {
	return maze[p.x][p.y]
}

func (p *position) outOfBounds() bool {
	return p.x < 0 || p.y < 0 || p.x > len(maze)-1 || p.y > len(maze[1])-1
}

func (p *position) isWall() bool {
	return maze[p.x][p.y] == "x"
}

func (p *position) isEmpty() bool {
	return maze[p.x][p.y] == " "
}

func (p *position) isVisited() bool {
	mutex.Lock()
	_, ok := visited[coordinatesToString(p.x, p.y)]
	mutex.Unlock()

	return ok
}

func coordinatesToString(x, y int) string {
	return strconv.Itoa(x) + " " + strconv.Itoa(y)
}
