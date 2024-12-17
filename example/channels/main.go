package main

import "fmt"

// https://gobyexample.com/channels
func main() {

	messages := make(chan string)

	go func() { messages <- "ping" }()

	msg := <-messages
	fmt.Println(msg)
}
