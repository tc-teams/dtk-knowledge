package tracker

import (
	"github.com/sirupsen/logrus"
	"github.com/tc-teams/fakefinder-crawler/api"
	"github.com/tc-teams/fakefinder-crawler/tracker/crawler"
)

func WebCrawlerNews(log *api.Logging) error {

	g1 := crawler.NewG1()
	g1.TrackNewsBasedOnCovid19()

	log.WithFields(logrus.Fields{}).Info("synchronization finished")

	err := g1.LoggingDocuments(log)
	if err != nil {
		return err
	}
	//TODO implements logs about any sites

	return nil

}
