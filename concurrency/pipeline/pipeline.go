package main

import "fmt"

type Task struct {
	Id    int
	Index int
}

func build(in []int) <-chan Task {
	out := make(chan Task)

	go func() {
		for _, v := range in {
			out <- Task{Id: v}
		}
		close(out)
	}()

	return out
}

func makeIndex(id int) int {
	return id - 10
}

func fillIndex(in <-chan Task) <-chan Task {
	out := make(chan Task)

	go func() {
		for curTask := range in {
			curTask.Index = makeIndex(curTask.Id)
			out <- curTask
		}

		close(out)
	}()

	return out
}

func main() {
	in := build([]int{1, 2, 3})
	out := fillIndex(in)

	for task := range out {
		fmt.Println(task)
	}
}
