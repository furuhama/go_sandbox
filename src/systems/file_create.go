// Package systems is for system layor program
package systems

import (
	"os"
)

// Create is making txt file & write byte data in it
func Create() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	file.Write([]byte("os.File examples"))
	file.Close()
}
