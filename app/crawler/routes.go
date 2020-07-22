package crawler

import (
	"encoding/json"
	"fmt"
	"github.com/gocolly/colly"
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

	return r
}

func NewsRelatedToCovid(w http.ResponseWriter, r *http.Request) *api.BaseError {
	var c DocumentNews

	err := json.NewDecoder(r.Body).Decode(&c)
	if err != nil {
		return &api.BaseError{
			Error:   err,
			Message: "Invalid request body",
			Code:    http.StatusBadRequest,
		}
	}


	err = HandlerCovid(c)
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

func HandlerCovid(d DocumentNews) error {

	tk := tracker.New(colly.NewCollector(
		colly.AllowedDomains(
		tracker.Folha,
		tracker.G1,
		tracker.Uol),
	),
		&logrus.Logger{
			Out:       os.Stdout,
			Formatter: &logrus.JSONFormatter{},
		}, nil,
	)

	logrus.WithFields(logrus.Fields{
		"Title": d.Title,
		"SubTitle":  d.SubTitle,
		"Url": d.Url,
	}).Warn("Init search by")

	err, news := tk.SearchAndInputNews()
	if err != nil {
		return err
	}
    for i := 0 ; i < len(news); i++{
    	fmt.Printf("%+v\n",news[i])
	}

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
	//return nil

}
