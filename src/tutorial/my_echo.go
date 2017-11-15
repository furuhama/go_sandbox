// behaves like "echo" linux command
package tutorial

import (
	"fmt"
	"os"
)

func MyEcho() {
	var s, sep string
	for i := 1; i < len(os.Args); i++ {
		s += sep + os.Args[i]
		sep = " "
	}
	fmt.Println(s)
}
