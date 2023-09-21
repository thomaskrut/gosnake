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
	food    *Food
	grow     int = 40
	tileImg     *ebiten.Image
	foodImg	 *ebiten.Image
	snakeImg      *ebiten.Image
	ops      = &ebiten.DrawImageOptions{}
	dirQueue DirectionQueue
)

const (
	elementSize = 5
	screenWidth  = 320
	screenHeight = 240
)

func init() {
	setRandomSource(0)
	snake = newSnake()
	food = newFood()
	snakeImg = ebiten.NewImage(elementSize, elementSize)
	snakeImg.Fill(color.White)
	tileImg = ebiten.NewImage(elementSize, elementSize)
	tileImg.Fill(color.RGBA{120, 120, 120, 12})
	foodImg = ebiten.NewImage(elementSize, elementSize)
	foodImg.Fill(color.RGBA{255, 0, 0, 255})
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

	for _, e := range snake.head.getAllBodyElements()[1:] {
		if e.p.x == snake.head.p.x && e.p.y == snake.head.p.y {
			log.Fatal("Game Over")
		}
	}

	if snake.head.p.x == food.p.x && snake.head.p.y == food.p.y {
		food = newFood()
		grow += 10
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	for x := 0; x < 32; x++ {
		for y := 0; y < 24; y++ {
			ops.GeoM.Reset()
			ops.GeoM.Translate(float64(x*elementSize), float64(y*elementSize))
			screen.DrawImage(tileImg, ops)
		}
	}

	for _, e := range snake.head.getAllBodyElements() {
		ops.GeoM.Reset()
		ops.GeoM.Translate(float64(e.p.x), float64(e.p.y))
		screen.DrawImage(snakeImg, ops)
	}

	ops.GeoM.Reset()
	ops.GeoM.Translate(float64(food.p.x), float64(food.p.y))
	screen.DrawImage(foodImg, ops)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (sw, sh int) {
	return screenWidth, screenHeight
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Gosnake")
	if err := ebiten.RunGame(&Game{}); err != nil {
		log.Fatal(err)
	}
}
