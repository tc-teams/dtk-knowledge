package collector

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	log "github.com/sirupsen/logrus"
)

type Collector struct {
	collector *colly.Collector
	log  *log.Logger
}

//NewCollector return new colly
func NewColly(colly *colly.Collector, logging *log.Logger) *Collector {

	return &Collector{
		collector: colly,
		log:   logging,
	}

}

//LoadNews NEWS
func (c *Collector) LoadNews() {

	c.collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		c.log.WithFields(log.Fields{
			"Text": e.Text,
			"link": link,
		}).Info("Page was found")

		c.collector.Visit(e.Request.AbsoluteURL(link))
	})

	c.collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.collector.Visit("https://g1.globo.com/")

}
