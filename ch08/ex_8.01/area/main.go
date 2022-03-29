package main

import (
	"github.com/exvimmer/exercise8.01/shape"
)

func main() {
	s := shape.Square{
		Side: 9,
	}

	t := shape.Triangle{
		Base:   4,
		Height: 9,
	}

	r := shape.Rectangle{
		Width:  4,
		Height: 9,
	}

	shape.PrintShapeDetails(s, t, r)
}
