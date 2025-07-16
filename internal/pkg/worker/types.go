package worker

import (
	"sync"

	"url-shortener/internal/app/repository"
	"url-shortener/internal/app/repository/model"
)

type Worker struct {
	Store    repository.URLStorage
	wg       sync.WaitGroup
	resultCh chan model.DeleteEvent
}
