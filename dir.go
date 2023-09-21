package main

var (
	None      = Direction{varX: 0, varY: 0}
	North     = Direction{varX: 0, varY: -1}
	South     = Direction{varX: 0, varY: 1}
	East      = Direction{varX: 1, varY: 0}
	West      = Direction{varX: -1, varY: 0}
)

type Direction struct {
	varX, varY int
}

type DirectionQueue []Direction

func (d Direction) isOpposite(dir Direction) bool {
	return d.varX == -dir.varX && d.varY == -dir.varY
}

func (q *DirectionQueue) pop() Direction {
	if len(*q) == 0 {
		return None
	}
	dir := (*q)[0]
	*q = (*q)[1:]
	return dir
}

func (q *DirectionQueue) push(dir Direction) {
		*q = append(*q, dir)
}
