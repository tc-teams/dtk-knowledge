package tracker

import (
	"github.com/tc-teams/fakefinder-crawler/tracker/crawler"
)

func WebCrawlerNews() error {

	g1 := crawler.NewG1()
	g1.TrackNewsBasedOnCovid19()
	err := g1.LoggingDocuments()
	if err != nil {
		return err
	}

	//TODO implements logs about any sites

	return nil

}
