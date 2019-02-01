Go Async-Group
=====

Go Async Group can be more easily manage goroutine.

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

If you want to use at your application, see example.

# Installation

    go get -u github.com/Code-Hex/go-async

# Contribution

1. Fork [https://github.com/Code-Hex/go-async/fork](https://github.com/Code-Hex/go-async/fork)
2. Commit your changes
3. Create a new Pull Request

I'm waiting for a lot of PR.

# Author

[codehex](https://twitter.com/CodeHex)