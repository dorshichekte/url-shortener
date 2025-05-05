package memory

import (
	"url-shortener/internal/app/common"
)

type URLType string

type MapURL map[string]string

type Storage struct {
	mapURL MapURL
	common.BaseStorageDependency
}
