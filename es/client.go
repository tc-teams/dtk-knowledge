package es

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"time"
)

type Elastic struct {
	client *elastic.Client
}

// Hit is a structure used for serializing/deserializing data in Elasticsearch.
type Hit struct {
	Link string
	Name string
	Date time.Time
}

//NewClient return a news instance client
func NewClient(url string) (*Elastic, error) {

	client, err := elastic.NewClient(elastic.SetURL(url))
	if err != nil {
		return nil, err
	}
	return &Elastic{client}, nil
}

func search() {

}

//TODO remove this method when create a mapping using kibana
func (e *Elastic) AddIndex(ctx context.Context, index string, body string) (string, error) {

	exists, err := e.client.IndexExists(index).Do(ctx)
	if err != nil {
		return "", err
	}
	if !exists {
		createIndex, err := e.client.CreateIndex(index).BodyString(body).Do(ctx)
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
	put, err := e.client.Index().
		Index(index).
		BodyJson(body).
		Do(ctx)

	if err != nil {
		return "", nil
	}
	fmt.Printf("Indexed tweet %s to index %s, type %s\n", put.Id, put.Index, put.Type)

	return put.Id, nil
}
