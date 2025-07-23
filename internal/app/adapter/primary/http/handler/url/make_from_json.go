package urlhandler

import (
	"context"
	"encoding/json"
	"net/http"

	"go.uber.org/zap"

	dto "url-shortener/internal/app/adapter/primary/http/dto/url"
	errorshandler "url-shortener/internal/app/adapter/primary/http/handler/errors"
	"url-shortener/internal/app/adapter/primary/http/middleware"
	"url-shortener/internal/pkg/constants"
)

// MakeFromJSON godoc
//
//	@Summary		Create single short url
//	@Description	MakeFromJSON is used to handle single url in request
//	@Security		ApiKeyAuth
//	@Accept			json
//	@Produce		json
//	@Tags			API
//	@Param			data	body		dto.ShortenRequest	true	"Request body"
//	@Success		201 {object}	dto.ShortenResponse
//	@Failure		400,401,409,500
//	@Router			/api/shorten   [post]
func (h *Handler) MakeFromJSON(res http.ResponseWriter, req *http.Request) {
	ctx, cancel := context.WithTimeout(req.Context(), constants.DefaultTimeRequest)
	defer cancel()

	userID, ok := req.Context().Value(middleware.UserIDKey).(string)
	if userID == "" && !ok {
		h.logger.Error(errMessageFailedGetUserIDFromContext)
		h.handleError(res, http.StatusUnauthorized)
		return
	}

	//ToDO переписать
	u, err := h.jsonDecode(req)
	if err != nil {
		h.logger.Error(errorshandler.ErrMessageFailedDecodeJSON, zap.Error(err))
		h.handleError(res, http.StatusInternalServerError)
		return
	}

	//ToDo поправить логику
	shortURL, err := h.useCase.AddShorten(ctx, u.OriginalURL, userID)
	baseURL := h.config.BaseURL
	fullURL := baseURL + "/" + shortURL
	response := dto.ShortenResponse{
		ShortURL: fullURL,
	}
	if err != nil {
		h.logger.Error(errMessageFailedCreateShortURL, zap.Error(err))
		res.Header().Set("Content-Type", "application/json")
		h.handleError(res, http.StatusConflict)
		err = json.NewEncoder(res).Encode(response)
		if err != nil {
			h.logger.Error(errorshandler.ErrMessageFailedWriteResponse, zap.Error(err))
			res.WriteHeader(http.StatusInternalServerError)
		}
		return
	}

	res.Header().Set("Content-Type", "application/json")
	res.WriteHeader(http.StatusCreated)
	err = json.NewEncoder(res).Encode(response)
	if err != nil {
		h.logger.Error(errorshandler.ErrMessageFailedWriteResponse, zap.Error(err))
		res.WriteHeader(http.StatusInternalServerError)
	}
}
