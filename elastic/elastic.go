package elastic

import (
	"encoding/json"
	"github.com/tc-teams/fakefinder-crawler/api"
	"github.com/tc-teams/fakefinder-crawler/elastic/es"
	"github.com/tc-teams/fakefinder-crawler/external"
)

//DocumentsByDescription return
func DocumentsByDescription(log *api.Logging, description string) ([]es.Data, error) {
	//Url := viper.GetString("Url")
	//User := viper.GetString("User")
	//Password := viper.GetString("Password")

	es, err := es.NewInstanceElastic("http://elasticsearch:9200", "elastic", "changeme")
	if err != nil {
		return nil, err
	}
	log.Println("New Instance of elastic search created")


	reqBody := external.ReqDocuments{}
	reqBody.Text[0] = description


	req, err := external.NewClient().Request(reqBody)
	if err != nil{
		return nil, err
	}

	var docs external.RespDocuments

	err = json.NewDecoder(req.Body).Decode(&docs)
	if err != nil {
		return nil,err
	}


	source, err := es.MatchQueryByIndex(docs.Text[0])
	if err != nil {
		return nil, err
	}

	return source, nil

}
