package main

import (
	"testing"
)

func TestDirectionQueuePop(t *testing.T) {
	dirQueue = DirectionQueue{North, South, East, West}
	if dirQueue.pop() != North {
		t.Errorf("Expected North, got %v", dirQueue.pop())
	}
	if dirQueue.pop() != South {
		t.Errorf("Expected South, got %v", dirQueue.pop())
	}
	if dirQueue.pop() != East {
		t.Errorf("Expected East, got %v", dirQueue.pop())
	}
	if dirQueue.pop() != West {
		t.Errorf("Expected West, got %v", dirQueue.pop())
	}
	if dirQueue.pop() != None {
		t.Errorf("Expected None, got %v", dirQueue.pop())
	}
	dirQueue.push(North)
	if dirQueue.pop() != North {
		t.Errorf("Expected North, got %v", dirQueue.pop())
	}
}