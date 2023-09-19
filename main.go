package main

import (
	"image/color"
	"log"
	"github.com/hajimehoshi/ebiten/v2"
)

type BodyElement struct {
	rect *ebiten.Image
	ops *ebiten.DrawImageOptions
}

func newBodyElement() *BodyElement {
	newElement := &BodyElement{ops: &ebiten.DrawImageOptions{}, rect: ebiten.NewImage(10, 10)}
	newElement.rect.Fill(color.Gray16{})
	return newElement
}

var (
	snake = []*BodyElement{}
)

func init() {
	snake = append(snake, newBodyElement())
}

type Game struct{
	
}

func (g *Game) Update() error {
	for e := range snake {
		snake[e].ops.GeoM.Translate(1, 0)
	}
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.NRGBA{0x00, 0x80, 0x80, 0xff})
	for e := range snake {
		screen.DrawImage(snake[e].rect, snake[e].ops)
	}
	
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