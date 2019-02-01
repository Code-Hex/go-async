package async

import (
	"context"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/Code-Hex/go-async"
)

func ExampleMain() {
	st := NewApp(8080)
	st.HandleBackground("/background")
	st.HandleForeground("/foreground")

	st.Serve()

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()
	if err := st.Shutdown(ctx); err != nil {
		log.Fatalf("shutdown err: %v", err)
	}
	st.bg.Wait()
}

type YourApp struct {
	srv *http.Server
	mux *http.ServeMux
	bg  async.Group
}

func NewApp(port int) *YourApp {
	mux := http.NewServeMux()
	return &YourApp{
		mux: mux,
		srv: &http.Server{
			Addr:    fmt.Sprintf(":%d", port),
			Handler: mux,
		},
	}
}

func (y *YourApp) Serve() {
	go func() {
		if err := y.srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("serve err: %v", err)
		}
	}()

	sig := make(chan os.Signal, 1)
	signal.Notify(sig, os.Interrupt, syscall.SIGTERM)
	<-sig
}

func (y *YourApp) Shutdown(ctx context.Context) error {
	return y.srv.Shutdown(ctx)
}

func (y *YourApp) HandleBackground(path string) {
	y.mux.Handle(path,
		http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			y.bg.Go(func() {
				Demo()
			})
		}),
	)
}

func (y *YourApp) HandleForeground(path string) {
	y.mux.Handle(path,
		http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			Demo()
		}),
	)
}

func Demo() {
	time.Sleep(time.Millisecond * 100)
}
