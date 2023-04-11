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
	// width, height := ebiten.WindowSize()
	// g.root.endH = float32(width)
	// g.root.endV = float32(height)

	x, y := ebiten.CursorPosition()
	if x < 0 || x > int(g.root.endH) || y < 0 || y > int(g.root.endV) { return nil }

	g.root.forEveryNode(func (node *QNode) {
		midX, midY := node.getMidValues()

		// 1st quadrant
		if x < int(midX) && y < int(midY) {
			if node.one == nil {
				node.birthAt1()
			}
			node.two = nil
			node.three = nil
			node.four = nil
			return
		}

		// 2nd quadrant
		if x > int(midX) && y < int(midY) {
			if node.two == nil {
				node.birthAt2()
			}
			node.one = nil
			node.three = nil
			node.four = nil
			return
		}

		// 3rd quadrant
		if x < int(midX) && y > int(midY) {
			if node.three == nil {
				node.birthAt3()
			}
			node.one = nil
			node.two = nil
			node.four = nil
			return
		}

		// if 4th quadrant
		if x > int(midX) && y > int(midY) {
			if node.four == nil {
				node.birthAt4()
			}
			node.one = nil
			node.two = nil
			node.three = nil
			return
		}

	}, g.maxDepth)

	g = &Game{ root: NewQNode(0, g.root.endH, 0, g.root.endV, 0), maxDepth: 0 }
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	const lineThickness = 1
	if g.root.one == nil && g.root.two == nil && g.root.three == nil && g.root.four == nil {
		return
	}
	g.root.forEveryNode(func (node *QNode) {
		midX, midY := node.getMidValues()
		vector.StrokeLine(screen, node.startH, midY, node.endH, midY, lineThickness, color.White, false)
		vector.StrokeLine(screen, midX, node.startV, midX, node.endV, lineThickness, color.White, false)
	}, g.maxDepth)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return outsideWidth, outsideHeight
}