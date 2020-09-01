package tracker

import (
	"github.com/sirupsen/logrus"
	"github.com/tc-teams/fakefinder-crawler/tracker/crawl"
)

func WebCrawlerNews() error {

	cl := crawl.NewCrawler()

	err, related := cl.TrackNewsBasedOnCovid19()
	if err != nil {
		return err
	}

	for index, news := range related {

		logrus.WithFields(logrus.Fields{
			"Url":      news.Url,
			"Date":     news.Date,
			"Title":    news.Title,
			"SubTitle": news.Subtitle,
			"Body":     news.Body,
		}).Warn("News related:",index)

	}
	return nil

}
