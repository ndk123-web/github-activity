package services

import (
	"github.com/ndk123-web/github-activity/internal/models"
)

type RepoService interface {
	HandleInfoRepo(limit int64) (models.RepoObject, error)
	HandlerInfoRepoEvents(limit int64, repoName string, jsonData []models.GitResponseObject) (models.RepoEventsServiceResponse, error)
}

type repoService struct {
	jsonData models.RepoObject
}

func (r *repoService) HandleInfoRepo(limit int64) (models.RepoObject, error) {
	return r.jsonData, nil
}

func (r *repoService) HandlerInfoRepoEvents(limit int64, repoName string, jsonData []models.GitResponseObject) (models.RepoEventsServiceResponse, error) {

	var response models.RepoEventsServiceResponse

	for _, event := range jsonData {
		switch event.Type {
		case "PushEvent":
			if limit > 0 && response.PushEvents >= limit {
				break
			}
			response.PushEvents++
		case "WatchEvent":
			if limit > 0 && response.WatchEvents >= limit {
				break
			}
			response.WatchEvents++
		case "PullRequestEvent":
			if limit > 0 && response.PullEventService >= limit {
				break
			}
			response.PullEventService++
		case "IssuesEvent":
			if limit > 0 && response.IssueEventService >= limit {
				break
			}
			response.IssueEventService++
		}
	}
	return response, nil
}

func NewRepoService(jsonData models.RepoObject) RepoService {
	return &repoService{
		jsonData: jsonData,
	}
}
