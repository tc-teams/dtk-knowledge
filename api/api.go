package api

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/tc-teams/fakefinder-crawler/api/server"
	"net/http"
	"os"
	"os/signal"
	"sync"
)

type API struct {
	Client     *server.Client
	Router     *mux.Router
	Routes     *Route
	Middleware *Middleware
	context    context.Context
	logging    *Logging
}

func (a *API) Serve() error {
	ctx, cancel := context.WithCancel(a.context)

	go func() {
		ch := make(chan os.Signal, 1)
		signal.Notify(ch, os.Interrupt)
		<-ch
		a.logging.Info("signal caught. shutting down...")
		cancel()
		a.Client.Shutdown(ctx)
	}()

	var wg sync.WaitGroup

	wg.Add(1)
	go func() {
		defer wg.Done()
		defer cancel()
		a.Client.Serve(ctx, a)
	}()

	wg.Wait()
	return nil
}
func (a *API) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var h http.Handler = a.Router
	h.ServeHTTP(w, r)
}

// NewContextApi returns a new instance API
func NewContextApi() (*API, error) {
	return &API{
		Client:     server.NewClient(),
		Middleware: newMiddlewareContext(),
		Router:     mux.NewRouter().StrictSlash(true),
		context:    context.Background(),
		logging:    NewLog(),
	}, nil
}
