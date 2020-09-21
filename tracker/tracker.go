package tracker

import (
	"github.com/sirupsen/logrus"
	"github.com/tc-teams/fakefinder-crawler/api"
	"github.com/tc-teams/fakefinder-crawler/tracker/crawler"
)

func WebCrawlerNews(log *api.Logging) error {

	g1 := crawler.NewG1()

	log.WithFields(logrus.Fields{"page": crawler.StartG1,}).Debug("starting synchronization")
	g1.TrackNewsBasedOnCovid19()
	log.WithFields(logrus.Fields{"page": crawler.StartG1,}).Debug("synchronization finished")

	if err := g1.LoggingDocuments(log); err != nil {
		return err
	}
	gv := crawler.NewGov()

	log.WithFields(logrus.Fields{"page": crawler.StartGV,}).Debug("starting synchronization")
	gv.TrackNewsBasedOnCovid19()
	log.WithFields(logrus.Fields{"page": crawler.StartGV,}).Debug("synchronization finished")

	if err := gv.LoggingDocuments(log); err != nil {
		return err

	}

	return nil

}
