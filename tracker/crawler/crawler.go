package crawler

import "github.com/tc-teams/fakefinder-crawler/api"

type Crawler interface {
	TrackNewsBasedOnCovid19()
	LoggingDocuments(log *api.Logging) error

}