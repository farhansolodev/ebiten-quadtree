package main

import (
	"fmt"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/basicfont"
)

type Game struct {
	root *QNode[*Sprite]
	// cursorTree *QNode[C]
	// datapoints []T
	maxDepth uint
}

func (g *Game) Update() error {
	g.root.forEach(func(node *QNode[*Sprite]) bool {
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

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight
}

func (g *Game) Draw(screen *ebiten.Image) {
	const strokeThickness = 1
	for _, spr := range g.root.datapoints {
		// spr := any(spr).(*Sprite)
		vector.DrawFilledCircle(screen, spr.x, spr.y, strokeThickness, spr.clr, true)
		vector.StrokeCircle(screen, spr.x, spr.y, float32(spr.radius), 2, spr.clr, true)
	}

	width, height := ebiten.WindowSize()
	x, y := ebiten.CursorPosition()
	if x > 0 && x <= width && y > 0 && y <= height {
		g.root.forEach(func(node *QNode[*Sprite]) bool {
			if !node.marked {
				return true
			}
			midX, midY := node.getMidValues()
			vector.StrokeLine(screen, node.x0, midY, node.x1, midY, strokeThickness, color.White, false)
			vector.StrokeLine(screen, midX, node.y0, midX, node.y1, strokeThickness, color.White, false)

			return false
		}, g.maxDepth)
	}

	drawDebug(screen,
		fmt.Sprintf("TPS: %.2f", ebiten.CurrentTPS()),
		fmt.Sprintf("FPS: %.2f", ebiten.CurrentFPS()),
		fmt.Sprintf("Sprites: %d", len(g.root.datapoints)),
	)
}

func drawDebug(screen *ebiten.Image, msgs ...string) {
	vector.DrawFilledRect(screen, 0, 0, initialScreenWidth*.1, initialScreenHeight*.035*float32(len(msgs)), color.RGBA{0x7f, 0x00, 0x7f, 0x7f}, true)
	font := text.FaceWithLineHeight(basicfont.Face7x13, 0)
	for i, msg := range msgs {
		text.Draw(screen, msg, font, 10, (i+1)*20, color.White)
	}
}
