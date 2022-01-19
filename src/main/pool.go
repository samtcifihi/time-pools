package main

// Pool holds a time pool
type Pool struct {
	name     string
	ztime    ztime
	running  bool
	overflow bool
}

// NewPool is the Constructor for pool
func NewPool(name string, time string) *Pool {
	p := new(Pool)

	p.name = name
	p.ztime.Set(time)
	p.running = false
	p.overflow = false

	return p
}

// SetRunning sets "running"
func (p *Pool) SetRunning(val bool) {
	p.running = val
}
