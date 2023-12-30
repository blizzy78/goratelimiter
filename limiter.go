package goratelimiter

import "time"

// Limiter implements a rate limiter that uses a token bucket algorithm:
// Each second, the bucket is filled with a number of tokens. For each event, e.g. a request,
// a token is consumed from the bucket. If the bucket is empty, the next event should be delayed,
// rejected, dropped, or whatever is appropriate for the use case.
type Limiter struct {
	updateTime time.Time
	tokens     float64
}

// Consume consumes a token from the bucket. It returns true if the token was consumed successfully,
// false otherwise.
//
// Consume also refills the limiter's bucket with new tokens, up to maxPerSec, which is the maximum
// number of tokens that can be consumed per second.
// If this is the first call to Consume, the bucket is filled with maxPerSec tokens.
//
// Concurrent calls to Consume must be synchronized by the caller.
func (l *Limiter) Consume(maxPerSec float64, now time.Time) bool {
	l.update(maxPerSec, now)

	if l.tokens < 1 {
		return false
	}

	l.tokens--

	return true
}

func (l *Limiter) update(maxPerSec float64, now time.Time) {
	elapsed := now.Sub(l.updateTime)
	if elapsed >= time.Second {
		l.updateTime = now
		l.tokens = maxPerSec

		return
	}

	l.updateTime = now
	l.tokens = min(l.tokens+maxPerSec*elapsed.Seconds(), maxPerSec)
}
