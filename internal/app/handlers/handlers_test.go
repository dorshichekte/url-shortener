package handlers

import (
	"io"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"url-shortener/internal/app/config"
	"url-shortener/internal/app/services/url"
)

func testRequest(t *testing.T, ts *httptest.Server, method,
	path, body string) (*http.Response, string) {

	client := &http.Client{
		CheckRedirect: func(req *http.Request, via []*http.Request) error {
			return http.ErrUseLastResponse
		},
	}

	req, err := http.NewRequest(method, ts.URL+path, strings.NewReader(body))
	require.NoError(t, err)

	resp, err := client.Do(req)
	require.NoError(t, err)
	defer resp.Body.Close()

	respBody, err := io.ReadAll(resp.Body)
	require.NoError(t, err)

	return resp, string(respBody)
}

func TestRoute(t *testing.T) {
	mockURL := "https://ya.ru"
	mockTestData := url.CreateShort(mockURL)
	ts := httptest.NewServer(Register())
	baseURL := config.Get().BaseURL

	defer ts.Close()

	type values struct {
		url    string
		method string
		body   string
	}

	type want struct {
		status int
		body   string
	}

	var tests = []struct {
		name   string
		values values
		want   want
	}{
		{
			name: "Test #1 Запись в хранилище",
			values: values{
				url:    "/",
				method: "POST",
				body:   mockURL,
			},
			want: want{
				status: http.StatusCreated,
			},
		},
		{
			name: "Test #2 Метод GET вместо POST",
			values: values{
				url:    "/",
				method: "GET",
				body:   mockURL,
			},
			want: want{
				status: http.StatusMethodNotAllowed,
			},
		},
		{
			name: "Test #3 без тела запроса",
			values: values{
				url:    "/",
				method: "POST",
				body:   "",
			},
			want: want{
				status: http.StatusBadRequest,
			},
		},
		{
			name: "Test #4 не валидный URL",
			values: values{
				url:    "/",
				method: "POST",
				body:   "ggf.fdfhk/fsdf",
			},
			want: want{
				status: http.StatusBadRequest,
			},
		},
		{
			name: "Test #5 добавление дубликата",
			values: values{
				url:    "/",
				method: "POST",
				body:   mockURL,
			},
			want: want{
				status: http.StatusCreated,
				body:   baseURL + "/" + mockTestData,
			},
		},
		{
			name: "Test #6 проверка извлечения URL по сокращенной ссылке",
			values: values{
				url:    "/" + mockTestData,
				method: "GET",
				body:   "",
			},
			want: want{
				status: http.StatusTemporaryRedirect,
				body:   mockURL,
			},
		},
		{
			name: "Test #7 метод POST вмсето GET",
			values: values{
				url:    "/" + mockTestData,
				method: "POST",
				body:   "",
			},
			want: want{
				status: http.StatusMethodNotAllowed,
			},
		},
		{
			name: "Test #8 несуществующая сокращенная ссылка",
			values: values{
				url:    "/sdfjvu88934nkdkl",
				method: "GET",
				body:   "",
			},
			want: want{
				status: http.StatusBadRequest,
			},
		},
	}

	for num, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			resp, URL := testRequest(t, ts, test.values.method, test.values.url, test.values.body)
			defer resp.Body.Close()

			assert.Equal(t, test.want.status, resp.StatusCode)

			if num == 4 {
				assert.Equal(t, test.want.body, URL)
			} else {
				assert.Equal(t, test.want.body, resp.Header.Get("Location"))
			}
		})
	}
}
