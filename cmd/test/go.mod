module test

go 1.21

toolchain go1.21.5

require (
	github.com/blizzy78/goratelimiter v0.1.0
	github.com/blizzy78/gotimeseries v0.1.1
)

replace github.com/blizzy78/goratelimiter => ../..
