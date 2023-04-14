package main

import (
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

type Game[T Locatable] struct {
	root *QNode[T]
	// cursorTree *QNode[C]
	// datapoints []T
	maxDepth uint
}

func (g *Game[T]) Update() error {
	g.root.forEach(func (node *QNode[T]) bool {
		if node.marked {
			node.marked = false
			return node.marked
		}
		return true
	}, g.maxDepth)

	// width, height := ebiten.WindowSize()
	x, y := ebiten.CursorPosition()
	if x <= 0 || x > initialScreenWidth || y <= 0 || y > initialScreenHeight {
		return nil
	}
	g.root.markPathTo(float32(x), float32(y))
	return nil
}

func (g *Game[T]) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}


func (g *Game[T]) Draw(screen *ebiten.Image) {
	const strokeThickness = 1
	for _, spr := range g.root.datapoints {
		spr := any(spr).(*Sprite)
		vector.DrawFilledCircle(screen, spr.x, spr.y, strokeThickness, spr.clr, true)
		vector.StrokeCircle(screen, spr.x, spr.y, float32(spr.radius), 2, spr.clr, true)
	}

	width, height := ebiten.WindowSize()
	x, y := ebiten.CursorPosition()
	if x <= 0 || x > width || y <= 0 || y > height {
		return
	}

	g.root.forEach(func (node *QNode[T]) bool {
		if !node.marked {
			return true
		}
		midX, midY := node.getMidValues()
		vector.StrokeLine(screen, node.x0, midY, node.x1, midY, strokeThickness, color.White, false)
		vector.StrokeLine(screen, midX, node.y0, midX, node.y1, strokeThickness, color.White, false)
		return false
	}, g.maxDepth)
}