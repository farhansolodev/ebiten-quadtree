package main

type Cursor struct {
	x, y float32
}

func (cur *Cursor) getPosition() (x, y float32) {
	return cur.x, cur.y
}
