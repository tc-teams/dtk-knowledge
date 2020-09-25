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
	FFstop = false
)

type FatoOuFake struct {
	Colly     *colly.Collector
	News      []RelatedNews
	validator *ctx.Validation
	Log       *api.Logging
}

//LoadNews returns related crawler by an entry
func (f *FatoOuFake) TrackNewsBasedOnCovid19() {
	detailsNews := RelatedNews{}

	f.Colly.OnHTML(".feed-root a[href]", func(e *colly.HTMLElement) {
		if !FFstop {
			e.Request.Visit(e.Attr("href"))
		}

	})

	f.Colly.OnHTML("head", func(e *colly.HTMLElement) {

		e.ForEach("meta", func(_ int, el *colly.HTMLElement) {
			switch el.Attr("property") {
			case "twitter:title":
				detailsNews.Title = el.Attr("content")
			case "twitter:description":
				detailsNews.Subtitle = el.Attr("content")
			}
		})

	})
	f.Colly.OnHTML("main", func(e *colly.HTMLElement) {
		e.ForEach(".content-text__container", func(_ int, el *colly.HTMLElement) {
			text := el.Text
			detailsNews.Body += text

		})

	})
	f.Colly.OnHTML("time", func(e *colly.HTMLElement) {

		//time, err := time.Parse(time.RFC3339, e.Attr("datetime"))
		//if err != nil {
		//	logrus.WithFields(logrus.Fields{
		//		"Date": time,
		//	}).Info("Was not possible convert string to Date")
		//
		//}
		var data string
		switch e.Attr("itemprop") {
		case "datePublished":
			data = e.Text
		}
		if detailsNews.Date == strEmpty {
			detailsNews.Date = data
		}
		hasNotice, err := f.validator.ValidateStruct(detailsNews)

		if err != nil {
			detailsNews = RelatedNews{}
			return
		}

		for _, re := range f.News {
			if re.Title == detailsNews.Title {
				hasNotice = false
				break

			}
		}
		if hasNotice {
			detailsNews.Url = e.Request.AbsoluteURL(e.Request.URL.Path)
			f.News = append(f.News, detailsNews)
		}

		if len(f.News) == 2 {
			FFstop = true
			return
		}
		detailsNews = RelatedNews{}

	})

	f.Colly.Limit(&colly.LimitRule{RandomDelay: 1 * time.Second})

	f.Colly.Visit(StartFatoOuFake)
	f.Colly.Wait()

}
func (f *FatoOuFake) LoggingDocuments(log *api.Logging) error {

	if f.News == nil {
		return errors.New("error to search data in Fato-ou-Fake")

	}
	space := regexp.MustCompile(`\s+`)

	for _, news := range f.News {
		s := space.ReplaceAllString(news.Body, " ")
		log.WithFields(logrus.Fields{
			"Url":      news.Url,
			"Date":     news.Date,
			"Title":    strings.ToLower(news.Title),
			"SubTitle": news.Subtitle,
			"Body":     s,
			"From":     GB,
		}).Info()

	}
	return nil
}

//NewFatoOuFake return crawler  instance of colly
func NewFatoOuFake() Crawler {
	return &G1{
		Colly: colly.NewCollector(colly.AllowedDomains(GB), colly.URLFilters(
			regexp.MustCompile(FilterFF),
		)),
		News:      []RelatedNews{},
		validator: ctx.NewValidate(),
	}
}
