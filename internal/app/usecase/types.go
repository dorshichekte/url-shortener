package usecase

import (
	dbusecase "url-shortener/internal/app/usecase/db"
	urlusecase "url-shortener/internal/app/usecase/url"
)

// UseCases агрегирует все бизнес-слои (use case) приложения.
type UseCases struct {
	URL urlusecase.IUrlUseCase
	DB  dbusecase.IDBUseCase
}
