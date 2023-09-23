package main

import (
	"strconv"

	"github.com/thomaskrut/gosnake/util"
)

type Maze struct {
	walls []*Wall
	game  *Game
}

type Wall struct {
	points []util.Point
}

func (m *Maze) addWall(wall *Wall) {
	m.walls = append(m.walls, wall)
}

func (m *Maze) addRandomWall() {
	m.addWall(newWall(*m.game, util.RandPoint(m.game.screenWidth, m.game.screenHeight, m.game.elementSize), util.RandomNumber(10), util.GetRandomDirection()))
}

func (m *Maze) loadFromFile(level string) {
	data := util.ReadCsv("levels/" + level + ".lvl")
	for _, row := range data {
		if len(row) != 4 {
			panic("Invalid level file")
		}

		x, err := strconv.Atoi(row[0])
		if err != nil {
			panic("Invalid x coordinate")
		}

		y, err := strconv.Atoi(row[1])
		if err != nil {
			panic("Invalid y coordinate")
		}

		len, err := strconv.Atoi(row[2])
		if err != nil {
			panic("Invalid length")
		}

		dir, err := util.GetDirFromString(row[3])
		if err != nil {
			panic("Invalid direction")
		}

		m.addWall(newWall(*m.game, util.NewPoint(x*game.elementSize, y*game.elementSize), len, dir))
	}
}

func (m Maze) Points() []util.Point {
	points := make([]util.Point, 0)
	for _, wall := range m.walls {
		points = append(points, wall.getPoints()...)
	}
	return points
}

func (w Wall) getPoints() []util.Point {
	return w.points
}

func newMaze(g *Game) *Maze {
	return &Maze{game: g}
}

func newWall(game Game, startingPoint util.Point, length int, direction util.Direction) *Wall {
	points := make([]util.Point, length)
	points[0] = startingPoint
	for i := 1; i < length; i++ {
		points[i] = startingPoint.GetAdjecentPoint(direction, game.elementSize)
		startingPoint = points[i]
	}
	return &Wall{points: points}
}
