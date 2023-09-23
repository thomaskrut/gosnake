package main

import (
	"github.com/thomaskrut/gosnake/util"
)

type Food struct {
	p util.Point
}

func newFood(game *Game, points ...util.Points) *Food {
	newFood := Food{p: util.EmptyRandPoint(game.screenWidth, game.screenHeight, game.elementSize, points)}
	return &newFood
}
