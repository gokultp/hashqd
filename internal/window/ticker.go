package window

import "time"

type ticker struct {
	*time.Ticker
	interval time.Duration
}

func NewTicker(interval time.Duration) *ticker {
	return &ticker{
		Ticker:   time.NewTicker(interval),
		interval: interval,
	}
}

func (t *ticker) Reset() {
	t.Ticker = time.NewTicker(t.interval)
}
