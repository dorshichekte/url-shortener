package worker

import (
	"context"

	"url-shortener/internal/app/repository/model"
)

func New(workerCount int, chLength int) *Worker {
	w := &Worker{
		resultCh: make(chan model.DeleteEvent, chLength),
	}

	for i := 0; i < workerCount; i++ {
		w.wg.Add(1)
		go w.Delete()
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

func (w *Worker) Delete() {
	defer w.wg.Done()
	for event := range w.resultCh {
		w.Store.BatchUpdate(event)
	}
}

func (w *Worker) Close() {
	close(w.resultCh)
	w.wg.Wait()
}
