package worker

import (
	"sync"

	entity "url-shortener/internal/app/domain/entity/url"
	"url-shortener/internal/app/repository/postgres"
)

type Worker struct {
	Store    postgres.Repositories
	wg       sync.WaitGroup
	resultCh chan entity.DeleteBatch
}
