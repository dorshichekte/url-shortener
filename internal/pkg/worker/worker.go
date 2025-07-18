package worker

import (
	"context"

	"url-shortener/internal/app/config/worker"
	"url-shortener/internal/app/repository/model"
)

func New(config *config.Worker) *Worker {
	w := &Worker{
		resultCh: make(chan model.DeleteEvent, config.ChanelLength),
	}

	w.wg.Add(config.WorkerCounter)
	for i := 0; i < config.WorkerCounter; i++ {
		go w.RunJob()
	}

	return w
}

func (w *Worker) SendEvent(ctx context.Context, event model.DeleteEvent) {
	select {
	case w.resultCh <- event:
	case <-ctx.Done():
		return
	}
}

func (w *Worker) RunJob() {
	defer w.wg.Done()
	for event := range w.resultCh {
		w.Store.Url.DeleteBatch(event)
	}
}

func (w *Worker) StopJob() {
	close(w.resultCh)
	w.wg.Wait()
}
