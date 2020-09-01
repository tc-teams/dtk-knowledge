package crawl

import (
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
	"time"
)

type Crawl struct {
	Colly *colly.Collector
}

//LoadNews returns related crawler by an entry
func (t *Crawl) TrackNewsBasedOnCovid19() (error, []RelatedNews) {

	news := make([]RelatedNews, 0, 2)
	detailsNews := RelatedNews{}
	stop := false
	var i int
	t.Colly.OnRequest(func(r *colly.Request) {
		log.WithFields(log.Fields{
			"Visiting": r.URL.String()}).Info("int search")
	})
	t.Colly.OnScraped(func(r *colly.Response) {
		log.WithFields(log.Fields{
			"Finished": r.Request.URL}).Info("end search")
	})

	t.Colly.OnError(func(_ *colly.Response, e error) {
		stop = true
		log.WithFields(log.Fields{
			"Error": e.Error()}).Info("error search")

	})

	t.Colly.OnHTML("#bstn-launcher a[href]", func(e *colly.HTMLElement) {
		e.Request.Visit(e.Attr("href"))
		i +=1
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
			log.WithFields(log.Fields{
				"Date": time,
			}).Info("Was not possible convert string to Date")

		}
		if detailsNews.Date.IsZero() {
			detailsNews.Date = time
		}
        news = append(news, detailsNews)

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
