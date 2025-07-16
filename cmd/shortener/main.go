package main

import (
	"context"
	"log"

	a "url-shortener/internal/app"
	c "url-shortener/internal/app/config"
	g "url-shortener/internal/pkg/graceful"
	l "url-shortener/internal/pkg/logger"
)

func main() {
	logger, err := l.New()
	if err != nil {
		log.Fatal(err)
		return
	}
	defer func() {
		_ = logger.Sync()
	}()

	config, err := c.New()
	if err != nil {
		logger.Fatal(err.Error())
		return
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	app := a.New(ctx, logger, config)

	graceful := g.New(g.NewProcess(app.HTTPAdapter))

	err = graceful.Start(ctx)
	if err != nil {
		logger.Fatal(err.Error())
		panic(err)
	}
}
