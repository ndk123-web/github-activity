package handlers

import (
	"fmt"

	"github.com/ndk123-web/github-activity/internal/models"
	"github.com/ndk123-web/github-activity/internal/services"
)

type IssueEventHandler interface {
	GetAllIssueEvents(jsonData []models.GitResponseObject) error
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

func NewIssueEventHandler(url string) IssueEventHandler {
	return &issueEventHandler{
		issueType: url,
	}
}
