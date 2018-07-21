package main

import (
	"fmt"
	"github.com/furuhama/go_sandbox/mytypes"
)

func main() {
	v := mytypes.Vertex{X: 10, Y: 20}
	fmt.Println(v.RecieverAbs())
}
