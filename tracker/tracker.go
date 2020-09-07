package tracker

import (
	"github.com/sirupsen/logrus"
	"github.com/tc-teams/fakefinder-crawler/api"
	"github.com/tc-teams/fakefinder-crawler/tracker/crawl"
)

func WebCrawlerNews(log *api.Logging) error {

	cl := crawl.NewCrawler()

	err, related := cl.TrackNewsBasedOnCovid19()
	if err != nil {
		return err
	}

	for index, news := range related {

		log.WithFields(logrus.Fields{
			"Url":      news.Url,
			"Date":     news.Time,
			"Title":    news.Title,
			"SubTitle": news.Subtitle,
			"Body":     news.Body,
		}).Info("News related:",index)

	}
	return nil

}
