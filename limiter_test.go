package goratelimiter

import (
	"math"
	"testing"
	"time"

	"github.com/matryer/is"
)

func TestLimiter_Consume(t *testing.T) {
	is := is.New(t)

	limiter := Limiter{}

	now := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC)

	ok := limiter.Consume(5, now)
	is.True(ok)
	equalFloat64(is, limiter.tokens, 4.0)

	ok = limiter.Consume(5, now.Add(10*time.Millisecond))
	is.True(ok)
	equalFloat64(is, limiter.tokens, 3.0+5.0*10.0*float64(time.Millisecond)/float64(time.Second))

	ok = limiter.Consume(5, now.Add(20*time.Millisecond))
	is.True(ok)
	equalFloat64(is, limiter.tokens, 2.0+5.0*20.0*float64(time.Millisecond)/float64(time.Second))

	ok = limiter.Consume(5, now.Add(30*time.Millisecond))
	is.True(ok)
	equalFloat64(is, limiter.tokens, 1.0+5.0*30.0*float64(time.Millisecond)/float64(time.Second))

	ok = limiter.Consume(5, now.Add(40*time.Millisecond))
	is.True(ok)
	equalFloat64(is, limiter.tokens, 0.0+5.0*40.0*float64(time.Millisecond)/float64(time.Second))

	ok = limiter.Consume(5, now.Add(40*time.Millisecond))
	is.True(!ok)

	ok = limiter.Consume(5, now.Add(2*time.Second))
	is.True(ok)
	equalFloat64(is, limiter.tokens, 4.0)
}

func TestLimiter_Consume_First(t *testing.T) {
	is := is.New(t)

	limiter := Limiter{}

	ok := limiter.Consume(100, time.Now())
	is.True(ok)
	equalFloat64(is, limiter.tokens, 99.0)
}

func equalFloat64(is *is.I, a float64, b float64) { //nolint:varnamelen // short names are okay here
	is.Helper()

	const epsilon = 0.0001

	// simple comparison should be okay here
	// https://floating-point-gui.de/errors/comparison/
	is.True(math.Abs(a-b) < epsilon)
}
