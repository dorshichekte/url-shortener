package worker

import (
	"sync"

	"url-shortener/internal/app/repository/model"
	"url-shortener/internal/app/repository/postgres"
)

type Worker struct {
	Store    postgres.Repositories
	wg       sync.WaitGroup
	resultCh chan model.DeleteEvent
}
