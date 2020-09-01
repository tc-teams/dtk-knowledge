package es

import (
	"context"
	"fmt"
	"github.com/olivere/elastic"
	"time"
)
//https://olivere.github.io/elastic/
type Elastic struct {
	client *elastic.Client
	ctx    context.Context
}

// Hit is a structure used for serializing/deserializing data in Elasticsearch.
type Hit struct {
	Url      string
	Date     time.Time
	Title    string
	Subtitle string
	Body     string
}
//Search
func (e *Elastic) Search() {

	result, err := e.client.Get().
		Index("twitter").
		Type("tweet").
		Id("1").
		Do(e.ctx)
	if err != nil {
		panic(err)
	}
	if result.Found {
		fmt.Printf("Got document %s in version %d from index %s, type %s\n", result.Id, result.Version, result.Index, result.Type)
	}

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

//NewInstanceElastic return a crawler instance client
func NewInstanceElastic(url string) (*Elastic, error) {

	client, err := elastic.NewClient(elastic.SetURL(url))
	if err != nil {
		return nil, err
	}
	return &Elastic{client, context.Background()}, nil
}
