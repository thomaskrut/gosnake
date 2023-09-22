package main

import (
	"github.com/thomaskrut/gosnake/util"
)

type Food struct {
	p *util.Point
}

func newFood(game Game) *Food {
	newFood := Food{p: util.GetRandomPoint(game.screenWidth, game.screenHeight, game.elementSize)}
	return &newFood
}