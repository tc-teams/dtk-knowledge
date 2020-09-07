package crawler

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/tc-teams/fakefinder-crawler/api"
	"github.com/tc-teams/fakefinder-crawler/elastic"
	"github.com/tc-teams/fakefinder-crawler/elastic/es"
	"github.com/tc-teams/fakefinder-crawler/tracker"
	"net/http"
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
	var information es.Info

	err := json.NewDecoder(r.Body).Decode(&information)
	if err != nil {
		return &api.BaseError{
			Error:   err,
			Message: "Invalid request body",
			Code:    http.StatusBadRequest,
		}
	}

	err = tracker.WebCrawlerNews(log)
	if err != nil {
		return &api.BaseError{
			Error:   err,
			Message: "The process could not be completed, pages was not found",
			Code:    http.StatusNotFound,
		}
	}
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

	for index, related := range documents {

		log.WithFields(logrus.Fields{
			"Url":      related.News.Url,
			"Date":     related.News.Date,
			"Title":    related.News.Title,
			"SubTitle": related.News.Subtitle,
			"Body":     related.News.Body,
		}).Info("News related:", index)

	}
	//var (
	//	request http.Request
	//	result []
	//)
	//
	//err = external.NewClient().Request(&request,result)
	//if err != nil {
	//	return err
	//}

	return &api.BaseError{
		Error:   nil,
		Message: "OK",
		Code:    http.StatusOK,
	}

}
