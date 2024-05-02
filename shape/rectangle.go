package shape

type Rectangle struct {
	Witdh  float32
	Heigth float32
}

func (rectangle *Rectangle) Area() float32 {
	return rectangle.Heigth * rectangle.Witdh
}

func (rectangle *Rectangle) Perimeter() float32 {
	return 2 * (rectangle.Heigth + rectangle.Witdh)
}
