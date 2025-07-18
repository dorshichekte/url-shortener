package worker

import (
	"context"

	"url-shortener/internal/app/config/worker"
	entity "url-shortener/internal/app/domain/entity/url"
)

func New(config *config.Worker) *Worker {
	w := &Worker{
		resultCh: make(chan entity.DeleteBatch, config.ChanelLength),
	}

	w.wg.Add(config.WorkerCounter)
	for i := 0; i < config.WorkerCounter; i++ {
		go w.RunJob()
	}

	return w
}

func (w *Worker) SendEvent(ctx context.Context, event entity.DeleteBatch) {
	select {
	case w.resultCh <- event:
	case <-ctx.Done():
		return
	}
}

func (w *Worker) RunJob() {
	defer w.wg.Done()
	for event := range w.resultCh {
		_ = w.Store.Url.DeleteBatch(event)
	}
}

func (w *Worker) StopJob() {
	close(w.resultCh)
	w.wg.Wait()
}
