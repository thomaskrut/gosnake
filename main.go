package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	_"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var (
	snake *Snake
	grow int
	tile *ebiten.Image
	img *ebiten.Image
	ops = &ebiten.DrawImageOptions{}
	dirQueue DirectionQueue
)

const (
	elementSize = 10
)

func init() {
	snake = newSnake()
	img = ebiten.NewImage(elementSize, elementSize)
	img.Fill(color.White)
	tile = ebiten.NewImage(elementSize - 1, elementSize -1)
	tile.Fill(color.RGBA{120,120,120,120})
}

func keys() (Direction, bool) {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		return North, true
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		return South, true
	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		return West, true
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		return East, true
	}
	return None, false
}

type Game struct {
}

func (g *Game) Update() error {

	if newDir, ok := keys(); ok {
		dirQueue.push(newDir)
	}

	if grow > 0 {
		grow--
		tail := snake.head.getEndOfTail()
		tail.tail = newBodyElement(tail.x, tail.y)
	}
	
	if len(dirQueue) > 0 && snake.head.x % elementSize == 0 && snake.head.y % elementSize == 0 {
		snake.dir = dirQueue.pop()
	}

	snake.head.x += snake.dir.varX
	snake.head.y += snake.dir.varY
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for x:=0; x<32; x++ {
		for y:=0; y < 24; y++ {
			ops.GeoM.Reset()
			ops.GeoM.Translate(float64(x * elementSize), float64(y * elementSize))
			screen.DrawImage(tile, ops)
		}
	}
	ops.GeoM.Reset()
	ops.GeoM.Translate(float64(snake.head.x), float64(snake.head.y))
	screen.DrawImage(img, ops)	
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Gosnake")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
