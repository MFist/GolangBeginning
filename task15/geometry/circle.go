package geometry

import "math"

//Circle ...
type Circle struct {
	Radius float64 `json:"radius"`
}

// Area ...
func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

// Perim ...
func (c Circle) Perim() float64 {
	return 2 * math.Pi * c.Radius
}
