package elastic

import (
	"github.com/spf13/viper"
	"github.com/tc-teams/fakefinder-crawler/api"
	"github.com/tc-teams/fakefinder-crawler/elastic/es"
)

//DocumentsByDescription return
func DocumentsByDescription(log *api.Logging, description string) ([]es.Data, error) {
	Url := viper.GetString("Url")
	User := viper.GetString("User")
	Password := viper.GetString("Password")

	es, err := es.NewInstanceElastic(Url, User, Password)
	if err != nil {
		return nil, err
	}
	log.Println("New Instance of elastic search created")

	source, err := es.MatchQueryByIndex(description)
	if err != nil {
		return nil, err
	}

	return source, nil

}
