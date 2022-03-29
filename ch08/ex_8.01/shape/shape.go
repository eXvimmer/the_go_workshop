package shape

import "fmt"

type Shape interface {
	Area() float64
	Name() string
}

type Triangle struct {
	Base   float64
	Height float64
}

func (t Triangle) Name() string {
	return "triangle"
}

func (t Triangle) Area() float64 {
	return (t.Base * t.Height) / 2
}

type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Height * r.Width
}

func (r Rectangle) Name() string {
	return "rectangle"
}

type Square struct {
	Side float64
}

func (s Square) Area() float64 {
	return s.Side * s.Side
}

func (s Square) Name() string {
	return "square"
}

func PrintShapeDetails(shapes ...Shape) {
	for _, v := range shapes {
		fmt.Printf("The area of %s is %.2f\n", v.Name(), v.Area())
	}
}
