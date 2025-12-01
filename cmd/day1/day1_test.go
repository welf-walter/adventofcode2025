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
