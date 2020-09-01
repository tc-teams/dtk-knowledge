package app

import (
	"github.com/tc-teams/fakefinder-crawler/api"
	"github.com/tc-teams/fakefinder-crawler/api/middlewares"
	"github.com/tc-teams/fakefinder-crawler/app/crawler"
)

func Run() (*api.API, error) {
	app, err := api.NewContextApi()
	if err != nil{
		return nil, err
	}
	app.InitRoute(crawler.Init())

	app.Middleware.Chain(middlewares.HelloWord(), middlewares.LogTime())
	return app, nil
}
