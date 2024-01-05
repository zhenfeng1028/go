package main

type Golimit struct {
	n int
	c chan struct{}
}

// initialization Golimit struct
func New(n int) *Golimit {
	return &Golimit{
		n: n,
		c: make(chan struct{}, n),
	}
}

// Run f in a new goroutine but with limit.
func (g *Golimit) Run(f func()) {
	g.c <- struct{}{}
	go func() {
		f()
		<-g.c
	}()
}

func (g *Golimit) Stop() {
	close(g.c)
}
