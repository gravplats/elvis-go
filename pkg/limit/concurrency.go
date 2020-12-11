package limit

// Poor man's concurrency limiting.
type Concurrency struct {
	ch chan struct{}
}

func NewConcurrency(max int) Concurrency {
	ch := make(chan struct{}, max)
	return Concurrency{ch: ch}
}

func (c Concurrency) Release() {
	<-c.ch
}

func (c Concurrency) Take() {
	c.ch <- struct{}{}
}
