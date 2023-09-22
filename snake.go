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

func (e BodyElement) getPoint() *util.Point {
	return e.p
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

func (e *BodyElement) append(element *BodyElement) {
	e.tail = element
}

func (e *BodyElement) move(p util.Point) {
	previous := *e.p
	e.p = &p
	if e.tail != nil {
		e.tail.move(previous)
	}
}

func newSnake() *Snake {
	newSnake := Snake{head: newBodyElement(util.NewPoint(0, 0)), dir: util.East}
	return &newSnake
}

func newBodyElement(p util.Point) *BodyElement {
	newElement := &BodyElement{p: &p, tail: nil}
	return newElement
}
