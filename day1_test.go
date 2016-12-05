package advent_of_code

import (
	"testing"
	"strings"
)

func TestExamples1(t *testing.T) {
	lostman1 := NewLostman()
	lostman2 := NewLostman()
	lostman3 := NewLostman()
	lostman4 := NewLostman()

	if lostman1.position.x != 0 && lostman1.position.y != 0 && lostman1.orrientation != NORTH {
		t.Errorf("Error initialising")
	}

	ex1_commands :=  strings.Split("R2, L3", ", ")
	ex2_commands :=  strings.Split("R2, R2, R2", ", ")
	ex3_commands :=  strings.Split("R5, L5, R5, R3", ", ")
	ex4_commands := strings.Split("R8, R4, R4, R8", ", ")

	lostman1.findHeadquarters(ex1_commands, false)
	if distance := lostman1.distance(); distance != 5 {
		t.Errorf("Got %v expected %v", distance, 5)
	}


	lostman2.findHeadquarters(ex2_commands, false)
	if distance := lostman2.distance(); distance != 2 {
		t.Errorf("Got %v expected %v", distance, 2)
	}



	lostman3.findHeadquarters(ex3_commands, false)
	if distance := lostman3.distance(); distance != 12 {
		t.Errorf("Got %v expected %v", distance, 12)
	}

	lostman4.findHeadquarters(ex4_commands, true)
	if distance := lostman4.distance(); distance != 4 {
		t.Errorf("Got %v expected %v", distance, 4)
	}
}

func TestDay1(t *testing.T) {

	// part1
	realy_lostman1 := NewLostman()

	commands := strings.Split("R5, R4, R2, L3, R1, R1, L4, L5, R3, L1, L1, R4, L2, R1, R4, R4, L2, L2, R4, L4, R1, R3, L3, L1, L2, R1, R5, L5, L1, L1, R3, R5, L1, R4, L5, R5, R1, L185, R4, L1, R51, R3, L2, R78, R1, L4, R188, R1, L5, R5, R2, R3, L5, R3, R4, L1, R2, R2, L4, L4, L5, R5, R4, L4, R2, L5, R2, L1, L4, R4, L4, R2, L3, L4, R2, L3, R3, R2, L2, L3, R4, R3, R1, L4, L2, L5, R4, R4, L1, R1, L5, L1, R3, R1, L2, R1, R1, R3, L4, L1, L3, R2, R4, R2, L2, R1, L5, R3, L3, R3, L1, R4, L3, L3, R4, L2, L1, L3, R2, R3, L2, L1, R4, L3, L5, L2, L4, R1, L4, L4, R3, R5, L4, L1, L1, R4, L2, R5, R1, R1, R2, R1, R5, L1, L3, L5, R2", ", ")
	realy_lostman1.findHeadquarters(commands, false)
	if distance := realy_lostman1.distance(); distance != 231 {
		t.Errorf("Got %v expected %v", distance, 231)
	}

	// part2
	realy_lostman2 := NewLostman()
	realy_lostman2.findHeadquarters(commands, true)
	if distance := realy_lostman2.distance(); distance != 147 {
		t.Errorf("Got %v expected %v", distance, 147)
	}

}


