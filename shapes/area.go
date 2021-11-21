package shapes

import "math"

type Rectangle struct {
	width  float64
	height float64
}

type Cirlce struct {
	radius float64
}

type Shape interface {
	Area() float64
}

func (rec Rectangle) Area() float64 {
	return rec.width * rec.height
}

func (circle Cirlce) Area() float64 {
	return circle.radius * math.Pi
}
