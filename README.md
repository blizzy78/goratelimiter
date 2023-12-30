[![GoDoc](https://pkg.go.dev/badge/github.com/blizzy78/goratelimiter)](https://pkg.go.dev/github.com/blizzy78/goratelimiter)


goratelimiter
=============

A Go package that provides a simple rate limiter that uses a token bucket algorithm.

```go
import "github.com/blizzy78/goratelimiter"
```


Code example
------------

```go
limiter := &goratelimiter.Limiter{}

// ... later ...
if !limiter.Consume(100, time.Now()) {
	// fail
	return
}

doWork()
```


License
-------

This package is licensed under the MIT license.
