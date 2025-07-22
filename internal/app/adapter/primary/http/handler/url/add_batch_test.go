package urlhandler

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/golang/mock/gomock"
	"go.uber.org/zap/zaptest"

	dto "url-shortener/internal/app/adapter/primary/http/dto/url"
	"url-shortener/internal/app/adapter/primary/http/middleware"
	env "url-shortener/internal/app/config/env"
	. "url-shortener/internal/app/domain/entity/url"
	"url-shortener/internal/app/domain/mock"
	urlusecase "url-shortener/internal/app/usecase/url"
	"url-shortener/internal/pkg/validator"
)

func initBenchmarkDataHandlerAddBatch(b *testing.B) *Handler {
	ctrl := gomock.NewController(b)
	defer ctrl.Finish()

	mockRepo := mock.NewMockIURLRepository(ctrl)

	mockRepo.
		EXPECT().
		AddBatch(gomock.Any(), gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, batches []Batch, userID string) error {
			for i := range batches {
				batches[i].ShortURL = "http://short/" + batches[i].ID
			}
			return nil
		}).
		AnyTimes()

	cfg := &env.Env{
		BaseURL: env.DefaultAddressWithProtocol,
	}

	urlUC := urlusecase.New(cfg, mockRepo)
	logger := zaptest.NewLogger(b)
	v := validator.New()

	handler := New(logger, cfg, urlUC, v)

	return handler
}

func BenchmarkHandlerAddBatch(b *testing.B) {
	handler := initBenchmarkDataHandlerAddBatch(b)

	testData := make([]dto.BatchRequest, b.N)
	for i := 0; i < b.N; i++ {
		testData[i] = dto.BatchRequest{
			ID:          fmt.Sprintf("%d", i),
			OriginalURL: fmt.Sprintf("https://www.youtube.com/%d", i),
		}
	}

	body, err := json.Marshal(testData)
	if err != nil {
		b.Fatalf("failed to marshal request body: %v", err)
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest(http.MethodPost, "/api/shorten/batch", bytes.NewReader(body))
		req = req.WithContext(context.WithValue(req.Context(), middleware.UserIDKey, "test-user"))

		rec := httptest.NewRecorder()
		handler.AddBatch(rec, req)

		res := rec.Result()
		_ = res.Body.Close()

		if res.StatusCode != http.StatusCreated {
			b.Fatalf("unexpected status: got %v, want %v", res.StatusCode, http.StatusCreated)
		}
	}
}
