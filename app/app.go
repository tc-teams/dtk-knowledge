package app

import (
	"github.com/tc-teams/fakefinder-crawler/api"
	"github.com/tc-teams/fakefinder-crawler/api/middlewares"
	"github.com/tc-teams/fakefinder-crawler/app/crawler"
)

func Run() *api.API {
	a := api.NewContextApi()

	a.InitRoute(crawler.Init())
	a.Middleware.Chain(middlewares.HelloWord(), middlewares.LogTime())
	return a
}
