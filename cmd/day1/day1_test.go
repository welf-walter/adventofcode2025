package main

import (
	"testing"
)

func TestRight(t *testing.T) {
	dial := Dial(2)
	dial = right(dial, Dial(12))
	if dial != 14 {
		t.Errorf("Expected 14 but got %v", dial)
	}
	dial = right(dial, Dial(112))
	if dial != 26 {
		t.Errorf("Expected 26 but got %v", dial)
	}
}

func TestLeft(t *testing.T) {
	dial := Dial(2)
	dial = left(dial, Dial(12))
	if dial != 90 {
		t.Errorf("Expected 90 but got %v", dial)
	}
}
