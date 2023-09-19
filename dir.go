package main

var (
	None      = direction{varX: 0, varY: 0}
	North     = direction{varX: 0, varY: -1}
	South     = direction{varX: 0, varY: 1}
	East      = direction{varX: 1, varY: 0}
	West      = direction{varX: -1, varY: 0}
)

type direction struct {
	varX, varY int
}