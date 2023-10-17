package main

import (
	"fmt"
	"strings"
)

type Locatable interface {
	getPosition() (x, y float32)
}

type QNode[T Locatable] struct {
	northwest, northeast, southwest, southeast *QNode[T]
	datapoints                                 []T
	marked                                     bool
	x0, x1, y0, y1                             float32
	depth                                      uint
}

func NewQNode[T Locatable](datapoints []T, x0, x1, y0, y1 float32, depth uint) *QNode[T] {
	return &QNode[T]{nil, nil, nil, nil, datapoints, false, x0, x1, y0, y1, depth}
}

func (node *QNode[T]) makeNorthWest(datapoints []T) *QNode[T] {
	midX, midY := node.getMidValues()
	node.northwest = NewQNode(datapoints, node.x0, midX, node.y0, midY, node.depth+1)
	return node.northwest
}

func (node *QNode[T]) makeNorthEast(datapoints []T) *QNode[T] {
	midX, midY := node.getMidValues()
	node.northeast = NewQNode(datapoints, midX, node.x1, node.y0, midY, node.depth+1)
	return node.northeast
}

func (node *QNode[T]) makeSouthWest(datapoints []T) *QNode[T] {
	midX, midY := node.getMidValues()
	node.southwest = NewQNode(datapoints, node.x0, midX, midY, node.y1, node.depth+1)
	return node.southwest
}

func (node *QNode[T]) makeSouthEast(datapoints []T) *QNode[T] {
	midX, midY := node.getMidValues()
	node.southeast = NewQNode(datapoints, midX, node.x1, midY, node.y1, node.depth+1)
	return node.southeast
}

func (node *QNode[T]) forEach(skip func(node *QNode[T]) bool, maxDepth uint) {
	if node.depth > maxDepth {
		return
	}
	if skip(node) {
		return
	}
	if node.northwest != nil {
		node.northwest.forEach(skip, maxDepth)
	}
	if node.northeast != nil {
		node.northeast.forEach(skip, maxDepth)
	}
	if node.southwest != nil {
		node.southwest.forEach(skip, maxDepth)
	}
	if node.southeast != nil {
		node.southeast.forEach(skip, maxDepth)
	}
}

func (node *QNode[T]) markPathTo(x, y float32) {
	node.marked = true
	midX, midY := node.getMidValues()
	switch {
	case x < midX && y < midY:
		if node.northwest == nil {
			return
		}
		node.northwest.markPathTo(x, y)
	case x > midX && y < midY:
		if node.northeast == nil {
			return
		}
		node.northeast.markPathTo(x, y)
	case x < midX && y > midY:
		if node.southwest == nil {
			return
		}
		node.southwest.markPathTo(x, y)
	case x > midX && y > midY:
		if node.southeast == nil {
			return
		}
		node.southeast.markPathTo(x, y)
	}
}

func (node *QNode[T]) generateTree(maxDepth uint) {
	node.forEach(func(node *QNode[T]) bool {
		nwDatapoints := make([]T, 0)
		neDatapoints := make([]T, 0)
		swDatapoints := make([]T, 0)
		seDatapoints := make([]T, 0)
		for _, v := range node.datapoints {
			x, y := v.getPosition()
			midX, midY := node.getMidValues()
			switch {
			case x < midX && y < midY:
				nwDatapoints = append(nwDatapoints, v)
			case x > midX && y < midY:
				neDatapoints = append(neDatapoints, v)
			case x < midX && y > midY:
				swDatapoints = append(swDatapoints, v)
			case x > midX && y > midY:
				seDatapoints = append(seDatapoints, v)
			}
		}
		noData := true
		if len(nwDatapoints) != 0 {
			noData = false
			node.makeNorthWest(nwDatapoints)
		}
		if len(neDatapoints) != 0 {
			noData = false
			node.makeNorthEast(neDatapoints)
		}
		if len(swDatapoints) != 0 {
			noData = false
			node.makeSouthWest(swDatapoints)
		}
		if len(seDatapoints) != 0 {
			noData = false
			node.makeSouthEast(seDatapoints)
		}
		if noData {
			return true
		}
		return false
	}, maxDepth)
}

func (node *QNode[T]) getMidValues() (x, y float32) {
	return (node.x0 + node.x1) * 0.5, (node.y0 + node.y1) * 0.5
}

func (node *QNode[T]) String() string {
	var sb strings.Builder
	node.forEach(func(node *QNode[T]) bool {
		midX, midY := node.getMidValues()
		sb.WriteString(fmt.Sprintf("%s[Node: (%f, %f)]\n", strings.Repeat("-> ", int(node.depth)), midX, midY))
		return false
	}, 10)
	return sb.String()
}
