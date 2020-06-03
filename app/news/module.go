package news

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	collector "github.com/idasilva/dtk-knowledge/collector"
	log "github.com/sirupsen/logrus"
	"net/http"
	"os"
)

//HandlerFakeFinder instance a new collector of news
func HandlerFakeFinder(w http.ResponseWriter,r *http.Request) {

	c := collector.NewColly(colly.NewCollector(),&log.Logger{
		Out:          os.Stdout,
		Formatter:    &log.JSONFormatter{},
		Level:        log.DebugLevel,
	})


	c.LoadNews()

	fmt.Fprint(w,"New instace of colly was created")


}

