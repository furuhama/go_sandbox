// Package systems is for system layor program
package systems

import (
	"os"
)

// FetchOSCreate fetch os.Create
func FetchOSCreate() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	file.Write([]byte("system call example"))
}
