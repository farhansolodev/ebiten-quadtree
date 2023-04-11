package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game struct {
	root *QNode
	maxDepth uint
}

func (g *Game) Update() error {
	width, height := ebiten.WindowSize()
	g.root.endH = float32(width)
	g.root.endV = float32(height)

	x, y := ebiten.CursorPosition()
	if x < 0 || x > int(g.root.endH) || y < 0 || y > int(g.root.endV) { return nil }

	g.root.collapse(float32(x), float32(y), g.maxDepth)
	
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	const lineThickness = 1
	if g.root.northwest == nil && g.root.northeast == nil && g.root.southwest == nil && g.root.southeast == nil {
		return
	}
	g.root.forEach(func (node *QNode) {
		midX, midY := node.getMidValues()
		vector.StrokeLine(screen, node.startH, midY, node.endH, midY, lineThickness, color.White, false)
		vector.StrokeLine(screen, midX, node.startV, midX, node.endV, lineThickness, color.White, false)
	}, g.maxDepth)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}