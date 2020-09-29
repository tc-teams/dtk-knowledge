package api

import (
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"net/http"
)

type Handler func(w http.ResponseWriter, r *http.Request, logging *Logging) *BaseError

type ContextRoute struct {
	api     *API       //router
	Method  string     `json:"method"` //tipo de método http
	Path    string     `json:"path"`   //o caminho até esse método
	mux     *mux.Route //Uma rota ela é definida pela seguinte estrutura(rota simples)
	Handler Handler    `json:"-"` //metodo consumido

}

//Almost any object can be a handlers,
//so long as it satisfies the http.Handler interface.
//In lay terms, that simply means it must have a
//ServeHTTP method with the following signature:
func (c ContextRoute) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	c.WContentType(w)
	log := NewLog()
	log.Config()

	err := c.api.Middleware.ChainMiddleware(c.Handler)(w, r, log)
	if err != nil {
		c.api.logging.WithFields(logrus.Fields{
			"Error": err.Error,
			"Code":  err.Code,
		}).Error("application error", err.Message)
		http.Error(w, err.Message, err.Code)
	}

}
func (c *ContextRoute) WContentType(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}
