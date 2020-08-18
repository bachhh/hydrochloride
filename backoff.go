package hydrochloride

import (
	"fmt"
	"math"
	"math/rand"
	"sync/atomic"
	"time"
)

type Backoff interface {

	// Next shows the waiting time until next attempts
	Next() time.Duration

	Counter() uint64 // shows number of attempts thus far

	Reset() // reset backoff counter
}

const maxInt64 = float64(math.MaxInt64 - 512)

type expo struct {
	min     time.Duration
	max     time.Duration
	exp     float64 // exponent factor
	counter uint64
	jitter  float64
}

// NewExponentialBackoff return a new exponential backoff
// set jitter != 0 to enable random jittering
func NewExponentialBackoff(min, max time.Duration, factor float64, jitter float64) (b Backoff, err error) {
	if jitter >= 1 {
		return nil, fmt.Errorf("jitter value must not be >= 1")
	}
	return &expo{min: min, max: max, exp: factor, jitter: jitter}, nil
}

// Next shows the waiting time until next attempts
func (e *expo) Next() time.Duration {
	c := atomic.AddUint64(&e.counter, 1) - 1

	next := float64(e.min) * math.Pow(e.exp, float64(c))
	if next > float64(maxInt64) || next > float64(e.max) {
		next = float64(e.max)
	}

	if e.jitter > 0 {
	again: // random float implementation copied from  math/rand
		f := float64(rand.Int63()) / (1 << 63)
		if f > e.jitter {
			// if jitter < [0.0, 1.0) then complexity follows a Geometric distribution Pr(p = e.jitter)
			// worst case is infinite loop,
			// quick calculation shows for p = 0.000001; p99 = 25
			goto again
		}
	}
	return time.Duration(next)
}

func (e *expo) Counter() uint64 {
	return atomic.LoadUint64(&e.counter)
}

func (e *expo) Reset() {
	atomic.StoreUint64(&e.counter, 0)
}
