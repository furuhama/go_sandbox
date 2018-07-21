package main

import (
	"fmt"
	"github.com/furuhama/go_sandbox/mytypes"
)

func main() {
	v := mytypes.Vertex{}
	v.X = 10
	v.Y = 20
	fmt.Println(v.RecieverAbs())
}
