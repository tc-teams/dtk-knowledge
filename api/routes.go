package api

type Route struct {
	Name  string        `json:"name"`
	Route *ContextRoute `json:"route"`
}

func (c *Route) RouteName(name string) {
	c.Name = name
}

func (c *Route) AddRoute(r *ContextRoute) {
	c.Route = r
}

//When a program runs concurrently, the parts of code which modify shared resources should not be accessed by multiple Goroutines at the same time
// InitModule initializes a module
func (a *API) InitRoute(r *Route) {
	r.Route.api = a
	r.Route.muxRoute = a.Router.Handle(r.Route.Path, r.Route).Methods(r.Route.Method)


}
