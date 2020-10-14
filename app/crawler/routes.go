package crawler

import (
	"encoding/json"
	"errors"
	"github.com/sirupsen/logrus"
	"github.com/tc-teams/fakefinder-crawler/api"
	"github.com/tc-teams/fakefinder-crawler/elastic"
	"github.com/tc-teams/fakefinder-crawler/external"
	"github.com/tc-teams/fakefinder-crawler/nlp"
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
		Path:    "/search/news",
		Handler: ElasticCrawlByDescription,
	},
	)

	return r
}

func CrawlNewsRelatedToCovid(w http.ResponseWriter, r *http.Request, log *api.Logging) *api.BaseError {

	err := tracker.WebCrawlerNews(log)
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

	w.Write([]byte("Search news success"))
	return nil
}

func ElasticCrawlByDescription(w http.ResponseWriter, r *http.Request, log *api.Logging) *api.BaseError {

	var info external.BotRequest

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
		errNoDataFound := errors.New("No data found")
		return &api.BaseError{
			Error:   errNoDataFound,
			Message: "No data found",
			Code:    http.StatusNotFound,
		}

	}
	reqBody := &external.PlnRequest{
		Description: info.Description,
	}

	for _, related := range documents {
		reqBody.News = append(reqBody.News, related.News.Body)
	}

	response, err := external.NewClient().Request(reqBody)
	if err != nil {
		return &api.BaseError{
			Error:   err,
			Message: "error request a external service",
			Code:    http.StatusBadGateway,
		}
	}
	defer response.Body.Close()


	var pln external.PlnResponse

	err = json.NewDecoder(response.Body).Decode(&pln)
	if err != nil {
		return &api.BaseError{
			Error:   err,
			Message: "Invalid request body nlp",
			Code:    http.StatusBadRequest,
		}
	}

	bot := nlp.NaturalLanguageProcess(pln,documents)


	log.WithFields(logrus.Fields{
		"nlp": "external process successfully completed",
	}).Info()

	err = json.NewEncoder(w).Encode(bot)
	if err != nil {
		errInvalidEncode := errors.New("was not possible enconde response body")
		return &api.BaseError{
			Error:   err,
			Message: errInvalidEncode.Error(),
			Code:    http.StatusBadRequest,
		}
	}
	w.WriteHeader(http.StatusOK)
	return nil

}
