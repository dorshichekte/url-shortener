package worker

import (
	"sync"

	"url-shortener/internal/app/common"
	"url-shortener/internal/app/model"
	"url-shortener/internal/app/storage"
)

type Service struct {
	Store storage.URLStorage
	common.BaseDependency
	wg       sync.WaitGroup
	resultCh chan model.DeleteEvent
}
