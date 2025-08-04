package usecase

import (
	urlusecase "url-shortener/internal/app/usecase/url"
)

// UseCases агрегирует все бизнес-слои (use case) приложения.
type UseCases struct {
	URL urlusecase.IUrlUseCase
}
