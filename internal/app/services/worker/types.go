package worker

import (
	"sync"
	"url-shortener/internal/app/common"
	"url-shortener/internal/app/storage"

	"url-shortener/internal/app/models"
)

type Service struct {
	Store storage.URLStorage
	common.BaseDependency
	wg       sync.WaitGroup
	resultCh chan models.DeleteEvent
}
