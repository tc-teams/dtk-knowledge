package main

import (
	"github.com/gocolly/colly/v2"
	"github.com/idasilva/dtk-knowledge/app/collector"
)

func main() {

	c := collector.NewColly(colly.NewCollector())
	c.LoadNews()

}
