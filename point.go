package main

type Point struct {
	x, y int
}

func newPoint(x, y int) *Point {
	return &Point{x: x, y: y}
}

func (p Point) move(dir Direction) Point {
	x := p.x + dir.varX
	y := p.y + dir.varY
	return Point{x: x, y: y}
}



