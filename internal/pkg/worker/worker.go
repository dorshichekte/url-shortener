package worker

import (
	"context"

	config "url-shortener/internal/app/config/worker"
	entity "url-shortener/internal/app/domain/entity/url"
)

// ToDo переписать полностью
// New создает новый экземпляр Worker с указанным количеством горутин.
func New(context context.Context, config *config.Worker) *Worker {
	w := &Worker{
		resultCh: make(chan entity.DeleteBatch, config.ChanelLength),
	}

	w.wg.Add(config.WorkerCounter)
	for i := 0; i < config.WorkerCounter; i++ {
		go w.RunJob(context)
	}

	return w
}

// SendEvent отправляет задание на удаление в очередь воркера.
func (w *Worker) SendEvent(ctx context.Context, event entity.DeleteBatch) {
	select {
	case w.resultCh <- event:
	case <-ctx.Done():
		return
	}
}

// RunJob основная рабочая функция воркера.
func (w *Worker) RunJob(context context.Context) {
	defer w.wg.Done()
	for event := range w.resultCh {
		_ = w.Store.URL.DeleteBatch(event)
	}
}

// StopJob корректно останавливает работу воркера.
func (w *Worker) StopJob() {
	close(w.resultCh)
	w.wg.Wait()
}
