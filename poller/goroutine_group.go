package poller

import "sync"

type goroutineGroup struct {
	waitGroup sync.WaitGroup
}

func newRoutineGroup() *goroutineGroup {
	return new(goroutineGroup)
}

func (g *goroutineGroup) Run(fn func()) {
	g.waitGroup.Add(1)

	go func() {
		defer g.waitGroup.Done()
		fn()
	}()
}

func (g *goroutineGroup) Wait() {
	g.waitGroup.Wait()
}
