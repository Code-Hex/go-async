package async

import (
	"sync"
)

// A Group is a collection of goroutines working on async subtasks that are part of
// the same overall task.
type Group struct {
	wg sync.WaitGroup
}

// Go calls the given function in a new goroutine.
func (g *Group) Go(f func()) {
	g.wg.Add(1)
	go func() {
		defer g.wg.Done()
		f()
	}()
}

// Wait blocks until all function calls working on background from the Go method have returned.
func (g *Group) Wait() {
	g.wg.Wait()
}
