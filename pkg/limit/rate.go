package limit

import "time"

// Poor man's rate limiting.
type Rate struct {
	ch chan<- struct{}
}

func NewRate(perSecond int) Rate {
	ch := make(chan struct{}, 0)
	go func() {
		ticker := time.NewTicker(1000 / time.Duration(perSecond) * time.Millisecond)
		defer ticker.Stop()

		for range ticker.C {
			if _, ok := <-ch; !ok {
				return
			}
		}
	}()

	return Rate{ch: ch}
}

func (r Rate) Stop() {
	close(r.ch)
}

func (r Rate) Take() {
	r.ch <- struct{}{}
}
