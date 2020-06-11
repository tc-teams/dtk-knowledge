package collector

import (
	"fmt"
	"github.com/gocolly/colly/v2"
	log "github.com/sirupsen/logrus"
	"strings"
)

const (
	visited = 10
)

type Collector struct {
	Collector *colly.Collector
	Log       *log.Logger
	Content   string
}

type Te struct {
	Title   string
	SubTile string
}

//NewCollector return new  instance of colly
func NewColly(Colly *colly.Collector, Logging *log.Logger, Content string) *Collector {

	return &Collector{
		Collector: Colly,
		Log:       Logging,
		Content:   Content,
	}

}

//LoadNews returns related news by an entry
func (c *Collector) SearchAndInputNews() {


	c.Collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL.String())
	})
	c.Collector.OnScraped(func(r *colly.Response) {
		fmt.Println("Finished", r.Request.URL)
	})
	c.Collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		subUrl := e.Request.AbsoluteURL(e.Attr("href"))
		if strings.Index(subUrl, "covid") > -1 || strings.Index(subUrl, "coronavirus") > -1 ||
			strings.Index(subUrl, "covid-19") > -1 {
			c.Collector.Visit(subUrl)
			fmt.Println("sub url permitida", subUrl)
		}
	})
	c.Collector.OnHTML("head", func(e *colly.HTMLElement) {

		teste := Te{}
		e.ForEach("meta", func(_ int, el *colly.HTMLElement) {

			switch el.Attr("property") {
			case "og:title":
				teste.Title = el.Attr("content")
			case "og:description":
				teste.SubTile = el.Attr("content")
			}
		})
		c.Log.WithFields(log.Fields{
			"Title":    teste.Title,
			"SubTitle": teste.SubTile,
			"based":    "on",
		}).Info(c.Content)
	})

	// Start scraping on....
	c.Collector.Visit("https://g1.globo.com/busca/?q=covid")

}
