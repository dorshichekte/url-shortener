package urlhandler

import (
	"encoding/json"
	"io"
	"net/http"

	"go.uber.org/zap"

	errorshandler "url-shortener/internal/app/adapter/primary/http/handler/errors"
	"url-shortener/internal/app/adapter/primary/http/middleware"
	entity "url-shortener/internal/app/domain/entity/url"
)

// DeleteBatch godoc
// @Summary      Пакетное удаление URL пользователя
// @Description  Помечает указанные URL как удаленные (устанавливает флаг is_deleted=true) для конкретного пользователя.
//
//	Операция выполняется асинхронно, сразу возвращает статус Accepted (202).
//	Требуется аутентификация по API-ключу.
//
// @Security     ApiKeyAuth
// @Accept       json
// @Produce      json
// @Tags         Управление URL
// @Param        request body []string true "Список сокращенных URL для удаления"
//
//	example: ["abc123", "def456"]
//
// @Success      202 "Запрос на удаление принят в обработку"
// @Failure      400
// @Failure      401
// @Failure      500
// @Router       /api/user/urls [delete]
func (h *Handler) DeleteBatch(res http.ResponseWriter, req *http.Request) {
	userID, ok := req.Context().Value(middleware.UserIDKey).(string)
	if userID == "" && !ok {
		h.logger.Error(errMessageFailedGetUserIDFromContext)
		h.handleError(res, http.StatusUnauthorized)
		return
	}

	body, err := io.ReadAll(req.Body)
	if err != nil || len(body) == 0 {
		h.logger.Error(errorshandler.ErrMessageFailedParseRequestBody, zap.Error(err))
		h.handleError(res, http.StatusBadRequest)
		return
	}

	var listURLs []string
	err = json.Unmarshal(body, &listURLs)
	if err != nil {
		h.logger.Error(errorshandler.ErrMessageFailedUnmarshalJSON, zap.Error(err))
		h.handleError(res, http.StatusBadRequest)
		return
	}

	event := entity.DeleteBatch{
		ListURL: listURLs,
		UserID:  userID,
	}

	go h.useCase.DeleteBatch(event)
	res.WriteHeader(http.StatusAccepted)
}
