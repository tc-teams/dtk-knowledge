package collector

import (
	"fmt"
	"github.com/gocolly/colly/v2"
)

type Collector struct {
	collector *colly.Collector

}

//NewCollector return new colly
func NewColly(colly *colly.Collector) *Collector{

	return &Collector{
		collector: colly,
	}

}

//LoadNews NEWS
func(c *Collector) LoadNews(){

	c.collector.OnHTML("a[href]", func(e *colly.HTMLElement) {
		link := e.Attr("href")
		fmt.Printf("Link found: %q -> %s\n", e.Text, link,e.DOM)

		c.collector.Visit(e.Request.AbsoluteURL(link))
	})

	c.collector.OnRequest(func(r *colly.Request) {
		fmt.Println("Visiting", r.URL)
	})

	c.collector.Visit("https://g1.globo.com/")


}


