package main

import (
	"flag"
	"log"
	"os"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

const (
	initialScreenWidth  = 1366
	initialScreenHeight = 768
)

func main() {
	maxDepth := flag.Uint("depth", 7, "set your own maximum quadtree depth")
	spriteCount := flag.Uint("sprites", 250, "how many sprites do you want?")
	// resizable := flag.Bool("resize", false, "should window be resizable? [true|false] (still experimental) (default false)")
	flag.Parse()

	ebiten.SetWindowSize(initialScreenWidth, initialScreenHeight)
	// if *resizable { ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled) }
	ebiten.SetWindowTitle(os.Args[0])

	sprites := make([]*Sprite, 0, *spriteCount)
	var i uint
	for ; i < *spriteCount; i++ {
		sprites = append(sprites, NewSprite(int64(i)+time.Now().UnixNano()))
	}

	root := NewQNode(sprites, 0, initialScreenWidth, 0, initialScreenHeight, 0)
	root.generateTree(*maxDepth)

	game := &Game{root, *maxDepth}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
