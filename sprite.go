package main

import (
	"image/color"
	"math/rand"
)

type Sprite struct {
	clr color.Color
	x, y float32
	radius uint
}

//nolint:exhaustruct
func NewSprite(seed int64) *Sprite {
	spr := &Sprite{}
	spr.setRandomColor(seed)
	spr.setRandomRadius(seed)
	spr.setRandomInitialPosition(seed)
	return spr
}

func (spr *Sprite) setRandomColor(seed int64) {
	rand.Seed(seed)
	// so that the color is not too dark
	const min = 30
	const max = 256
	r := uint8(rand.Intn(max-min)+min)
	g := uint8(rand.Intn(max-min)+min)
	b := uint8(rand.Intn(max-min)+min)
	spr.clr = color.RGBA{r, g, b, 255}
}

func (spr *Sprite) setRandomRadius(seed int64) {
	rand.Seed(seed)
	const min = 10
	const max = 30
	spr.radius = uint(rand.Intn(max-min)+min)
}

func (spr *Sprite) setRandomInitialPosition(seed int64) {
	rand.Seed(seed)
	min := int(spr.radius)

	maxX := int(initialScreenWidth - spr.radius)
	spr.x = float32(rand.Intn(maxX-min)+min)

	maxY := int(initialScreenHeight - spr.radius)
	spr.y = float32(rand.Intn(maxY-min)+min)
}

func (spr *Sprite) getPosition() (x, y float32) {
	return spr.x, spr.y
}
