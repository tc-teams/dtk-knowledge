package collector

import (
	"github.com/gocolly/colly/v2"
	validate "github.com/idasilva/dtk-knowledge/app/news/valid"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

const (
	visited = 10
)

type Collector struct {
	Colly     *colly.Collector
	Log       *log.Logger
	Valid      validate.Validation
	Content    string
}

type News struct {
	Title    string `validate:"required,max=60"`
	SubTitle string `validate:"required,max=60"`
	Date     string `validate:"required,max=60"`
	Page     string `validate:"required,max=60"`
}


//LoadNews returns related news by an entry
func (c *Collector) SearchAndInputNews() {
	detailColly := c.Colly.Clone()
	stop := false

	c.Colly.OnRequest(func(r *colly.Request) {
		c.Log.WithFields(log.Fields{"Visiting": r.URL.String()}).Info("start search")
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
		if strings.Index(subUrl, "covid") > -1 || strings.Index(subUrl, "coronavirus") > -1 ||
			strings.Index(subUrl, "covid-19") > -1 && !stop {
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
		_,err := c.Valid.ValidateStruct(detailsNews)
		if err != nil {
			return
		}

		c.Log.WithFields(log.Fields{
			"Title":    detailsNews.Title,
			"SubTitle": detailsNews.SubTitle,
			"Page":     detailsNews.Page,
		}).Info(c.Content)

	})
	c.Colly.Limit(&colly.LimitRule{DomainGlob: "*", Parallelism: 3, RandomDelay: 5 * time.Second})

	// Start scraping on....
	c.Colly.Visit(StartFolha)
	c.Colly.Visit(StartG1)
	c.Colly.Visit(StartUol)
	c.Colly.Wait()

}

//NewCollector return new  instance of colly
func NewColly(Colly *colly.Collector, Log *log.Logger,Valid validate.Validation, Content string) *Collector {

	return &Collector{
		Colly:     Colly,
		Log:       Log,
		Valid:     Valid,
		Content:   Content,

	}

}
