package main

import "fmt"

type QNode struct {
	northwest		*QNode
	northeast       *QNode
	southwest       *QNode
	southeast       *QNode
	startH, endH, startV, endV float32
	depth uint
}

func NewQNode(startH, endH, startV, endV float32, depth uint) *QNode {
	return &QNode{nil, nil, nil, nil, startH, endH, startV, endV, depth}
}

func (node *QNode) makeNorthWest() *QNode {
	midX, midY := node.getMidValues()
	node.northwest = NewQNode(node.startH, midX, node.startV, midY, node.depth+1)
	return node.northwest
}

func (node *QNode) makeNorthEast() *QNode {
	midX, midY := node.getMidValues()
	node.northeast = NewQNode(midX, node.endH, node.startV, midY, node.depth+1)
	return node.northeast
}

func (node *QNode) makeSouthWest() *QNode {
	midX, midY := node.getMidValues()
	node.southwest = NewQNode(node.startH, midX, midY, node.endV, node.depth+1)
	return node.southwest
}

func (node *QNode) makeSouthEast() *QNode {
	midX, midY := node.getMidValues()
	node.southeast = NewQNode(midX, node.endH, midY, node.endV, node.depth+1)
	return node.southeast
}

func (node *QNode) forEach(cb func(node *QNode), maxDepth uint) {
	if node.depth == maxDepth {
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
	return (node.startH+node.endH)*0.5, (node.startV+node.endV)*0.5
}

func (node *QNode) String() string {
	midX, midY := node.getMidValues()
	return fmt.Sprintf("[Node: (%f, %f)]", midY, midX)
}