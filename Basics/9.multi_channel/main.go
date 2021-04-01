package main

import "fmt"

func main() {

	messages := make(chan string, 2)

	messages <- "Message 1"
	messages <- "Message 2"

	fmt.Println(<-messages)
	fmt.Println(<-messages)
}
