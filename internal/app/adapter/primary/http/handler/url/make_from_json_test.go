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
	"url-shortener/internal/app/domain/mock"
	urlusecase "url-shortener/internal/app/usecase/url"
	"url-shortener/internal/pkg/validator"
)

func initBenchmarkDataHandlerMakeFromJSON(b *testing.B) *Handler {
	ctrl := gomock.NewController(b)
	defer ctrl.Finish()

	mockRepo := mock.NewMockIURLRepository(ctrl)

	mockRepo.
		EXPECT().
		AddShorten(gomock.Any(), gomock.Any(), gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, originalURL, shortURL, userID string) (string, error) {
			return "", nil
		}).
		AnyTimes()

	mockRepo.
		EXPECT().
		GetByOriginalURL(gomock.Any(), gomock.Any()).
		DoAndReturn(func(ctx context.Context, originalURL string) (string, error) {
			return "", nil
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

func BenchmarkHandlerMakeFromJSON(b *testing.B) {
	handler := initBenchmarkDataHandlerMakeFromJSON(b)

	testData := make([]dto.ShortenRequest, b.N)
	for i := 0; i < b.N; i++ {
		testData[i] = dto.ShortenRequest{
			OriginalURL: fmt.Sprintf("https://www.youtube.com/%d", i),
		}
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		body, err := json.Marshal(testData[i])
		if err != nil {
			b.Fatalf("failed to marshal request body: %v", err)
		}
		
		req := httptest.NewRequest(http.MethodPost, "/api/shorten", bytes.NewReader(body))
		req.Header.Set("Content-Type", "application/json")
		req = req.WithContext(context.WithValue(req.Context(), middleware.UserIDKey, "test-user"))

		recorder := httptest.NewRecorder()
		handler.MakeFromJSON(recorder, req)

		res := recorder.Result()
		_ = res.Body.Close()

		if res.StatusCode != http.StatusCreated {
			b.Fatalf("handler returned wrong status code: got %v want %v", res.StatusCode, http.StatusOK)
		}
	}
}
