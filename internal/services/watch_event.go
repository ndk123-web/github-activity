package services

import (
	"github.com/ndk123-web/github-activity/internal/models"
)

type WatchEventService interface {
	GetAllWatchEvent(limit int64) (map[string]int64, error)
}

type watchEventService struct {
	jsonData []models.GitResponseObject
}

func (w *watchEventService) GetAllWatchEvent(limit int64) (map[string]int64, error) {
	mapp := make(map[string]int64)
	for _, item := range w.jsonData {
		if item.Type != "WatchEvent" {
			continue
		}
		mapp[item.Repo.Name]++
		if limit > 0 && int64(len(mapp)) >= limit {
			break
		}
	}

	return mapp, nil
}

func NewWatchEventService(jsonData []models.GitResponseObject) WatchEventService {
	return &watchEventService{
		jsonData: jsonData,
	}
}
