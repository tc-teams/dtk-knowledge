package crawler

type Crawler interface {
	TrackNewsBasedOnCovid19()
	LoggingDocuments() error

}