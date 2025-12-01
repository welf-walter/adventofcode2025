package main

import (
	"testing"
)

func TestRotate(t *testing.T) {
	dial := Dial(11)
	dial = right(dial, Dial(8))
	if dial != 19 {
		t.Errorf("Expected 19 but got %v", dial)
	}
	dial = left(dial, Dial(19))
	if dial != 0 {
		t.Errorf("Expected 0 but got %v", dial)
	}
}

func TestRotateCircle(t *testing.T) {
	dial := Dial(0)
	dial = left(dial, Dial(1))
	if dial != 99 {
		t.Errorf("Expected 99 but got %v", dial)
	}
	dial = right(dial, Dial(1))
	if dial != 0 {
		t.Errorf("Expected 0 but got %v", dial)
	}
}

func TestPart1(t *testing.T) {
	input := `L68
L30
R48
L5
R60
L55
L1
L99
R14
L82`
	operations := parseInput(input)
	exp1 := Operation{direction: LEFT, ticks: Dial(30)}
	if operations[1] != exp1 {
		t.Errorf("Expected %v but got %v", exp1, operations[1])
	}
	exp2 := Operation{direction: RIGHT, ticks: Dial(48)}
	if operations[2] != exp2 {
		t.Errorf("Expected %v but got %v", exp1, operations[2])
	}
}
