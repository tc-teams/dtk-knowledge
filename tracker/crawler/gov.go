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
	"io/ioutil"
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

var (
	Govstop = false
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
		data := strings.Split(e.Text, ",")

		if detailsNews.Date == strEmpty {
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

		if len(g.News) == 6 {
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

	if len(g.News) == 0 {
		return errors.New("error to search data in GOV")

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
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		panic(err.Error())
	}
	fmt.Println(string(body))

	err = json.NewDecoder(req.Body).Decode(&docs)
	if err != nil {
		return err
	}


	for index, news := range g.News {
		log.WithFields(logrus.Fields{
			"Url":      news.Url,
			"Date":     news.Date,
			"Title":    strings.ToLower(news.Title),
			"SubTitle": news.Subtitle,
			"Body":     docs.Text[index],
			"From":     GV,
		}).Info()

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
