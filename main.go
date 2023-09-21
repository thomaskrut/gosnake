package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	_ "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

var (
	snake    *Snake
	grow     int
	tile     *ebiten.Image
	img      *ebiten.Image
	ops      = &ebiten.DrawImageOptions{}
	dirQueue DirectionQueue
)

const (
	elementSize = 10
)

func init() {
	snake = newSnake()
	img = ebiten.NewImage(elementSize, elementSize)
	img.Fill(color.White)
	tile = ebiten.NewImage(elementSize-1, elementSize-1)
	tile.Fill(color.RGBA{120, 120, 120, 12})
}

func checkKeys() {
	if inpututil.IsKeyJustPressed(ebiten.KeyUp) && snake.dir != South {
		dirQueue.push(North)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyDown) && snake.dir != North {
		dirQueue.push(South)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) && snake.dir != East {
		dirQueue.push(West)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyRight) && snake.dir != West {
		dirQueue.push(East)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		grow+=10
	}
}

type Game struct {
}

func (g *Game) Update() error {

	checkKeys()

	if grow > 0 {
		grow--
		snake.head.getEndOfTail().tail = newBodyElement(snake.head.getEndOfTail().p.x, snake.head.getEndOfTail().p.y)
	}

	if len(dirQueue) > 0 && snake.head.p.x%elementSize == 0 && snake.head.p.y%elementSize == 0 {
		snake.dir = dirQueue.pop()
	}

	snake.head.move(snake.head.p.move(snake.dir))
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	/*for x := 0; x < 32; x++ {
		for y := 0; y < 24; y++ {
			ops.GeoM.Reset()
			ops.GeoM.Translate(float64(x*elementSize), float64(y*elementSize))
			screen.DrawImage(tile, ops)
		}
	}*/
	currentElement := snake.head
	for {
		ops.GeoM.Reset()
		ops.GeoM.Translate(float64(currentElement.p.x), float64(currentElement.p.y))
		screen.DrawImage(img, ops)
		if currentElement.tail == nil {
			break
		}
		currentElement = currentElement.tail
	}
	
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
