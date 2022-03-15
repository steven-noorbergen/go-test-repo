package shapes

type Rectangle struct {
	Width  float64
	Height float64
}

func Area(r Rectangle) float64 {
	return r.Width * r.Height
}

func Perimeter(r Rectangle) float64 {
	return 2 * (r.Width + r.Height)
}
