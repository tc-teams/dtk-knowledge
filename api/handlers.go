package api

import (
	"github.com/gorilla/mux"
	"log"
	"net/http"
)


type Handler func(w http.ResponseWriter, r *http.Request) error

type ContextRoute struct {
	api      *API   //router
	Method   string `json:"method"` //tipo de método http
	Path     string `json:"path"`   //o caminho até esse método
	mux     *mux.Route             //Uma rota ela é definida pela seguinte estrutura(rota simples)
	Handler  Handler `json:"-"` //metodo consumido

}

//Almost any object can be a handlers,
//so long as it satisfies the http.Handler interface.
//In lay terms, that simply means it must have a
//ServeHTTP method with the following signature:
func (c ContextRoute) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.ServeContentType(w)

	err := c.api.Middleware.ChainMiddleware(c.Handler)(w,r)
	if err != nil {
		log.Print("ola mundo")
	}
	log.Print(c.Path)
	log.Print(c.Method)


}
func (c *ContextRoute) ServeContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json ;charset=UTF-8")
}



