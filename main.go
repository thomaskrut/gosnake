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
	head = newBodyElement()
	body = []*BodyElement{}
)

func init() {
	body = append(body, newBodyElement())
}

type Game struct{
	
}

func (g *Game) Update() error {
	head.ops.GeoM.Translate(1, 0)
	
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.NRGBA{0x00, 0x80, 0x80, 0xff})
	screen.DrawImage(head.rect, head.ops)
	for e := range body {
		screen.DrawImage(body[e].rect, body[e].ops)
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