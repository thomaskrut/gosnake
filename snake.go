package main

import (
	"github.com/thomaskrut/gosnake/util"
)

type Snake struct {
	head *BodyElement
	dir  util.Direction
}

type BodyElement struct {
	p    *util.Point
	tail *BodyElement
}

func (s Snake) Points() []util.Point {
	body := s.head.body()
	points := make([]util.Point, len(body))
	for i, element := range body {
		points[i] = *element.p
	}
	return points

}

func (e BodyElement) point() *util.Point {
	return e.p
}

func (e BodyElement) body() []*BodyElement {
	elements := []*BodyElement{&e}
	if e.tail != nil {
		elements = append(elements, e.tail.body()...)
	}
	return elements
}

func (s Snake) tail() *BodyElement {
	return s.head.getEndOfTail()
}

func (e *BodyElement) grow() {
	e.tail = newBodyElement(*e.point())
}

func (e BodyElement) isOnGrid() bool {
	return e.point().IsOnGrid(elementSize)
}

func (e *BodyElement) getEndOfTail() *BodyElement {
	if e.tail == nil {
		return e
	}
	return e.tail.getEndOfTail()
}

func (e *BodyElement) move(p util.Point) {
	previous := *e.p
	e.p = &p
	if e.tail != nil {
		e.tail.move(previous)
	}
}

func newSnake(pos util.Point, dir util.Direction) *Snake {
	newSnake := Snake{head: newBodyElement(pos), dir: dir}
	return &newSnake
}

func newBodyElement(p util.Point) *BodyElement {
	newElement := &BodyElement{p: &p, tail: nil}
	return newElement
}
