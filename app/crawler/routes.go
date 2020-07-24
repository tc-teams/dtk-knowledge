package crawler

import (
	"encoding/json"
	"github.com/sirupsen/logrus"
	"github.com/tc-teams/fakefinder-crawler/api"
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
		Handler: SearchNewsRelatedToCovid,
	})

	return r
}

func SearchNewsRelatedToCovid(w http.ResponseWriter, r *http.Request) *api.BaseError {
	var b BaseUrl

	err := json.NewDecoder(r.Body).Decode(&b)
	if err != nil {
		return &api.BaseError{
			Error:   err,
			Message: "Invalid request body",
			Code:    http.StatusBadRequest,
		}
	}

	err = HandlerCovid(b)
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

func HandlerCovid(base BaseUrl) error {

	tk := tracker.New()

	related, err := tk.TrackNewsOnUrl(base.Url)
	if err != nil {
		return err
	}

	logrus.WithFields(logrus.Fields{
		"Url":      related.Url,
		"Date":     related.Date,
		"Title":    related.Title,
		"SubTitle": related.Subtitle,
		"Body":     related.Body,
	}).Warn("Create a new crawler based on url")

	//err, _ = tk.TrackNewsBasedOnUrl(related.Title)
	//if err != nil {
	//	return err
	//}

	//var (
	//	request http.Request
	//	result []DocRequest
	//)
	//
	//err = external.NewClient().Request(&request,result)
	//if err != nil {
	//	return err
	//}
	//
	return nil

}
