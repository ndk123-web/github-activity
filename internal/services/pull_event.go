package services

import "github.com/ndk123-web/github-activity/internal/models"

type PullEventService interface {
	GetAllPullRequests() (int64, error)
}

type pullEventService struct {
	data []models.GitResponseObject
}

func (p *pullEventService) GetAllPullRequests() (int64, error) {
	data := p.data
	var cnt int64 = 0

	for _, obj := range data {
		if obj.Type == "PullRequestEvent" {
			cnt++
		}
	}

	return cnt, nil
}

func NewPullEventService(data []models.GitResponseObject) PullEventService {
	return &pullEventService{
		data: data,
	}
}
