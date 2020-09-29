package crawler

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/gocolly/colly"
	"github.com/sirupsen/logrus"
	"github.com/tc-teams/fakefinder-crawler/api"
	ctx "github.com/tc-teams/fakefinder-crawler/context/validator"
	"github.com/tc-teams/fakefinder-crawler/external"
	"regexp"
	"strings"
	"time"
)

var (
	G1stop = false
)

type G1 struct {
	Colly     *colly.Collector
	News      []RelatedNews
	validator *ctx.Validation
	Log       *api.Logging
}

//LoadNews returns related crawler by an entry
func (g *G1) TrackNewsBasedOnCovid19() {
	detailsNews := RelatedNews{}

	g.Colly.OnHTML("#bstn-launcher a[href]", func(e *colly.HTMLElement) {
		if !G1stop {
			e.Request.Visit(e.Attr("href"))
		}

	})

	g.Colly.OnHTML("head", func(e *colly.HTMLElement) {

		e.ForEach("meta", func(_ int, el *colly.HTMLElement) {
			switch el.Attr("property") {
			case "og:title":
				detailsNews.Title = el.Attr("content")
			case "og:description":
				detailsNews.Subtitle = el.Attr("content")
			}
		})

	})
	g.Colly.OnHTML("main", func(e *colly.HTMLElement) {
		e.ForEach(".content-text__container", func(_ int, el *colly.HTMLElement) {
			text := el.Text
			detailsNews.Body += text

		})

	})
	g.Colly.OnHTML("time", func(e *colly.HTMLElement) {

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
			G1stop = true
			return
		}
		detailsNews = RelatedNews{}

	})

	g.Colly.Limit(&colly.LimitRule{RandomDelay: 1 * time.Second})

	g.Colly.Visit(StartG1)
	g.Colly.Wait()

}
func (g *G1) LoggingDocuments(log *api.Logging) error {

	if g.News == nil {
		return errors.New("error to search data in G1")

	}
	reqBody := external.ReqDocuments{}

	for _, related := range g.News {
		result := fmt.Sprintf("%s %s",related.Body,related.Title)
		reqBody.Text = append(reqBody.Text, result)
	}
	fmt.Println("req:",reqBody)


	req, err := external.NewClient().Request(reqBody)
	if err != nil{
		return err
	}

	var docs external.RespDocuments
	defer req.Body.Close()

	err = json.NewDecoder(req.Body).Decode(&docs)
	if err != nil {
		return err
	}

	for index , news := range g.News {
			log.WithFields(logrus.Fields{
				"Url":      news.Url,
				"Date":     news.Date,
				"Title":    strings.ToLower(news.Title),
				"SubTitle": news.Subtitle,
				"Body":     docs.Text[index],
				"From":     GB,
			}).Info()
		}

	return nil
}

//NewG1 return crawler  instance of colly
func NewG1() Crawler {
	return &G1{
		Colly: colly.NewCollector(colly.AllowedDomains(GB), colly.URLFilters(
			regexp.MustCompile(FilterGB),
		)),
		News:      []RelatedNews{},
		validator: ctx.NewValidate(),
	}
}
