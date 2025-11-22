package point

import "math"

type Point struct {
	x float64
	y float64
}

func NewPoint(x, y float64) Point {
	return Point{x: x, y: y}
}

func (p Point) Distance(p2 Point) float64 {
	dx := p.x - p2.x
	dy := p.y - p2.y
	return math.Sqrt(dx*dx + dy*dy)
}
