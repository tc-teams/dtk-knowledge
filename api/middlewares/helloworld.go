package middlewares

import (
	"github.com/tc-teams/fakefinder-crawler/api"
	"log"
	"net/http"
	"time"
)

//HelloWord is the first middlewares for api
func HelloWord() api.MiddlewareFunc {
	return func(h http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			log.Println("This request was sent at",time.Now())
			h.ServeHTTP(w, r)
		})
	}
}

