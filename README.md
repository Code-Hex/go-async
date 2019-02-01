Go Async-Group
=====

Go Async Group can be more easily manage goroutine.

[![CircleCI](https://circleci.com/gh/Code-Hex/go-async.svg?style=svg)](https://circleci.com/gh/Code-Hex/go-async)
[![codecov](https://codecov.io/gh/Code-Hex/go-async/branch/master/graph/badge.svg)](https://codecov.io/gh/Code-Hex/go-async)
[![Go Report Card](https://goreportcard.com/badge/github.com/Code-Hex/go-async)](https://goreportcard.com/report/github.com/Code-Hex/go-async)
[![GoDoc](https://godoc.org/github.com/Code-Hex/go-async?status.svg)](https://godoc.org/github.com/Code-Hex/go-async)

# Synopsis

```go
func main() {
    var ag async.Group
	var sum, expected uint32
	for i := 0; i < times; i++ {
		expected++
		ag.Go(func() {
			atomic.AddUint32(&sum, 1)
		})
	}
	ag.Wait()
	fmt.Println(sum, expected)
}
```

If you want to use at your application, see [example](https://github.com/Code-Hex/go-async/blob/master/example_test.go).

# Installation

    go get -u github.com/Code-Hex/go-async

# Contribution

1. Fork [https://github.com/Code-Hex/go-async/fork](https://github.com/Code-Hex/go-async/fork)
2. Commit your changes
3. Create a new Pull Request

I'm waiting for a lot of PR.

# Author

[codehex](https://twitter.com/CodeHex)