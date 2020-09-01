package elastic

import (
	es "github.com/tc-teams/fakefinder-crawler/elastic/es"
	"github.com/tc-teams/fakefinder-crawler/tracker/crawl"
)

var url = "http://localhost:5061"

func ElasticDocumentsByDescription() ([]*crawl.RelatedNews, error) {

	es, err := es.NewInstanceElastic(url)
	if err != nil {
		return nil, err
	}
	es.Search()
	return nil, nil

}
