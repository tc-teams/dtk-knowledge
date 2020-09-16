package es

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/olivere/elastic/v7"
	log "github.com/sirupsen/logrus"
	"strings"
)

type Elastic struct {
	*elastic.Client
	ctx context.Context
}

//Search
func (e *Elastic) MatchQueryByIndex(description string) ([]Data, error) {

	var source []Data

	searchSource := elastic.NewSearchSource()
	searchSource.Query(elastic.NewMatchQuery(Fields, description))

	_, err := searchSource.Source()
	if err != nil {
		return nil, err
	}

	searchService := e.Search().
		Index(Index).
		SearchSource(searchSource)

	searchResult, err := searchService.Do(e.ctx)
	if err != nil {
		return nil, err
	}
	log.WithFields(log.Fields{
		"hits:": searchResult.Hits.TotalHits,
	}).Info()

	for _, hit := range searchResult.Hits.Hits {
		var data Data
		err := json.Unmarshal(hit.Source, &data)
		if err != nil {
			return nil, err
		}
		if strings.Index(data.News.Title, description) > -1 {
			source = append(source, data)
		}
	}

	if err != nil {
		return nil, err

	}

	return source, nil
}

func (e *Elastic) EspecifiedValues(value string) error {
	get1, err := e.Get().
		Index(Index).
		Id(value).
		Do(e.ctx)
	if err != nil {
		return err
	}
	if get1.Found {
		var hits Data
		err := json.Unmarshal(get1.Source, &hits)
		if err != nil {
			return err
		}
		fmt.Printf("Got document %s", hits)
	}
	return nil
}

func (e *Elastic) Version(url string) (bool, error) {

	esversion, err := e.Client.ElasticsearchVersion(url)
	if err != nil {
		return false, err

	}
	fmt.Printf("Elasticsearch version %s\n", esversion)
	return true, nil
}

//TODO remove this method when create a mapping using kibana
func (e *Elastic) AddIndex(ctx context.Context, index string, body string) (string, error) {

	exists, err := e.Client.IndexExists(index).Do(ctx)
	if err != nil {
		return "", err
	}
	if !exists {
		createIndex, err := e.Client.CreateIndex(index).BodyString(body).Do(ctx)
		if err != nil {
			return "", err
		}
		if !createIndex.Acknowledged {
			// Not acknowledged
		}
	}
	return "", nil

}
func (e *Elastic) Index(ctx context.Context, index string, body interface{}) (string, error) {
	put, err := e.Client.Index().
		Index(index).
		BodyJson(body).
		Do(ctx)

	if err != nil {
		return "", nil
	}
	fmt.Printf("index %s, type %s\n", put.Index, put.Type)

	return put.Id, nil
}

//NewInstanceElastic return a crawler instance client
func NewInstanceElastic(url string, user string, password string) (*Elastic, error) {

	client, err := elastic.NewClient(elastic.SetURL(url), elastic.SetBasicAuth(user, password))
	if err != nil {
		fmt.Println("erro na autenticação", err)
		return nil, err
	}
	return &Elastic{client, context.Background()}, nil
}
