package main

type Food struct {
	p *Point
}

func newFood() *Food {
	newFood := Food{p: getRandomPoint()}
	return &newFood
}