package main

type Snake struct {
	head *BodyElement
	dir Direction
}

type BodyElement struct {
	x, y int
	tail *BodyElement
}

func (e *BodyElement) getEndOfTail() *BodyElement {
	if e.tail == nil {
		return e
	}
	return e.tail.getEndOfTail()
}

func (e *BodyElement) move() {
	
}

func newSnake() *Snake {
	newSnake := Snake{head: newBodyElement(0, 0), dir: East}
	return &newSnake
}

func newBodyElement(startX, startY int) *BodyElement {
	newElement := &BodyElement{x: startX, y: startY}
	return newElement
}