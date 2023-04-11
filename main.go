package main

import (
	"flag"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 1280
	screenHeight = 720
)

func main() {
	maxDepth := flag.Uint("depth", 8, "set your own maximum quadtree depth")
	// resizable := flag.Bool("resize", false, "should window be resizable? [true|false] (still experimental) (default false)")
	flag.Parse()

	ebiten.SetWindowSize(screenWidth, screenHeight)
	// if *resizable { ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled) }
	ebiten.SetWindowTitle("ebiten-quadtree")
	
	game := Game{ root: NewQNode(0, screenWidth, 0, screenHeight, 0), maxDepth: *maxDepth }

	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}