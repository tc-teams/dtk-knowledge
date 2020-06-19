package collector

import (
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
	"github.com/tc-teams/fakefinder-crawler/app/news/valid"
	"strings"
	"time"
)

type Collector struct {
	Colly   *colly.Collector
	Log     *log.Logger
	Valid   *valid.Validation
	Content string
}

//LoadNews returns related news by an entry
func (c *Collector) SearchAndInputNews() {
	detailColly := c.Colly.Clone()
	stop := false

	c.Colly.OnRequest(func(r *colly.Request) {
		c.Log.WithFields(log.Fields{"Visiting": r.URL.String()}).Info("int search ")
	})
	c.Colly.OnScraped(func(r *colly.Response) {
		c.Log.WithFields(log.Fields{"Finished": r.Request.URL}).Info("end search")
	})

	c.Colly.OnError(func(_ *colly.Response, e error) {
		stop = true
		c.Log.WithFields(log.Fields{"Error": e.Error()}).Info("error search")

	})

	c.Colly.OnHTML("a[href]", func(e *colly.HTMLElement) {
		subUrl := e.Request.AbsoluteURL(e.Attr("href"))
		if strings.Index(subUrl, "Covid") > -1 || strings.Index(subUrl, "coronavírus") > -1 ||
			strings.Index(subUrl, "Covid-19") > -1 || strings.Index(subUrl, "pandemia") > -1 || strings.Index(subUrl, "quarentena") > -1 && !stop {
			detailColly.Visit(subUrl)
		}
	})
	detailColly.OnHTML("head", func(e *colly.HTMLElement) {

		detailsNews := News{}
		e.ForEach("meta", func(_ int, el *colly.HTMLElement) {
			switch el.Attr("property") {
			case "og:title":
				detailsNews.Title = el.Attr("content")
			case "og:description":
				detailsNews.SubTitle = el.Attr("content")
			}
		})
		detailsNews.Page = e.Request.URL.Host
		_, err := c.Valid.ValidateStruct(detailsNews)

		if err != nil {
			c.Log.WithFields(log.Fields{
				"ErrorID": err,
			}).Info(c.Content)

			return
		}

		c.Log.WithFields(log.Fields{
			"Title":    detailsNews.Title,
			"SubTitle": detailsNews.SubTitle,
			"Page":     detailsNews.Page,
		}).Info(c.Content)

	})
	c.Colly.Limit(&colly.LimitRule{Parallelism: 3, RandomDelay: 1 * time.Second})

	// Start scraping on....
	c.Colly.Visit(StartFolha)
	c.Colly.Visit(StartG1)
	c.Colly.Visit(StartUol)
	c.Colly.Wait()

}

//NewCollector return news  instance of colly
func NewColly(Colly *colly.Collector, Log *log.Logger, Valid *valid.Validation, Content string) *Collector {

	return &Collector{
		Colly:   Colly,
		Log:     Log,
		Valid:   Valid,
		Content: Content,
	}

}
