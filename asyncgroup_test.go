package async

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"sync"
	"sync/atomic"
	"testing"
	"time"
)

const times = 1000

func TestRunTask(t *testing.T) {
	var ag Group
	var sum, expected uint32
	for i := 0; i < times; i++ {
		expected++
		ag.Go(func() {
			atomic.AddUint32(&sum, 1)
		})
	}
	ag.Wait()
	if sum != expected {
		t.Fatalf("got %d, but expected %d", sum, expected)
	}
}

func TestHTTPServe(t *testing.T) {
	var ag Group
	var sum uint32
	ts := httptest.NewServer(
		http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			// we shouldn't use the request context like `ctx := r.Context()`
			// please generate new context.
			// ag.Go(func() {
			// 	ctx := context.Background()
			// })
			ag.Go(func() {
				atomic.AddUint32(&sum, 1)
			})
		}),
	)

	expected := request(ts.URL)

	ag.Wait()
	if sum != expected {
		t.Fatalf("got %d, but expected %d", sum, expected)
	}
}

func request(url string) uint32 {
	var wg sync.WaitGroup
	var expected uint32
	for i := 0; i < times; i++ {
		expected++
		wg.Add(1)
		go func() {
			defer wg.Done()
			_, err := http.Get(url)
			if err != nil {
				panic(fmt.Sprintf("Get: %v", err))
			}
		}()
		if i%100 == 0 {
			// keep for resource
			time.Sleep(time.Millisecond * 200)
		}
	}
	wg.Wait()
	return expected
}
