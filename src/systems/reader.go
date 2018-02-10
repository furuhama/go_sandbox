// Package systems is for system layor program
package systems

import (
	"encoding/binary"
	"fmt"
	"io"
	"os"
	"strings"
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

// ReadOnlyHead reads only first some bytes of given data
func ReadOnlyHead() {
	reader := strings.NewReader("Example of io.SectionReader\n")
	sectionReader := io.NewSectionReader(reader, 14, 7)
	io.Copy(os.Stdout, sectionReader)
}

func dumpChunk(chunk io.Reader) {
	var length int32
	binary.Read(chunk, binary.BigEndian, &length)
	buffer := make([]byte, 4)
	chunk.Read(buffer)
	fmt.Printf("chunk '%v' (%d bytes)\n", string(buffer), length)
}

func readChunks(file *os.File) []io.Reader {
	// slice for setting chunk
	var chunks []io.Reader

	// skip first 8 bits
	file.Seek(8, 0)
	var offset int64 = 8

	for {
		var length int32
		err := binary.Read(file, binary.BigEndian, &length)
		if err == io.EOF {
			break
		}
		chunks = append(chunks, io.NewSectionReader(file, offset, int64(length)+12))
		// Move to head of next chunk
		offset, _ = file.Seek(int64(length+8), 1)
	}
	return chunks
}

// ReadPNGChunck reads PNG bytes and return it as a chunks
func ReadPNGChunck() {
	file, err := os.Open("Lenna.png")
	if err != nil {
		panic(err)
	}
	chunks := readChunks(file)
	for _, chunk := range chunks {
		dumpChunk(chunk)
	}
}
