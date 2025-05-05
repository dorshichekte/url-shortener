package worker

import (
	"context"
	"url-shortener/internal/app/models"
)

func New(workerCount int, chLength int) *Service {
	w := &Service{
		resultCh: make(chan models.DeleteEvent, chLength),
	}

	for i := 0; i < workerCount; i++ {
		w.wg.Add(1)
		//go w.RunJobDeleteBatch()
	}

	return w
}

func (w *Service) SendDeleteBatchRequest(ctx context.Context, event models.DeleteEvent) {
	select {
	case w.resultCh <- event:
	case <-ctx.Done():
		return
	}
}

//func (w *Service) RunJobDeleteBatch() {
//	defer w.wg.Done()
//	for event := range w.resultCh {
//		w.BaseService.Store.BatchUpdate(event)
//	}
//}

func (w *Service) Close() {
	close(w.resultCh)
	w.wg.Wait()
}
