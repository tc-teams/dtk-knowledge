package crawler

import (
	"encoding/json"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/tc-teams/fakefinder-crawler/api"
	"github.com/tc-teams/fakefinder-crawler/elastic"
	"github.com/tc-teams/fakefinder-crawler/elastic/es"
	"github.com/tc-teams/fakefinder-crawler/external"
	"github.com/tc-teams/fakefinder-crawler/tracker"
	"net/http"
	"time"
)

const routerName = "crawler"

//Init return a new route instance
func Init() *api.Route {
	r := &api.Route{}
	r.RouteName(routerName)
	r.AddRoute(&api.ContextRoute{
		Method:  http.MethodPost,
		Path:    "/search/news/related/covid",
		Handler: CrawlNewsRelatedToCovid,
	}, &api.ContextRoute{
		Method:  http.MethodPost,
		Path:    "/teste",
		Handler: ElasticCrawlByDescription,
	},
	)

	return r
}

func CrawlNewsRelatedToCovid(w http.ResponseWriter, r *http.Request, log *api.Logging) *api.BaseError {

	err := tracker.WebCrawlerNews()
	if err != nil {
		return &api.BaseError{
			Error:   err,
			Message: "The process could not be completed, pages was not found",
			Code:    http.StatusNotFound,
		}
	}
	log.WithFields(logrus.Fields{
		"WebCrawler": "Search news success",
	}).Info()

	return &api.BaseError{
		Error:   nil,
		Message: "OK",
		Code:    http.StatusOK,
	}

}


func ElasticCrawlByDescription(w http.ResponseWriter, r *http.Request, log *api.Logging) *api.BaseError {

	var info es.Info

	err := json.NewDecoder(r.Body).Decode(&info)
	if err != nil {
		return &api.BaseError{
			Error:   err,
			Message: "Invalid request body",
			Code:    http.StatusBadRequest,
		}
	}

	documents, err := elastic.DocumentsByDescription(log, info.Description)
	if err != nil {
		return &api.BaseError{
			Error:   err,
			Message: "Elastic error",
			Code:    http.StatusNotFound,
		}
	}

	if documents == nil {
		return &api.BaseError{
			Error:   err,
			Message: "No data found",
			Code:    http.StatusNotFound,
		}

	}
	reqBody := &external.BotRequest{
		Description: info.Description,
	}

	for _, related := range documents {
		reqBody.Related = append(reqBody.Related, related.News.Body)
	}
	fmt.Println("quantidade:",len(reqBody.Related))
	time.Sleep(10*time.Second)


	for q, i := range reqBody.Related{
		fmt.Printf("Documents%s: %s",q,i)

	}



	//related, err := external.NewClient().Request(reqBody)
	//if err != nil {
	//	return &api.BaseError{
	//		Error:   err,
	//		Message: "error request a external service",
	//		Code:    http.StatusBadGateway,
	//	}
	//}
	//fmt.Println(related)

	return &api.BaseError{
		Error:   nil,
		Message: "OK",
		Code:    http.StatusOK,
	}

}
