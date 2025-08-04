package main

import (
	"context"
	"log"

	a "url-shortener/internal/app"
	c "url-shortener/internal/app/config"
	g "url-shortener/internal/pkg/graceful"
	l "url-shortener/internal/pkg/logger"
	w "url-shortener/internal/pkg/worker"
)

var (
	buildVersion string = "N/A"
	buildCommit  string = "N/A"
	buildDate    string = "N/A"
)

func main() {
	log.Printf("Build version: %s", buildVersion)
	log.Printf("Build date: %s", buildDate)
	log.Printf("Build commit: %s", buildCommit)

	logger, err := l.New()
	if err != nil {
		log.Fatal(err)
	}
	defer func() {
		_ = logger.Sync()
	}()

	config, err := c.New()
	if err != nil {
		logger.Fatal(err.Error())
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app := a.New(ctx, logger, config)

	graceful := g.New(g.NewProcess(app.HTTPAdapter))

	worker := w.New(ctx, config.Worker)
	defer worker.StopJob()

	err = graceful.Start(ctx)
	if err != nil {
		logger.Fatal(err.Error())
	}
}
