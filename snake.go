package main

type Snake struct {
	head *BodyElement
	dir Direction
}

type BodyElement struct {
	p   *Point
	tail *BodyElement
}

func (e BodyElement) getAllBodyElements() []*BodyElement {
	elements := []*BodyElement{&e}
	if e.tail != nil {
		elements = append(elements, e.tail.getAllBodyElements()...)
	}
	return elements
}

func (e *BodyElement) getEndOfTail() *BodyElement {
	if e.tail == nil {
		return e
	}
	return e.tail.getEndOfTail()
}

func (e *BodyElement) move(p Point) {
	previous := *e.p
	e.p = &p
	if e.tail != nil {
		e.tail.move(previous)
	}
}

func newSnake() *Snake {
	newSnake := Snake{head: newBodyElement(0, 0), dir: East}
	return &newSnake
}

func newBodyElement(startX, startY int) *BodyElement {
	newElement := &BodyElement{p: newPoint(startX, startY), tail: nil}
	return newElement
}