package main

type Semaphore struct {
	sem chan struct{}
}

func (s *Semaphore) Acquire() {
	s.sem <- struct{}{}
}

func (s *Semaphore) Release() {
	_ = <-s.sem
}

func NewSemaphore(capacity int) *Semaphore {
	sem := make(chan struct{}, capacity)
	return &Semaphore{sem: sem}
}
