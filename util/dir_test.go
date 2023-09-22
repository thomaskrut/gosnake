package util

import (
	"testing"
)

func TestDirectionQueuePop(t *testing.T) {
	dirQueue := DirectionQueue{North, South, East, West}
	if dir := dirQueue.Pop(); dir != North {
		t.Errorf("Expected North, got %v", dir)
	}
	if dir := dirQueue.Pop(); dir != South {
		t.Errorf("Expected South, got %v", dir)
	}
	if dir := dirQueue.Pop(); dir != East {
		t.Errorf("Expected East, got %v", dir)
	}
	if dir := dirQueue.Pop(); dir != West {
		t.Errorf("Expected West, got %v", dir)
	}
	if dir := dirQueue.Pop(); dir != None {
		t.Errorf("Expected None, got %v", dir)
	}
}