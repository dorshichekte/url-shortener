package worker

import (
	"context"
	"url-shortener/internal/app/model"
)

func NewService(workerCount int, chLength int) *Service {
	w := &Service{
		resultCh: make(chan model.DeleteEvent, chLength),
	}

	for i := 0; i < workerCount; i++ {
		w.wg.Add(1)
		go w.Delete()
	}

	return w
}

func (w *Service) SendEvent(ctx context.Context, event model.DeleteEvent) {
	select {
	case w.resultCh <- event:
	case <-ctx.Done():
		return
	}
}

func (w *Service) Delete() {
	defer w.wg.Done()
	for event := range w.resultCh {
		w.Store.BatchUpdate(event)
	}
}

func (w *Service) Close() {
	close(w.resultCh)
	w.wg.Wait()
}
