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

var (
	BBCstop = false
)

type BBCNews struct {
	Colly     *colly.Collector
	News      []RelatedNews
	validator *ctx.Validation
	Log       *api.Logging
}

//LoadNews returns related crawler by an entry
func (b *BBCNews) TrackNewsBasedOnCovid19() {
	detailsNews := RelatedNews{}

	b.Colly.OnHTML(".lx-stream a[href]", func(e *colly.HTMLElement) {
		if !BBCstop {
			e.Request.Visit(e.Attr("href"))
		}

	})

	b.Colly.OnHTML("head", func(e *colly.HTMLElement) {

		e.ForEach("meta", func(_ int, el *colly.HTMLElement) {
			switch el.Attr("name") {
			case "twitter:title":
				detailsNews.Title = el.Attr("content")
			case "twitter:description":
				detailsNews.Subtitle = el.Attr("content")
			}
		})

	})
	b.Colly.OnHTML("main", func(e *colly.HTMLElement) {
		e.ForEach(".GridItemConstrainedMedium-sc-12lwanc-2 .Paragraph-k859h4-0", func(_ int, el *colly.HTMLElement) {
			text := el.Text
			detailsNews.Body += text

		})

	})
	b.Colly.OnHTML("time", func(e *colly.HTMLElement) {

		//time, err := time.Parse(time.RFC3339, e.Attr("datetime"))
		//if err != nil {
		//	logrus.WithFields(logrus.Fields{
		//		"Date": time,
		//	}).Info("Was not possible convert string to Date")
		//
		//}
		var data string
		data = e.Text

		if detailsNews.Date == strEmpty {
			detailsNews.Date = data
		}
		hasNotice, err := b.validator.ValidateStruct(detailsNews)

		if err != nil {
			detailsNews = RelatedNews{}
			return
		}

		for _, re := range b.News {
			if re.Title == detailsNews.Title {
				hasNotice = false
				break

			}
		}
		if hasNotice {
			detailsNews.Url = e.Request.AbsoluteURL(e.Request.URL.Path)
			b.News = append(b.News, detailsNews)
		}

		if len(b.News) == 2 {
			BBCstop = true
			return
		}
		detailsNews = RelatedNews{}

	})

	b.Colly.Limit(&colly.LimitRule{RandomDelay: 1 * time.Second})

	b.Colly.Visit(StartBBCNews)
	b.Colly.Wait()

}
func (b *BBCNews) LoggingDocuments(log *api.Logging) error {

	if b.News == nil {
		return errors.New("error to search data in BBCNews")

	}
	space := regexp.MustCompile(`\s+`)

	for _, news := range b.News {
		s := space.ReplaceAllString(news.Body, " ")
		log.WithFields(logrus.Fields{
			"Url":      news.Url,
			"Date":     news.Date,
			"Title":    strings.ToLower(news.Title),
			"SubTitle": news.Subtitle,
			"Body":     s,
			"From":     BBC,
		}).Info()

	}
	return nil
}

//NewFatoOuFake return crawler  instance of colly
func NewBBCNews() Crawler {
	return &G1{
		Colly: colly.NewCollector(colly.AllowedDomains(BBC), colly.URLFilters(
			regexp.MustCompile(FilterBBC),
		)),
		News:      []RelatedNews{},
		validator: ctx.NewValidate(),
	}
}
