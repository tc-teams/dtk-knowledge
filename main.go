package main

import (
	"github.com/gocolly/colly/v2"
	"github.com/idasilva/dtk-knowledge/app"
)

func main() {

	c := app.NewColly(colly.NewCollector())
	c.LoadNews()

}
