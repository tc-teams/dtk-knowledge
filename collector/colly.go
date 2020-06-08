package collector

import "C"
import (
	"fmt"
	"github.com/PuerkitoBio/goquery"
	"github.com/gocolly/colly/v2"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

const (
	visited = 10
)

type Collector struct {
	Collector *colly.Collector
	Log       *log.Logger
	Content   string
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

	// Callback for when a scraped page contains an article element
	c.Collector.OnHTML("article", func(e *colly.HTMLElement) {

		// Extract meta tags from the document
		metaTags := e.DOM.ParentsUntil("~").Find("meta")
		metaTags.Each(func(_ int, s *goquery.Selection) {
			// Search for og:type meta tags
			property, _ := s.Attr("property")
			if strings.EqualFold(property, "og:type") {
				content, _ := s.Attr("content")

				// Emoji pages have "article" as their og:type
				isEmojiPage := strings.EqualFold(content, "article")
				fmt.Println(isEmojiPage)
			}
		})


	})

	c.Collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")


		c.Log.WithFields(log.Fields{
			"Text":  e.Text,
			"link":  link,
			"based": "on",
		}).Info(c.Content)
		c.Collector.Visit(e.Request.AbsoluteURL(link))

	})
	c.Collector.Limit(&colly.LimitRule{
		DomainGlob:  "*" ,
		RandomDelay: 5 * time.Second,
	})

	countVisited := 0
	c.Collector.OnRequest(func(r *colly.Request) {

		if countVisited >= visited {
			r.Abort()
		}

		countVisited++
	})

	// Start scraping on....
	c.Collector.Visit("https://g1.globo.com/")

}
