package crawler

import (
	"fmt"
	"github.com/gocolly/colly"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
	"github.com/tc-teams/fakefinder-crawler/api"
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
		Method:  http.MethodPost,
		Path:    "/covid",
		Handler: NewsRelatedToCovid,
	})
	fmt.Printf("adicionando a rota %v\n", r)

	return r
}


func NewsRelatedToCovid(w http.ResponseWriter, r *http.Request) {

	param := mux.Vars(r)

	//validation := api.NewValidate()

	c := tracker.NewColly(colly.NewCollector(
		colly.AllowedDomains(tracker.Folha, tracker.G1, tracker.Uol),
		colly.MaxDepth(3),
		colly.Async(true),
	),
		&logrus.Logger{
			Out:       os.Stdout,
			Formatter: &logrus.JSONFormatter{},
		}, nil, param["content"],
	)

	logrus.WithFields(logrus.Fields{"Text": param["content"]}).Warn("Search by content input")

	c.SearchAndInputNews()

	w.WriteHeader(http.StatusOK)

}
