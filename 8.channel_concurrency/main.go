package main

import "fmt"

func main() {

	messages := make(chan string)

	go func() { messages <- "this statement went via a channel" }()

	received := <-messages

	fmt.Println(received)
}
