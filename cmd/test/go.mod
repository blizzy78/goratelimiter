module test

go 1.22

toolchain go1.22.1

require (
	github.com/blizzy78/goratelimiter v0.1.0
	github.com/blizzy78/gotimeseries v0.2.0
)

replace github.com/blizzy78/goratelimiter => ../..
