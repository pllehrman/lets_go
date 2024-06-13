package shapes

import "math"

// A struct is a named collection of fields where one can store data. These structs become types which help to indicate
// to others the various concepts which are being tested/iterated on. Notice the function signatures becomes "rectangle Rectangle"

// Since both rectangle and circle have a method called Area() that returns a float64, it satisfies the Shape interface.
// You can think of interfaces as containing collections of structs
type Shape interface {
	Area() float64
}

type Rectangle struct {
	Width float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
} 

func Perimeter (rectangle Rectangle) float64 {
	return 2 * (rectangle.Width + rectangle.Height)
}

type Circle struct {
	Radius float64
}

func (c Circle) Area () float64 {
	return math.Pi * c.Radius * c.Radius
}


