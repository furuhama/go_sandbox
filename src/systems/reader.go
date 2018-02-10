// Package systems is for system layor program
package systems

import (
	"fmt"
	"io"
	"os"
)

// ReadStdio reads Std input & return it
func ReadStdio() {
	fmt.Println("return input")
	for {
		buffer := make([]byte, 5)
		size, err := os.Stdin.Read(buffer)
		if err == io.EOF {
			fmt.Println("EOF")
			break
		}
		fmt.Printf("size=%d input='%s'\n", size, string(buffer))
	}
}

// OpenFile opens file & copy data to Stdout
func OpenFile() {
	file, err := os.Open("hello.txt")
	if err != nil {
		panic(err)
	}
	defer file.Close()
	io.Copy(os.Stdout, file)
}
