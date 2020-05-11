package main

import (
	"github.com/gocolly/colly/v2"
	"github.com/idasilva/dtk-knowledge/app/collector"
	log "github.com/sirupsen/logrus"
	"os"
)

func main() {
	f, err := os.OpenFile("logrus.log", os.O_APPEND|os.O_CREATE|os.O_RDWR, 0666)
	if err != nil {
		panic(0)
	}

	defer f.Close()

	c := collector.NewColly(colly.NewCollector(),&log.Logger{
		Out:          f,
		Formatter:    &log.JSONFormatter{},
		Level:        log.DebugLevel,
	})

	c.LoadNews()

}
