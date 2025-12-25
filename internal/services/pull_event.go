package services

import "github.com/ndk123-web/github-activity/internal/models"

type PullEventService interface {
	GetAllPullRequests() (int64, error)
	GetPullRequestsRepoWise(limit int64, state string) (map[string]int64, error)
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

func (p *pullEventService) GetPullRequestsRepoWise(limit int64, state string) (map[string]int64, error) {
	data := p.data
	mapp := make(map[string]int64)

	for _, obj := range data {
		if obj.Type == "PullRequestEvent" {
			prState := obj.Payload.Action
			if state == "all" || prState == state {
				repoName := obj.Repo.Name
				mapp[repoName]++
				if int64(len(mapp)) >= limit {
					break
				}
			}
		}
		if len(mapp) >= int(limit) {
			break
		}
	}

	return mapp, nil
}

func NewPullEventService(data []models.GitResponseObject) PullEventService {
	return &pullEventService{
		data: data,
	}
}
