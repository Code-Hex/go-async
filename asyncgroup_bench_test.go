package async

import (
	"net/http"
	"net/http/httptest"
	"sync"
	"testing"
	"time"
)

func BenchmarkBackground(b *testing.B) {
	var ag Group
	for i := 1; i <= b.N; i++ {
		ag.Go(func() {
			Demo()
		})
	}
	ag.Wait()
}

func BenchmarkForeground(b *testing.B) {
	for i := 1; i <= b.N; i++ {
		Demo()
	}
}

func BenchmarkHTTPBackground(b *testing.B) {
	var ag Group
	ts := PrepareHTTP(func() {
		ag.Go(func() { Demo() })
	})
	defer ts.Close()
	var wg sync.WaitGroup
	b.ResetTimer()
	for i := 1; i <= b.N; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			http.Get(ts.URL)
		}()
	}
	wg.Wait()
	ag.Wait()
}

func BenchmarkHTTPForeground(b *testing.B) {
	ts := PrepareHTTP(func() {
		Demo()
	})
	defer ts.Close()
	var wg sync.WaitGroup
	b.ResetTimer()
	for i := 1; i <= b.N; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			http.Get(ts.URL)
		}()
	}
	wg.Wait()
}

func PrepareHTTP(f func()) *httptest.Server {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		f()
	}))
	return ts
}

func Demo() {
	time.Sleep(time.Millisecond * 100)
}
