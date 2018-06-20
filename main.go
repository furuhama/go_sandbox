package main

import (
	"fmt"
	"log"
)

func worker(msg string) <-chan string {
	receiver := make(chan string)
	for i := 0; i < 3; i++ {
		// go routine で無名関数を呼び出している
		go func(i int) {
			// ここで msg に文字列を入れて、そのあと receiver chan から返している
			msg := fmt.Sprintf("%d %s done", i, msg)
			receiver <- msg
		}(i)
	}
	return receiver
}

func main() {
	receiver := worker("job")
	for i := 0; i < 3; i++ {
		log.Println(<-receiver)
	}
}
