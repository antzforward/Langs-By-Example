package main

import "fmt"

/**
This example also showed that it’s possible to close a non-empty channel but still have the remaining values be received.
*/
func main() {

	queue := make(chan string, 2)
	queue <- "one"
	queue <- "two"
	close(queue) //chan closed

	for elem := range queue {
		fmt.Println(elem)
	}
}

/**go run .\channel-range-over.go
one
two*/
