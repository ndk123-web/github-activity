package services

import "github.com/ndk123-web/github-activity/internal/models"

type PushEvents interface {
	GetTotalPushEvents(data []models.GitResponseObject) (int64, error)
}

type pushEvents struct{}

func (p *pushEvents) GetTotalPushEvents(data []models.GitResponseObject) (int64, error) {
	// mapp := make(map[string]int)
	var total int64 = 0

	for _, obj := range data {
		if obj.Type == "PushEvent" {
			total++
		}
	}

	return total, nil
}

func NewPushEventsService() PushEvents {
	return &pushEvents{}
}
