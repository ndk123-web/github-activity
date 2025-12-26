package services

import "github.com/ndk123-web/github-activity/internal/models"

type IssueEventService interface {
	GetAllIssueEvents() (int64, error)
}

type issueEventService struct {
	jsonData []models.GitResponseObject
}

func (s *issueEventService) GetAllIssueEvents() (int64, error) {
	var cnt int64 = 0

	for _, item := range s.jsonData {
		if item.Type == "IssuesEvent" {
			cnt++
		}
	}

	return cnt, nil
}

func NewIssueEventService(data []models.GitResponseObject) IssueEventService {
	return &issueEventService{
		jsonData: data,
	}
}
