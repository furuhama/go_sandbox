package main

import (
	"github.com/furuhama/go_sandbox/mytypes"
)

func main() {
	v := mytypes.Vertex{X: 10, Y: 20}
	v.PrintParameters()

	v.ReverseVertex()
	v.PrintParameters()
}
