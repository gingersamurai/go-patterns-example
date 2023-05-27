package main

import "fmt"

func count(start, finish int) <-chan int {
	ch := make(chan int)
	go func() {
		defer close(ch)
		for i := start; i <= finish; i++ {
			ch <- i
		}
	}()

	return ch
}

func main() {
	for i := range count(1, 5) {
		fmt.Println(i)
	}
}
