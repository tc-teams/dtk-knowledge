package crawler

import (
	"errors"
	"github.com/gocolly/colly"
	"github.com/sirupsen/logrus"
	"time"
)

type G1 struct {
	Colly *colly.Collector
	News  []RelatedNews
}

//LoadNews returns related crawler by an entry
func (g *G1) TrackNewsBasedOnCovid19() {
	detailsNews := RelatedNews{}

	g.Colly.OnHTML("#bstn-launcher a[href]", func(e *colly.HTMLElement) {
		if !stop {
			e.Request.Visit(e.Attr("href"))
			detailsNews.Url = e.Attr("href")

		}

	})

	g.Colly.OnHTML("head", func(e *colly.HTMLElement) {

		e.ForEach("meta", func(_ int, el *colly.HTMLElement) {
			switch el.Attr("property") {
			case "og:title":
				detailsNews.Title = el.Attr("content")
			case "og:description":
				detailsNews.Subtitle = el.Attr("content")
			}
		})

	})
	g.Colly.OnHTML("main", func(e *colly.HTMLElement) {
		e.ForEach("p", func(_ int, el *colly.HTMLElement) {
			text := el.Text
			detailsNews.Body += text



		})

	})
	g.Colly.OnHTML("time", func(e *colly.HTMLElement) {

		time, err := time.Parse(time.RFC3339, e.Attr("datetime"))
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"Date": time,
			}).Info("Was not possible convert string to Date")

		}
		if detailsNews.Date.IsZero() {
			detailsNews.Date = time
		}
		g.News = append(g.News, detailsNews)
		if len(g.News) == 2 {
			stop = true
		}
		logrus.WithFields(logrus.Fields{
			"title": detailsNews.Title,
			"Body":  detailsNews.Body,
			"url": detailsNews.Url,
		}).Info()
		detailsNews.Body = ""

	})

	g.Colly.Limit(&colly.LimitRule{Parallelism: 3, RandomDelay: 1 * time.Second})

	g.Colly.Visit(StartG1)
	g.Colly.Wait()

}
func (g *G1) LoggingDocuments() error {
	if g.News == nil {
		return errors.New("error to search data in G1")

	}
	for index, news := range g.News {

		logrus.WithFields(logrus.Fields{
			"Url":      news.Url,
			"Date":     news.Time,
			"Title":    news.Title,
			"SubTitle": news.Subtitle,
			"Body":     news.Body,
		}).Info("News related:", index)

	}
	return nil
}

//NewG1 return crawler  instance of colly
func NewG1() Crawler {
	return &G1{
		Colly: colly.NewCollector(colly.AllowedDomains(Folha, GB, Uol)),
		News:  nil,
	}
}
