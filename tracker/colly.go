package tracker

import (
	"github.com/gocolly/colly"
	log "github.com/sirupsen/logrus"
	"strings"
	"time"
)

const LayouISO8601 = "01/02/2006"

type Track struct {
	Colly *colly.Collector
}

//LoadNews returns related crawler by an entry
func (t *Track) TrackNewsBasedOnUrl(bind string) (error, []RelatedNews) {

	detailColly := t.Colly.Clone()
	news := make([]RelatedNews, 0, 2)
	stop := false

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

	t.Colly.OnHTML("a[href]", func(e *colly.HTMLElement) {
		subUrl := e.Request.AbsoluteURL(e.Attr("href"))
		if strings.Index(subUrl, "Covid") > -1 || strings.Index(subUrl, "coronavírus") > -1 ||
			strings.Index(subUrl, "Covid-19") > -1 || strings.Index(subUrl, "pandemia") > -1 ||
			strings.Index(subUrl, "quarentena") > -1 && !stop {
			detailColly.Visit(subUrl)
		}
	})
	detailColly.OnHTML("head", func(e *colly.HTMLElement) {

		detailsNews := RelatedNews{}
		e.ForEach("meta", func(_ int, el *colly.HTMLElement) {
			switch el.Attr("property") {
			case "og:title":
				detailsNews.Title = el.Attr("content")
			case "og:description":
				detailsNews.Subtitle = el.Attr("content")
			}
		})
		detailsNews.Url = e.Request.URL.Host
		news = append(news, detailsNews)

	})
	t.Colly.Limit(&colly.LimitRule{Parallelism: 3, RandomDelay: 1 * time.Second})

	t.Colly.Visit(StartFolha)
	t.Colly.Visit(StartG1)
	t.Colly.Visit(StartUol)
	t.Colly.Wait()

	return nil, news
}

func (t *Track) TrackNewsOnUrl(baseUrl string) (RelatedNews, error) {
	detailsNews := RelatedNews{}

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

	})

	t.Colly.Visit(baseUrl)
	t.Colly.Wait()

	return detailsNews, nil
}

//New return crawler  instance of colly
func New() *Track {
	return &Track{
		Colly: colly.NewCollector(colly.AllowedDomains(Folha, G1, Uol)),
	}

}
