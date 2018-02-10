// Package systems is for system layor program
package systems

import (
	"bufio"
	"bytes"
	"fmt"
	"os"
)

// CreateText is making txt file & write byte data in it
func CreateText() {
	file, err := os.Create("test.txt")
	if err != nil {
		panic(err)
	}
	file.Write([]byte("os.File examples"))
	file.Close()
}

// StdOut send Stdout to example texts
func StdOut() {
	os.Stdout.Write([]byte("os.Stdout example\n"))
}

// Buffer put byte data into buffer, and send it to Stdout
func Buffer() {
	var buffer bytes.Buffer
	buffer.Write([]byte("bytes.Buffer example\n"))
	fmt.Println(buffer.String())
}

// FlushBuf uses Buffer and flush its data
func FlushBuf() {
	buffer := bufio.NewWriter(os.Stdout)
	buffer.WriteString("bufio.Writer ")
	buffer.Flush()
	buffer.WriteString("example\n")
	buffer.Flush()
}
