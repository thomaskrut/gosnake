package util



type Point struct {
	x, y int
}

func (p Point) GetX() float64 {
	return float64(p.x)
}

func (p Point) GetY() float64 {
	return float64(p.y)
}

func NewPoint(x, y int) Point {
	return Point{x: x, y: y}
}

func (p Point) IsOnGrid(elementSize int) bool {
	return p.x%elementSize == 0 && p.y%elementSize == 0
}

func (p Point) CollidesWith(other Point) bool {
	return p.x == other.x && p.y == other.y
}

func (p Point) Move(dir Direction) Point {
	x := p.x + dir.varX
	y := p.y + dir.varY
	return Point{x: x, y: y}
}

func GetRandomPoint(width, height, elementSize int) Point {
	x := RandomNumber(width)
	y := RandomNumber(height)
	x = x - (x % elementSize)
	y = y - (y % elementSize)

	return NewPoint(x, y)
}

func GetRandomEmptyPoint(width, height, elementSize int, points [][]Point) Point {
	p := GetRandomPoint(width, height, elementSize)
	for _, pointList := range points {
		for _, point := range pointList {
			if p.CollidesWith(point) {
				return GetRandomEmptyPoint(width, height, elementSize, points)
			}
		}
	}
	return p
	
}

func (p Point) GetAdjecentPoint(dir Direction, elementSize int) Point {
	return NewPoint(p.x+dir.varX*elementSize, p.y+dir.varY*elementSize)
}
