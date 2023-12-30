package main //nolint:revive // no documentation needed here

import (
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/blizzy78/goratelimiter"
	"github.com/blizzy78/gotimeseries"
)

const (
	maxPerSec = 666

	clientMaxWait = 500 * time.Millisecond
)

const (
	seriesFineGranularity = 20
	seriesGranularity     = time.Second / seriesFineGranularity
	seriesBuckets         = 5 * seriesFineGranularity
)

func main() {
	mu := sync.Mutex{}
	limiter := goratelimiter.Limiter{}
	series := gotimeseries.New(seriesGranularity, seriesBuckets, time.Now())

	for i := 1; i < 10000; i++ {
		go runClient(&mu, &limiter, series)
	}

	moveUp := false
	dots := 0

	for {
		time.Sleep(500 * time.Millisecond)

		mu.Lock()

		total := series.Total()

		mu.Unlock()

		if moveUp {
			fmt.Print("\033[A")
		}

		fmt.Printf("\033[K%.1f ", float64(total)/(seriesGranularity*seriesBuckets).Seconds())

		for i := 0; i <= dots; i++ {
			fmt.Print(".")
		}

		dots++
		dots %= 3

		fmt.Println()

		moveUp = true
	}
}

func runClient(mu *sync.Mutex, limiter *goratelimiter.Limiter, series *gotimeseries.TimeSeries) {
	doWork := func() {
		now := time.Now()

		mu.Lock()
		defer mu.Unlock()

		if !limiter.Consume(maxPerSec, now) {
			return
		}

		series.Update(now)
		series.Increase()
	}

	for {
		time.Sleep(time.Duration(rand.Float64() * float64(clientMaxWait))) //nolint:gosec // we don't need a secure random number here

		doWork()
	}
}
