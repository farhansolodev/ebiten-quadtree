package main

import "fmt"

type QNode struct {
	one        *QNode
	two        *QNode
	three      *QNode
	four       *QNode
	startH, endH, startV, endV float32
	depth uint
}

func NewQNode(startH, endH, startV, endV float32, depth uint) *QNode {
	return &QNode{nil, nil, nil, nil, startH, endH, startV, endV, depth}
}

func (node *QNode) birthAt1() *QNode {
	midX, midY := node.getMidValues()
	node.one = NewQNode(node.startH, midX, node.startV, midY, node.depth+1)
	return node.one
}

func (node *QNode) birthAt2() *QNode {
	midX, midY := node.getMidValues()
	node.two = NewQNode(midX, node.endH, node.startV, midY, node.depth+1)
	return node.two
}

func (node *QNode) birthAt3() *QNode {
	midX, midY := node.getMidValues()
	node.three = NewQNode(node.startH, midX, midY, node.endV, node.depth+1)
	return node.three
}

func (node *QNode) birthAt4() *QNode {
	midX, midY := node.getMidValues()
	node.four = NewQNode(midX, node.endH, midY, node.endV, node.depth+1)
	return node.four
}

func (node *QNode) forEveryNode(cb func(node *QNode), maxDepth uint) {
	if node.depth == maxDepth {
		return
	}
	
	cb(node)
	if node.one != nil {
		node.one.forEveryNode(cb, maxDepth)
	} else if node.two != nil {
		node.two.forEveryNode(cb, maxDepth)
	} else if node.three != nil {
		node.three.forEveryNode(cb, maxDepth)
	} else if node.four != nil {
		node.four.forEveryNode(cb, maxDepth)
	}
}

func (node *QNode) getMidValues() (x, y float32) {
	return (node.startH+node.endH)*0.5, (node.startV+node.endV)*0.5
}

func (node *QNode) String() string {
	midX, midY := node.getMidValues()
	return fmt.Sprintf("[Node: (%f, %f)]", midY, midX)
}