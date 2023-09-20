package main

import (
	"strconv"
)



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

func (q DirectionQueue) String() string {
	str := ""
	for _, dir := range q {
		str += "(x:" + strconv.Itoa(dir.varX) + " y:" + strconv.Itoa(dir.varY) + ")" 
	}
	return str
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
	if (len(*q) > 0 && (*q)[len(*q)-1] != dir) || len(*q) == 0 {
		
		*q = append(*q, dir)

}
}