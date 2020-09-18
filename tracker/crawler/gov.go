package crawler

import (
	"errors"
	"github.com/gocolly/colly"
	"github.com/sirupsen/logrus"
	"github.com/tc-teams/fakefinder-crawler/api"
	ctx "github.com/tc-teams/fakefinder-crawler/context/validator"
	"regexp"
	"strings"
	"time"
)


type GOV struct {
	Colly     *colly.Collector
	News      []RelatedNews
	validator *ctx.Validation
	Log       *api.Logging
}
var(
	Govstop      = false
)

//LoadNews returns related crawler by an entry
func (g *GOV) TrackNewsBasedOnCovid19() {
	detailsNews := RelatedNews{}

	//#content .cat-items .tile-list-1 .tileItem
	g.Colly.OnHTML(".category-listnoticias a[href] ", func(e *colly.HTMLElement) {
		if !Govstop {
			e.Request.Visit(e.Attr("href"))
		}

	})

	g.Colly.OnHTML("title", func(e *colly.HTMLElement) {
		detailsNews.Title = e.Text

	})
	g.Colly.OnHTML(".item-pagenoticias", func(e *colly.HTMLElement) {
		e.ForEach("p", func(_ int, el *colly.HTMLElement) {
			text := el.Text
			detailsNews.Body += text

		})

	})
	g.Colly.OnHTML(".documentPublished", func(e *colly.HTMLElement) {
		data := strings.Split(e.Text,",")

		if detailsNews.Date == strEmpty{
			detailsNews.Date = data[1]
		}

		hasNotice, err := g.validator.ValidateStruct(detailsNews)

		if err != nil {
			detailsNews = RelatedNews{}
			return
		}

		for _, re := range g.News {
			if re.Title == detailsNews.Title {
				hasNotice = false
				break

			}
		}
		if hasNotice {
			detailsNews.Url = e.Request.AbsoluteURL(e.Request.URL.Path)
			g.News = append(g.News, detailsNews)
		}


		if len(g.News) == 2 {
			Govstop = true
			return
		}
		detailsNews = RelatedNews{}

	})

	g.Colly.Limit(&colly.LimitRule{RandomDelay: 1 * time.Second})

	g.Colly.Visit(StartGV)
	g.Colly.Wait()


}
func (g *GOV) LoggingDocuments(log *api.Logging) error {

	if g.News == nil {
		return errors.New("error to search data in G1")

	}
	space := regexp.MustCompile(`\s+`)

	for index, news := range g.News {
		s := space.ReplaceAllString(news.Body, " ")
		log.WithFields(logrus.Fields{
			"Url":      news.Url,
			"Date":     news.Date,
			"Title":    strings.ToLower(news.Title),
			"SubTitle": news.Subtitle,
			"Body":     s,
		}).Info("News related:", index)

	}
	return nil
}

//NewG1 return crawler  instance of colly
func NewGov() Crawler {
	return &GOV{
		Colly: colly.NewCollector(colly.AllowedDomains(GV), colly.URLFilters(
			regexp.MustCompile(FilterGV),
		)),
		News:      []RelatedNews{},
		validator: ctx.NewValidate(),
	}
}
