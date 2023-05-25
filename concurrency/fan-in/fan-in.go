package main

import (
	"fmt"
	"sync"
)

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

func mergeTasks(in ...<-chan Task) <-chan Task {
	wg := sync.WaitGroup{}
	out := make(chan Task)

	wg.Add(len(in))
	for _, v := range in {
		v := v
		go func() {
			for curTask := range v {
				out <- curTask
			}
			wg.Done()
		}()
	}

	go func() {
		wg.Wait()
		close(out)
	}()

	return out
}

func main() {
	in := build([]int{1, 2, 3})
	out1 := fillIndex(in)
	out2 := fillIndex(in)

	for task := range mergeTasks(out1, out2) {
		fmt.Println(task.Id, task.Index)
	}
}
