package handlers

import (
	"fmt"

	"github.com/ndk123-web/github-activity/internal/models"
	"github.com/ndk123-web/github-activity/internal/services"
)

type IssueEventHandler interface {
	GetAllIssueEvents(jsonData []models.GitResponseObject) error
	GetIssueByState(state string, limit int64, jsonData []models.GitResponseObject) error
}

type issueEventHandler struct {
	issueType string
}

func (h *issueEventHandler) GetAllIssueEvents(jsonData []models.GitResponseObject) error {

	issueService := services.NewIssueEventService(jsonData)
	cnt, err := issueService.GetAllIssueEvents()

	if err != nil {
		return err
	}

	fmt.Printf("- Total Issues Events: %d\n", cnt)

	return nil
}

func (h *issueEventHandler) GetIssueByState(state string, limit int64, jsonData []models.GitResponseObject) error {
	issueService := services.NewIssueEventService(jsonData)
	result, err := issueService.GetIssueByState(state, limit)
	if err != nil {
		return err
	}

	for repo, count := range result {
		fmt.Printf("- Repo: %s, Issues with state '%s': %d\n", repo, state, count)
	}

	if len(result) == 0 {
		fmt.Printf("No issues found with state '%s'\n", state)
	}

	return nil
}

func NewIssueEventHandler(url string) IssueEventHandler {
	return &issueEventHandler{
		issueType: url,
	}
}
