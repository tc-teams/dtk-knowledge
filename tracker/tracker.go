package tracker

import (
	"github.com/sirupsen/logrus"
	"github.com/tc-teams/fakefinder-crawler/api"
	"github.com/tc-teams/fakefinder-crawler/tracker/crawler"
	"time"
)

func WebCrawlerNews(log *api.Logging) error {

	g1 := crawler.NewG1()
	log.WithFields(logrus.Fields{}).Info("starting synchronization")
	g1.TrackNewsBasedOnCovid19()

	log.WithFields(logrus.Fields{}).Info("synchronization finished")
	time.Sleep(20 * time.Second)
	err := g1.LoggingDocuments(log)
	if err != nil {
		return err
	}
	//TODO implements logs about any sites

	return nil

}
