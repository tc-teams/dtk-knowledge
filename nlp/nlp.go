package nlp

import (
	"github.com/tc-teams/fakefinder-crawler/elastic/es"
	"github.com/tc-teams/fakefinder-crawler/external"

	"sort"
)

var (
	equals = false
)

func NaturalLanguageProcess(pln external.PlnResponse, documents []es.Data) external.BotResponse {

	var (
		values  = map[float64]string{}
		keys    = []float64{}
		ordered = map[string]float64{}
	)

	for key, val := range pln.PlnProcess {
		values[val] = key
		keys = append(keys, val)
	}
	sort.Sort(sort.Reverse(sort.Float64Slice(keys)))

	for index, val := range keys {
		if index == 4 {
			break
		}
		ordered[values[val]] = val
	}

	var bot external.BotResponse
	bot.Description = pln.Description
	for _, j := range documents {
		var text external.TextResult
		if value := ordered[j.News.Body]; value != 0.0 {
			if len(bot.Text) == 0 {
				text.Similarity = value
				text.Link = j.News.Url
				text.Title = j.News.Title
				text.Date = j.News.Date
				bot.Text = append(bot.Text, text)
				continue
			} else {
				for _, iqls := range bot.Text {

					if iqls.Title == j.News.Title {
						equals = true
						break
					}

				}

			}

			if !equals {
				text.Similarity = value
				text.Link = j.News.Url
				text.Title = j.News.Title
				text.Date = j.News.Date
				bot.Text = append(bot.Text, text)

			}
			equals = false

		}

	}

	return bot

}
