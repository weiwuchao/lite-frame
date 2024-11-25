package workpool

import (
	"fmt"
	"sync"
)

type task struct {
	handler func(args ...interface{})
	args    []interface{}
}
type PoolState string

var (
	ActiveState   PoolState = "active"
	DisabledState PoolState = "disabled"
)

type WorkerPool struct {
	capacity      int
	state         PoolState
	runningWorker int
	taskQueue     chan task
	waitQueue     []task
	quit          chan struct{}
	sync.Mutex
}

func NewWorkerPool(capacity int) *WorkerPool {
	if capacity <= 0 {
		capacity = 100
	}
	p := &WorkerPool{capacity: capacity, runningWorker: 0, state: ActiveState, taskQueue: make(chan task, capacity), waitQueue: make([]task, 0)}
	p.start()
	return p
}

func (p *WorkerPool) start() {
	if p.state == DisabledState {
		return
	}
	for i := 0; i < p.capacity; i++ {
		go func() {
			for {
				select {
				case t := <-p.taskQueue:
					p.IncRunningWorker()
					fmt.Printf("t.agrs=%s\n", t.args)
					t.handler(t.args)
					p.DecRunningWorker()
				case <-p.quit:
					p.state = DisabledState
				}
			}
		}()
	}
}

func (p *WorkerPool) Stop() {
	p.quit <- struct{}{}
}

func (p *WorkerPool) Submit(workFunc func(args ...interface{}), args ...interface{}) {
	if p.state == DisabledState {
		return
	}
	t := task{handler: workFunc, args: args}
	p.taskQueue <- t
}

func (p *WorkerPool) RunningWorker() int {
	p.Lock()
	defer p.Unlock()
	return p.runningWorker
}

func (p *WorkerPool) Capacity() int {
	p.Lock()
	defer p.Unlock()
	return p.capacity
}

func (p *WorkerPool) IncRunningWorker() int {
	p.Lock()
	defer p.Unlock()
	p.runningWorker++
	return p.runningWorker
}

func (p *WorkerPool) DecRunningWorker() int {
	p.Lock()
	defer p.Unlock()
	p.runningWorker--
	return p.runningWorker
}
