package geometry

// Rect ...
type Rect struct {
	Width  float64 `json:"width"`
	Height float64 `json:"height"`
}

// Area ...
func (r Rect) Area() float64 {
	return r.Width * r.Height
}

// Perim ...
func (r Rect) Perim() float64 {
	return 2*r.Width + 2*r.Height
}
