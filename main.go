package main

import (
	"github.com/hajimehoshi/ebiten/v2"
	"image/color"
	"log"
)

type BodyElement struct {
	x, y int
	dir  direction
	rect *ebiten.Image
	ops  *ebiten.DrawImageOptions
	tail *BodyElement
}

func (e *BodyElement) getEndOfTail() *BodyElement {
	if e.tail == nil {
		return e
	}
	return e.tail.getEndOfTail()
}

func (e *BodyElement) move() {
	
}

func newBodyElement(startX, startY int) *BodyElement {
	newElement := &BodyElement{x: startX, y: startY, dir: East, ops: &ebiten.DrawImageOptions{}, rect: ebiten.NewImage(8, 8)}
	newElement.rect.Fill(color.RGBA{255, 0, 255, 255})
	return newElement
}

var (
	head *BodyElement
	grow int
)

func init() {
	head = newBodyElement(10, 10)
	
}

func getNewDirection(currentDir direction) direction {
	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		return North
	} else if ebiten.IsKeyPressed(ebiten.KeyDown) {
		return South
	} else if ebiten.IsKeyPressed(ebiten.KeyLeft) {
		return West
	} else if ebiten.IsKeyPressed(ebiten.KeyRight) {
		return East
	} else if ebiten.IsKeyPressed(ebiten.KeySpace) {
		grow++
	}
	return currentDir
}

type Game struct {
}

func (g *Game) Update() error {
	if grow > 0 {
		grow--
		tail := head.getEndOfTail()
		tail.tail = newBodyElement(tail.x, tail.y)
	}
	head.dir = getNewDirection(head.dir)
	head.ops.GeoM.Reset()
	head.x += head.dir.varX
	head.y += head.dir.varY
	head.ops.GeoM.Translate(float64(head.x), float64(head.y))
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.DrawImage(head.rect, head.ops)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Hello, World!")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
