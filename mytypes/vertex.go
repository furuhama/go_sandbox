// Package mytypes is practicing the OOP programming
package mytypes

import (
	"math"
)

// Vertex has two float parameters
type Vertex struct {
	X, Y float64
}

// RecieverAbs is Vertex type method
// this returns the length of Vertex
func (v Vertex) RecieverAbs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// NormalAbs is Vertex type method
// this returns the length of Vertex
func NormalAbs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
