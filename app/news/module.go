package news

import (
	"github.com/gocolly/colly/v2"
	"github.com/gorilla/mux"
	collector "github.com/idasilva/dtk-knowledge/collector"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

//Limiting Colly to parse only links that are on the clienturl.com domain
//Turning on Async processing of links (this is where we get a HUGE speed increase as we'll talk about in a bit)

//HandlerFakeFinder instance a new collector of news
func HandlerFakeFinder(w http.ResponseWriter, r *http.Request) {
	param := mux.Vars(r)
	c := collector.NewColly(colly.NewCollector(
		//colly.AllowedDomains(collector.Folha,collector.G1,collector.Uol),
		colly.Async(true),
	),
		&log.Logger{
			Out:       os.Stdout,
			Formatter: &log.JSONFormatter{},
			Level:     log.DebugLevel,
		}, param["content"])

	log.WithFields(log.Fields{
		"Text":  param["content"],
	}).Warn("Search by content input")

	c.SearchAndInputNews()

	w.WriteHeader(http.StatusOK)

}
