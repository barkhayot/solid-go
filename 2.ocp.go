package main

import (
	"math"
)

/*
Open/Closed Principle (OCP):
- Software entities (classes, structs, functions) should be open for extension but closed for modification.
- Instead of modifying existing code, we extend it.
*/

// Define a Shape interface that can be extended without modifying existing code.
type Shape interface {
	Area() float64
}

// Rectangle implements Shape interface.
type Rectangle struct {
	Width  float64
	Height float64
}

func (r Rectangle) Area() float64 {
	return r.Width * r.Height
}

// Circle implements Shape interface.
type Circle struct {
	Radius float64
}

func (c Circle) Area() float64 {
	return math.Pi * math.Pow(c.Radius, 2)
}

// Function to calculate total area of multiple shapes.
func TotalArea(shapes ...Shape) float64 {
	var total float64
	for _, shape := range shapes {
		total += shape.Area()
	}
	return total
}

// func main() {
// 	r := Rectangle{Width: 10, Height: 5}
// 	c := Circle{Radius: 7}

// 	fmt.Println("Rectangle Area:", r.Area())
// 	fmt.Println("Circle Area:", c.Area())

// 	shapes := []Shape{r, c}
// 	fmt.Println("Total Area:", TotalArea(shapes...))
// }
