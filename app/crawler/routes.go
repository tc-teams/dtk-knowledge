package crawler

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/tc-teams/fakefinder-crawler/api"
	"github.com/tc-teams/fakefinder-crawler/elastic"
	"github.com/tc-teams/fakefinder-crawler/tracker"
	"github.com/tc-teams/fakefinder-crawler/tracker/crawl"
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

func CrawlNewsRelatedToCovid(w http.ResponseWriter, r *http.Request) *api.BaseError {
	var information crawl.Info

	err := json.NewDecoder(r.Body).Decode(&information)
	if err != nil {
		return &api.BaseError{
			Error:   err,
			Message: "Invalid request body",
			Code:    http.StatusBadRequest,
		}
	}

	err = tracker.WebCrawlerNews()
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

func ElasticCrawlByDescription(w http.ResponseWriter, r *http.Request) *api.BaseError {

	documents, err := elastic.ElasticDocumentsByDescription()
	if err != nil {
		return &api.BaseError{
			Error:   err,
			Message: "Elastic error",
			Code:    http.StatusNotFound,
		}
	}

	for index, news := range documents {

		logrus.WithFields(logrus.Fields{
			"Url":      news.Url,
			"Date":     news.Date,
			"Title":    news.Title,
			"SubTitle": news.Subtitle,
			"Body":     news.Body,
		}).Warn("News related:", index)

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
