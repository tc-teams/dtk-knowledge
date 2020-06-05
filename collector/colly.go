package collector

import (
	"github.com/gocolly/colly/v2"
	log "github.com/sirupsen/logrus"
)

const (
	visited = 10
)

type Collector struct {
	collector *colly.Collector
	log       *log.Logger
	domain    string
	content   string
}

//NewCollector return new  instance of colly
func NewColly(colly *colly.Collector, logging *log.Logger) *Collector {

	return &Collector{
		collector: colly,
		log:       logging,
	}

}

//LoadNews returns related news by an entry
func (c *Collector) LoadNews() {

	c.collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")

		c.log.WithFields(log.Fields{
			"Text":  e.Text,
			"link":  link,
			"based": "on",
		}).Info(c.content)

		c.collector.Visit(e.Request.AbsoluteURL(link))
	})
	countVisited := 0
	c.collector.OnRequest(func(r *colly.Request) {

		if countVisited >= visited {
			r.Abort()
		}

		countVisited++
	})

	c.collector.Visit(c.domain)

}
