package crawl

import (
	"github.com/gocolly/colly"
	"github.com/sirupsen/logrus"
	"time"
)
var (
	stop = bool(false)
)
type Crawl struct {
	Colly *colly.Collector
}

//LoadNews returns related crawler by an entry
func (t *Crawl) TrackNewsBasedOnCovid19() (error, []RelatedNews) {

	news := make([]RelatedNews, 0, 2)
	detailsNews := RelatedNews{}

	t.Colly.OnHTML("#bstn-launcher a[href]", func(e *colly.HTMLElement) {
		if !stop {
			e.Request.Visit(e.Attr("href"))
		}

	})

	t.Colly.OnHTML("head", func(e *colly.HTMLElement) {


		e.ForEach("meta", func(_ int, el *colly.HTMLElement) {
			switch el.Attr("property") {
			case "og:title":
				detailsNews.Title = el.Attr("content")
			case "og:description":
				detailsNews.Subtitle = el.Attr("content")
			}
		})
		detailsNews.Url = e.Request.URL.Host

	})
	t.Colly.OnHTML("main", func(e *colly.HTMLElement) {

		e.ForEach("p", func(_ int, el *colly.HTMLElement) {
				detailsNews.Body += el.Text


		})

	})
	t.Colly.OnHTML("time", func(e *colly.HTMLElement) {

		time, err := time.Parse(time.RFC3339, e.Attr("datetime"))
		if err != nil {
			logrus.WithFields(logrus.Fields{
				"Date": time,
			}).Info("Was not possible convert string to Date")

		}
		if detailsNews.Date.IsZero() {
			detailsNews.Date = time
		}
        news = append(news, detailsNews)
        if len(news) == 100{
        	stop = true
		}

	})

	t.Colly.Limit(&colly.LimitRule{Parallelism: 3, RandomDelay: 1 * time.Second})


	t.Colly.Visit(StartG1)
	t.Colly.Wait()

	return nil, news
}

//New return crawler  instance of colly
func NewCrawler() *Crawl {
	return &Crawl{
		Colly: colly.NewCollector(colly.AllowedDomains(Folha, G1, Uol)),
	}

}
