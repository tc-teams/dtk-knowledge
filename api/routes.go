package api

import (
	"github.com/sirupsen/logrus"
)

type Route struct {
	Name  string          `json:"name"`
	Route []*ContextRoute `json:"route"`
}

func (c *Route) RouteName(name string) {
	c.Name = name
}

func (c *Route) AddRoute(r ...*ContextRoute) {
	for _, route := range r {
		c.Route = append(c.Route, route)
	}
}

//When a program runs concurrently, the parts of code which modify shared resources should not be accessed by multiple Goroutines at the same time
// InitModule initializes a module
func (a *API) InitRoute(r ...*Route) {
	for _, rt := range r[len(r)-1].Route {
		rt.api = a
		rt.mux = a.Router.Handle(rt.Path, rt).Methods(rt.Method)
		logrus.WithFields(logrus.Fields{
			"Method": rt.Method,
		}).Info("Create a new rote: ", rt.Path)
	}

}
