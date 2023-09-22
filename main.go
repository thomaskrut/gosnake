package main

import (
	"image/color"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	_ "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/thomaskrut/gosnake/util"
)

var (
	game     *Game
	snake    *Snake
	food     *Food
	grow     int
	tileImg  *ebiten.Image
	foodImg  *ebiten.Image
	snakeImg *ebiten.Image
	ops      = &ebiten.DrawImageOptions{}
	dirQueue util.DirectionQueue
)

const (
	elementSize  = 5
	screenWidth  = 320
	screenHeight = 240
)

func init() {
	game = &Game{elementSize: elementSize, screenWidth: screenWidth, screenHeight: screenHeight}
	util.SetRandomSource(0)
	grow = 40
	snake = newSnake()
	food = newFood(*game)
	snakeImg = ebiten.NewImage(elementSize, elementSize)
	snakeImg.Fill(color.RGBA{255, 255, 255, 255})
	tileImg = ebiten.NewImage(elementSize, elementSize)
	tileImg.Fill(color.RGBA{120, 120, 120, 12})
	foodImg = ebiten.NewImage(elementSize, elementSize)
	foodImg.Fill(color.RGBA{255, 0, 0, 255})
}

func checkKeys() {

	if inpututil.IsKeyJustPressed(ebiten.KeyUp) && snake.dir != util.South {
		dirQueue.Push(util.North)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyDown) && snake.dir != util.North {
		dirQueue.Push(util.South)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyLeft) && snake.dir != util.East {
		dirQueue.Push(util.West)
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyRight) && snake.dir != util.West {
		dirQueue.Push(util.East)
	}
	if inpututil.IsKeyJustPressed(ebiten.KeySpace) {
		grow += 10
	}
}

type Game struct {
	elementSize  int
	screenWidth  int
	screenHeight int
}

func (g *Game) Update() error {

	checkKeys()

	if grow > 0 {
		grow--
		snake.head.getEndOfTail().append(newBodyElement(*snake.head.getEndOfTail().getPoint()))
	}

	if len(dirQueue) > 0 && snake.head.getPoint().IsOnGrid(elementSize) {
		snake.dir = dirQueue.Pop()
	}

	snake.head.move(snake.head.getPoint().Move(snake.dir))

	if snake.head.getPoint().IsOnGrid(elementSize) {

		for _, e := range snake.head.getAllBodyElements()[1:] {
			if e.getPoint().CollidesWith(*snake.head.getPoint()) {
				log.Fatal("Game Over")
			}
		}

		if snake.head.getPoint().CollidesWith(*food.p) {
			food = newFood(*game)
			grow += 10
		}

	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	for _, e := range snake.head.getAllBodyElements() {
		ops.GeoM.Reset()
		ops.GeoM.Translate(e.p.GetX(), e.p.GetY())
		screen.DrawImage(snakeImg, ops)
	}

	ops.GeoM.Reset()
	ops.GeoM.Translate(food.p.GetX(), food.p.GetY())
	screen.DrawImage(foodImg, ops)

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (sw, sh int) {
	return screenWidth, screenHeight
}

func main() {

	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Gosnake")
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
