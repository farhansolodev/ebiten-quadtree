package main

import "fmt"

type QNode struct {
	northwest		*QNode
	northeast       *QNode
	southwest       *QNode
	southeast       *QNode
	x0, x1, y0, y1 float32
	depth uint
}

func NewQNode(x0, x1, y0, y1 float32, depth uint) *QNode {
	return &QNode{nil, nil, nil, nil, x0, x1, y0, y1, depth}
}

func (node *QNode) makeNorthWest() *QNode {
	midX, midY := node.getMidValues()
	node.northwest = NewQNode(node.x0, midX, node.y0, midY, node.depth+1)
	return node.northwest
}

func (node *QNode) makeNorthEast() *QNode {
	midX, midY := node.getMidValues()
	node.northeast = NewQNode(midX, node.x1, node.y0, midY, node.depth+1)
	return node.northeast
}

func (node *QNode) makeSouthWest() *QNode {
	midX, midY := node.getMidValues()
	node.southwest = NewQNode(node.x0, midX, midY, node.y1, node.depth+1)
	return node.southwest
}

func (node *QNode) makeSouthEast() *QNode {
	midX, midY := node.getMidValues()
	node.southeast = NewQNode(midX, node.x1, midY, node.y1, node.depth+1)
	return node.southeast
}

func (node *QNode) forEach(cb func(node *QNode), maxDepth uint) {
	if node.depth > maxDepth {
		return
	}
	
	cb(node)
	if node.northwest != nil {
		node.northwest.forEach(cb, maxDepth)
	} else if node.northeast != nil {
		node.northeast.forEach(cb, maxDepth)
	} else if node.southwest != nil {
		node.southwest.forEach(cb, maxDepth)
	} else if node.southeast != nil {
		node.southeast.forEach(cb, maxDepth)
	}
}

func (node *QNode) collapse(x, y float32, maxDepth uint) {
	node.forEach(func (node *QNode) {
		midX, midY := node.getMidValues()

		// 1st quadrant
		if x < midX && y < midY {
			if node.northwest == nil {
				node.makeNorthWest()
			}
			node.northeast = nil
			node.southwest = nil
			node.southeast = nil
			return
		}

		// 2nd quadrant
		if x > midX && y < midY {
			if node.northeast == nil {
				node.makeNorthEast()
			}
			node.northwest = nil
			node.southwest = nil
			node.southeast = nil
			return
		}

		// 3rd quadrant
		if x < midX && y > midY {
			if node.southwest == nil {
				node.makeSouthWest()
			}
			node.northwest = nil
			node.northeast = nil
			node.southeast = nil
			return
		}

		// if 4th quadrant
		if x > midX && y > midY {
			if node.southeast == nil {
				node.makeSouthEast()
			}
			node.northwest = nil
			node.northeast = nil
			node.southwest = nil
			return
		}

	}, maxDepth)

}

func (node *QNode) getMidValues() (x, y float32) {
	return (node.x0+node.x1)*0.5, (node.y0+node.y1)*0.5
}

func (node *QNode) String() string {
	midX, midY := node.getMidValues()
	return fmt.Sprintf("[Node: (%f, %f)]", midY, midX)
}