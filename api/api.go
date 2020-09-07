package api

import (
	"context"
	"github.com/gorilla/mux"
	"github.com/tc-teams/fakefinder-crawler/api/server"
	elastic2 "github.com/tc-teams/fakefinder-crawler/elastic/es"
	"gopkg.in/go-playground/validator.v8"
	"net/http"
	"os"
	"os/signal"
	"sync"
)

var esurl = "http://localhost:8080"

type API struct {
	Client     *server.Client
	Router     *mux.Router
	Routes     *Route
	Middleware *Middleware
	Validator  *validator.Validate
	context    context.Context
	elastic    *elastic2.Elastic
	logging   *Logging
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
	//es, err := es.NewInstanceElastic(esurl)
	//if err != nil {
	//	return nil, err
	//
	//}
	return &API{
		Client:     server.NewClient(),
		Middleware: newMiddlewareContext(),
		Router:     mux.NewRouter().StrictSlash(true),
		Validator:  NewValidate().Validate,
		context:    context.Background(),
		elastic:    nil,
		logging: NewLog(),
	}, nil
}
