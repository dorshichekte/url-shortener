package urlhandler

import (
	"context"
	"net/http"

	"go.uber.org/zap"

	errorshandler "url-shortener/internal/app/adapter/primary/http/handler/errors"
	"url-shortener/internal/app/adapter/primary/http/middleware"
	"url-shortener/internal/pkg/constants"
)

// AddShorten godoc
//
//	@Summary		Создать короткую ссылку для одной URL
//	@Description	Принимает оригинальный URL в теле запроса (plain text), возвращает короткую ссылку
//	@Accept			plain
//	@Produce		plain
//	@Param			data	body		string	true	"Оригинальный URL"
//	@Tags			URL
//	@Success		201	{string}	string	"Короткая ссылка"
//	@Failure		400	{string}	string	"Ошибка в запросе"
//	@Failure		401	{string}	string	"Пользователь не авторизован"
//	@Failure		409	{string}	string	"Короткая ссылка уже существует"
//	@Failure		500	{string}	string	"Внутренняя ошибка сервера"
//	@Router			/ [post]
func (h *Handler) AddShorten(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), constants.DefaultTimeRequest)
	defer cancel()

	userID, ok := req.Context().Value(middleware.UserIDKey).(string)
	if userID == "" && !ok {
		h.logger.Error(errMessageFailedGetUserIDFromContext)
		h.handleError(res, http.StatusUnauthorized)
		return
	}

	originalURL, err := h.parseRequest(req)
	if err != nil {
		h.logger.Error(errorshandler.ErrMessageFailedParseRequestURI, zap.Error(err))
		h.handleError(res, http.StatusBadRequest)
		return
	}
	defer func() {
		_ = req.Body.Close()
	}()

	//ToDO переделать логику
	shortURL, err := h.useCase.AddShorten(ctx, originalURL, userID)
	baseURL := h.config.BaseURL
	fullURL := baseURL + "/" + shortURL
	if err != nil {
		h.logger.Error(errMessageFailedCreateShortURL, zap.Error(err))
		res.Header().Set("Content-Type", "text/plain")
		h.handleError(res, http.StatusConflict)

		_, jsonWriteErr := res.Write([]byte(fullURL))
		if jsonWriteErr != nil {
			h.logger.Error(errorshandler.ErrMessageFailedWriteResponse, zap.Error(err))
			res.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	res.Header().Set("Content-Type", "text/plain")
	res.WriteHeader(http.StatusCreated)
	_, jsonWriteErr := res.Write([]byte(fullURL))
	if jsonWriteErr != nil {
		h.logger.Error(errorshandler.ErrMessageFailedWriteResponse, zap.Error(err))
		res.WriteHeader(http.StatusInternalServerError)
		return
	}
}
