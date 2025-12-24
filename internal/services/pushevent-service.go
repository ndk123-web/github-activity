package services

import (
	"errors"

	customerror "github.com/ndk123-web/github-activity/internal/custom-error"
	"github.com/ndk123-web/github-activity/internal/models"
)

type PushEvents interface {
	GetTotalPushEvents() (int64, error)
	GetPushEventsRepoWise(limit int64) (map[string]int64, error)
}

type pushEvents struct {
	data []models.GitResponseObject
}

func (p *pushEvents) GetTotalPushEvents() (int64, error) {
	// mapp := make(map[string]int)
	var total int64 = 0

	for _, obj := range p.data {
		if obj.Type == "PushEvent" {
			total++
		}
	}

	return total, nil
}

func (p *pushEvents) GetPushEventsRepoWise(limit int64) (map[string]int64, error) {
	mapp := make(map[string]int64)

	for _, obj := range p.data {
		repoName := obj.Repo.Name
		if repoName == "" {
			customerror.Wrap("Repo Name is Empty", errors.New("Repository Name Should not be Empty"))
		}
		mapp[repoName]++
		if int64(len(mapp)) >= limit {
			break
		}
	}

	return mapp, nil
}

func NewPushEventsService(data []models.GitResponseObject) PushEvents {
	return &pushEvents{
		data: data,
	}
}
