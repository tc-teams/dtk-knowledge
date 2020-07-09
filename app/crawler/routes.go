package crawler

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/sirupsen/logrus"
	"github.com/tc-teams/fakefinder-crawler/api"
	h "github.com/tc-teams/fakefinder-crawler/app/hack"
	"github.com/tc-teams/fakefinder-crawler/tracker"
	"net/http"
	"os"
)

const routerName = "crawler"

//Init return a new route instance
func Init() *api.Route {
	r := &api.Route{}
	r.RouteName(routerName)
	r.AddRoute(&api.ContextRoute{
		Method:  http.MethodGet,
		Path:    "/covid",
		Handler: NewsRelatedToCovid,
	})
	fmt.Printf("adicionando a rota %v\n", r)

	return r
}

func NewsRelatedToCovid(w http.ResponseWriter, r *http.Request) error {
	hack := h.NewHack()
	c := Covid{}

	if err := hack.ParseBody(c, r); err != nil {
		http.Error(w, err.Error(), http.StatusBadRequest)
		return err
	}
	err := HandlerCovid(c)
	if err != nil {
		http.Error(w, err.Error(), http.StatusNotFound)
		return err
	}
	w.WriteHeader(http.StatusOK)
	return nil

}

func HandlerCovid(cd Covid) error {

	c := tracker.NewColly(colly.NewCollector(
		colly.AllowedDomains(tracker.Folha, tracker.G1, tracker.Uol),
		colly.MaxDepth(3),
		colly.Async(true),
	),
		&logrus.Logger{
			Out:       os.Stdout,
			Formatter: &logrus.JSONFormatter{},
		}, nil, "",
	)

	logrus.WithFields(logrus.Fields{"Text": "hello world"}).Warn("Search by content input")

	c.SearchAndInputNews()

	return nil

}
