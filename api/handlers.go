package api

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)

//Uma rota ela é definida pela seguinte estrutura(rota simples)
type ContextRoute struct {
	api      *API   //router
	Method   string `json:"method"` //tipo de método http
	Path     string `json:"path"`   //o caminho até esse método
	muxRoute *mux.Route
	Handler  http.HandlerFunc `json:"handler,omitempty"` //metodo consumido
}

//Almost any object can be a handlers,
//so long as it satisfies the http.Handler interface.
//In lay terms, that simply means it must have a
//ServeHTTP method with the following signature:
func (c ContextRoute) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	err := c.api.Middleware.ChainMiddleware(c.Handler)
	if err != nil {
		log.Print("ola mundo")
	}
	log.Print(c.Path)
	log.Print(c.Method)

}
