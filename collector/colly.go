package collector

import (
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

type News struct {
	Title    string
	SubTitle string
	Date     string
}

//LoadNews returns related news by an entry
func (c *Collector) SearchAndInputNews() {
	detailCollector := c.Collector.Clone()

	c.Collector.OnRequest(func(r *colly.Request) {
		c.Log.WithFields(log.Fields{"Visiting": r.URL.String()}).Info("start search")
	})
	c.Collector.OnScraped(func(r *colly.Response) {
		c.Log.WithFields(log.Fields{"Finished": r.Request.URL}).Info("end search")
	})
	c.Collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		subUrl := e.Request.AbsoluteURL(e.Attr("href"))
		if strings.Index(subUrl, "covid") > -1 || strings.Index(subUrl, "coronavirus") > -1 ||
			strings.Index(subUrl, "covid-19") > -1 {
			detailCollector.Visit(subUrl)
		}
	})
	detailCollector.OnHTML("head", func(e *colly.HTMLElement) {

		detailsNews := News{}
		e.ForEach("meta", func(_ int, el *colly.HTMLElement) {
			switch el.Attr("property") {

			case "og:title":
				detailsNews.Title = el.Attr("content")
			case "og:description":
				detailsNews.SubTitle = el.Attr("content")
			}
		})
		c.Log.WithFields(log.Fields{"Title": detailsNews.Title,"SubTitle": detailsNews.SubTitle}).Info(c.Content)

	})

	// Start scraping on....
	c.Collector.Visit(StartFolha)
	c.Collector.Visit(StartG1)
	c.Collector.Visit(StartUol)
	c.Collector.Wait()

}
//NewCollector return new  instance of colly
func NewColly(Colly *colly.Collector, Logging *log.Logger, Content string) *Collector {

	return &Collector{
		Collector: Colly,
		Log:       Logging,
		Content:   Content,
	}

}

