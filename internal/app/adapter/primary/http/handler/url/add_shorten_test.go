package urlhandler

import (
	"context"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/golang/mock/gomock"
	"go.uber.org/zap/zaptest"

	"url-shortener/internal/app/adapter/primary/http/middleware"
	env "url-shortener/internal/app/config/env"
	"url-shortener/internal/app/domain/mock"
	urlusecase "url-shortener/internal/app/usecase/url"
	"url-shortener/internal/pkg/validator"
)

func initBenchmarkDataHandlerAddShorten(b *testing.B) *Handler {
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

func BenchmarkHandlerAddShorten(b *testing.B) {
	handler := initBenchmarkDataHandlerAddShorten(b)

	testData := make([]string, b.N)
	for i := 0; i < b.N; i++ {
		testData[i] = fmt.Sprintf("https://www.youtube.com/%d", i)
	}

	b.ReportAllocs()
	b.ResetTimer()

	for i := 0; i < b.N; i++ {
		req := httptest.NewRequest(http.MethodPost, "/", strings.NewReader(testData[i]))
		req.Header.Set("Content-Type", "text/plain")
		req = req.WithContext(context.WithValue(req.Context(), middleware.UserIDKey, "test-user"))

		recorder := httptest.NewRecorder()
		handler.AddShorten(recorder, req)

		res := recorder.Result()
		_ = res.Body.Close()

		if res.StatusCode != http.StatusCreated {
			b.Fatalf("handler returned wrong status code: got %v want %v", res.StatusCode, http.StatusOK)
		}
	}
}
