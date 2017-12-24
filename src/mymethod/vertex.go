// Package mymethod is practicing the OOP programming
package mymethod

import (
	"math"
)

// Vertex has two float parameters
type Vertex struct {
	X, Y float64
}

// ReceiverAbs is Vertex type method
// this returns the length of Vertex
func (v Vertex) ReceiverAbs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

// NormalAbs is Vertex type method
// this returns the length of Vertex
func NormalAbs(v Vertex) float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}
