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
	"strings"
	"time"
)

var (
	UOLstop  = false
	hasTitle = false
)

type Uol struct {
	Colly     *colly.Collector
	News      []RelatedNews
	validator *ctx.Validation
	Log       *api.Logging
}

//LoadNews returns related crawler by an entry
func (u *Uol) TrackNewsBasedOnCovid19() {
	detailsNews := RelatedNews{}

	u.Colly.OnHTML(".flex-wrap a[href]", func(e *colly.HTMLElement) {
		//if !UOLstop {
		e.Request.Visit(e.Attr("href"))
		//}

	})

	u.Colly.OnHTML("head", func(e *colly.HTMLElement) {

		e.ForEach("meta", func(_ int, el *colly.HTMLElement) {
			switch el.Attr("property") {
			case "og:title":
				for _, related := range nl {
					if strings.Contains(strings.ToLower(el.Attr("content")), related) {
						detailsNews.Title = el.Attr("content")
						continue
					}
					hasTitle = true

				}
			case "og:description":
				if hasTitle {
					detailsNews.Subtitle = el.Attr("content")
				}
			}

		})

	})
	u.Colly.OnHTML(".row", func(e *colly.HTMLElement) {
		e.ForEach(".text", func(_ int, el *colly.HTMLElement) {
			text := el.Text
			detailsNews.Body += text

		})

	})
	u.Colly.OnHTML("time", func(e *colly.HTMLElement) {

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
		hasNotice, err := u.validator.ValidateStruct(detailsNews)

		if err != nil {
			detailsNews = RelatedNews{}
			return
		}

		for _, re := range u.News {
			if re.Title == detailsNews.Title {
				hasNotice = false
				break

			}
		}
		if hasNotice {
			detailsNews.Url = e.Request.AbsoluteURL(e.Request.URL.Path)
			u.News = append(u.News, detailsNews)
		}

		//if len(u.News) == 20 {
		//	UOLstop = true
		//	return
		//}
		detailsNews = RelatedNews{}

	})

	u.Colly.Limit(&colly.LimitRule{RandomDelay: 1 * time.Second})

	u.Colly.Visit(StartUol)
	u.Colly.Wait()
	fmt.Println("total de noticia g1", len(u.News))

}
func (u *Uol) LoggingDocuments(log *api.Logging) error {

	if u.News == nil {
		return errors.New("error to search data in Uol")

	}
	reqBody := external.ReqDocuments{}

	for _, related := range u.News {
		result := fmt.Sprintf("%s %s",related.Body,related.Title)
		reqBody.Text = append(reqBody.Text, result)
	}


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

	for index, news := range u.News {
		log.WithFields(logrus.Fields{
			"Url":      news.Url,
			"Date":     news.Date,
			"Title":    strings.ToLower(news.Title),
			"SubTitle": news.Subtitle,
			"Body":     docs.Text[index],
			"From":     UolNews,
		}).Info()

	}
	return nil
}

//NewFatoOuFake return crawler  instance of colly
func NewUol() Crawler {
	return &Uol{
		Colly:     colly.NewCollector(colly.AllowedDomains(UolNews)),
		News:      []RelatedNews{},
		validator: ctx.NewValidate(),
	}
}
