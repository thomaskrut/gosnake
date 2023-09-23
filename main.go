package main

import (
	"image/color"
	"log"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	_ "github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/thomaskrut/gosnake/util"
)

var (
	game         *Game
	snake        *Snake
	food         *Food
	maze         *Maze
	grow         int
	foodCounter  int
	maxFood      int
	currentLevel int
	tileImg      *ebiten.Image
	foodImg      *ebiten.Image
	snakeImg     *ebiten.Image
	foodMeterImg *ebiten.Image
	ops          *ebiten.DrawImageOptions
	dirQueue     util.DirectionQueue
)

const (
	elementSize  = 5
	screenWidth  = 320
	screenHeight = 240
)

func init() {
	game = &Game{elementSize: elementSize, screenWidth: screenWidth, screenHeight: screenHeight}
	util.SetRandomSource(0)
	currentLevel = 1
	game.newLevel()
	ops = &ebiten.DrawImageOptions{}
	snakeImg = ebiten.NewImage(elementSize, elementSize)
	snakeImg.Fill(color.RGBA{0, 255, 0, 255})
	tileImg = ebiten.NewImage(elementSize, elementSize)
	tileImg.Fill(color.RGBA{120, 120, 120, 255})
	foodImg = ebiten.NewImage(elementSize, elementSize)
	foodImg.Fill(color.RGBA{255, 0, 0, 255})
	foodMeterImg = ebiten.NewImage(elementSize, elementSize)

}

func (g Game) newLevel() {
	grow = 40
	foodCounter = 0
	maxFood = 4
	snake = newSnake()
	maze = newMaze(game)
	maze.loadFromFile(strconv.Itoa(currentLevel))
	food = newFood(game, snake, maze)
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
		snake.tail().grow()
	}

	if len(dirQueue) > 0 && snake.head.point().IsOnGrid(elementSize) {
		snake.dir = dirQueue.Pop()
	}

	snake.head.move(snake.head.point().Move(snake.dir))

	if snake.head.isOnGrid() {

		for _, e := range snake.head.body()[1:] {
			if e.point().Overlaps(*snake.head.point()) {
				log.Fatal("Game Over")
			}
		}

		for _, p := range maze.Points() {
			if p.Overlaps(*snake.head.point()) {
				log.Fatal("Game Over")
			}
		}

		if snake.head.point().Overlaps(food.p) {
			food = newFood(game, snake, maze)
			foodCounter++
			if foodCounter == maxFood {
				currentLevel++
				game.newLevel()
			}
			grow += 20
		}

	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {

	//Snake
	for _, e := range snake.head.body() {
		ops.GeoM.Reset()
		ops.GeoM.Translate(e.p.GetX(), e.p.GetY())
		screen.DrawImage(snakeImg, ops)
	}

	//Maze
	for _, p := range maze.Points() {
		ops.GeoM.Reset()
		ops.GeoM.Translate(p.GetX(), p.GetY())
		screen.DrawImage(tileImg, ops)
	}

	//Food
	ops.GeoM.Reset()
	ops.GeoM.Translate(food.p.GetX(), food.p.GetY())
	screen.DrawImage(foodImg, ops)

	//Food meter
	for i := 0; i < foodCounter; i++ {
		ops.GeoM.Reset()
		ops.GeoM.Translate(float64(10+i*g.elementSize), 10)
		foodMeterImg.Fill(color.RGBA{255, 0, 0, 255})
		screen.DrawImage(foodMeterImg, ops)
	}
	for i := foodCounter; i < maxFood; i++ {
		ops.GeoM.Reset()
		ops.GeoM.Translate(float64(10+i*g.elementSize), 10)
		foodMeterImg.Fill(color.RGBA{100, 0, 0, 255})
		screen.DrawImage(foodMeterImg, ops)
	}

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
