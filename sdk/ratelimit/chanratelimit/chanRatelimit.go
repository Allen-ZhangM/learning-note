package chanratelimit

import (
	"time"
)

type ChanRateLimiter struct {
	limitChan <-chan time.Time
	rate      uint64
	allowance uint64
	max       uint64
	unit      uint64
}

func New(rate int, per time.Duration) *ChanRateLimiter {
	nano := uint64(per)
	if nano < 1 {
		nano = uint64(time.Second)
	}
	if rate < 1 {
		rate = 1
	}
	return &ChanRateLimiter{
		limitChan: time.Tick(time.Duration(int64(per) / int64(rate))),
		rate:      uint64(rate),
		allowance: uint64(rate) * nano,
		max:       uint64(rate) * nano,
		unit:      nano,
	}
}

func (crl *ChanRateLimiter) Limit() bool {
	select {
	case <-crl.limitChan:
		return true
	default:
		return false
	}
}
