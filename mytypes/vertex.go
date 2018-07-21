// Package mytypes is practicing the OOP programming
package mytypes

import (
	"fmt"
	"math"
)

// Vertex has two float parameters
type Vertex struct {
	X, Y float64
}

// RecieverAbs is Vertex type method
// this returns the length of Vertex
func (v *Vertex) RecieverAbs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// NormalAbs is Vertex type method
// this returns the length of Vertex
func NormalAbs(v *Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// ReverseVertex converts struct Vertex's parameters
func (v *Vertex) ReverseVertex() {
	v.X, v.Y = v.Y, v.X
}

// PrintParameters puts struct Vertex's parameters to STDOUT
func (v *Vertex) PrintParameters() {
	fmt.Printf("X: %f, Y: %f\n", v.X, v.Y)
}
