package services

import (
	// "fmt"
	"github.com/ndk123-web/github-activity/internal/models"
)

type IssueEventService interface {
	GetAllIssueEvents() (int64, error)
	GetIssueByState(state string, limit int64) (map[string]int, error)
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

func (s *issueEventService) GetIssueByState(state string, limit int64) (map[string]int, error) {
	result := make(map[string]int)

	for _, item := range s.jsonData {
		if item.Type == "IssuesEvent" {
			// link repo to count
			issueState := item.Payload.Issues.State
			// fmt.Println("State: ", state)
			// fmt.Println("Issue State: ", issueState)
			if issueState == state {
				result[item.Repo.Name]++
			}
			limit--
		}
		if limit < 0 {
			break 
		}
	}

	return result, nil
}

func NewIssueEventService(data []models.GitResponseObject) IssueEventService {
	return &issueEventService{
		jsonData: data,
	}
}
