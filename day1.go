package advent_of_code

import (
	"strconv"
)

const (
	NORTH = 0
	EAST = 1
	SOUTH = 2
	WEST = 3
)

func(man *Lostman) turn(direction byte) {
	if direction == 'R' {
		if man.orrientation == WEST {
			man.orrientation = NORTH
		} else {
			man.orrientation++
		}
	} else {
		if man.orrientation == NORTH {
			man.orrientation = WEST
		} else {
			man.orrientation--
		}
	}
}

func(man *Lostman) move() Point {
	switch man.orrientation {
		case NORTH:
			return Point{x : man.position.x, y : man.position.y + 1}
		case WEST:
			return Point{x : man.position.x - 1, y : man.position.y}
		case SOUTH:
			return Point{x : man.position.x, y : man.position.y - 1}
		case EAST:
			return Point{x : man.position.x + 1, y : man.position.y}
		default:
			return Point{x : 0, y : 0}
	}
}

func(man *Lostman) moveSteps(numsteps int, visited *List, partTwo bool) bool {
	for i := 0; i < numsteps; i++ {
		p := man.move()
		man.position = p
		if(partTwo && visited.Contains(p)) {
			return true
		} else {
			visited.Add(p)
		}
	}
	return false
}

func(man *Lostman) findHeadquarters(commands []string, partTwo bool) {
	visited := NewList()
	for _, command := range commands {
		man.turn(command[0])
		numsteps, _ := strconv.Atoi(command[1:len(command)])
		if(man.moveSteps(numsteps, visited, partTwo)) {
			return
		}
	}
}

func(man *Lostman) distance() int {
	return Abs(man.position.x) + Abs(man.position.y)
}

func Abs(n int) int {
	if n >= 0 {
		return n
	} else {
		return -n
	}
}

func NewLostman() *Lostman {
	return &Lostman{}
}
