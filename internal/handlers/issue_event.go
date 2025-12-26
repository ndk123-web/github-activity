package handlers

import (
	"fmt"
	"strings"

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

	fmt.Printf("\n🐛 Issue Activity Overview\n")
	fmt.Printf("- Issue Events: %d total\n", cnt)

	return nil
}

func (h *issueEventHandler) GetIssueByState(state string, limit int64, jsonData []models.GitResponseObject) error {
	issueService := services.NewIssueEventService(jsonData)
	result, err := issueService.GetIssueByState(state, limit)
	if err != nil {
		return err
	}

	repoHeader := "REPOSITORY"
	countHeader := fmt.Sprintf("ISSUES (%s)", state)
	maxRepoLen := len(repoHeader)
	for repo := range result {
		if len(repo) > maxRepoLen {
			maxRepoLen = len(repo)
		}
	}
	repoWidth := maxRepoLen
	countWidth := len(countHeader)

	fmt.Printf("\n%-*s  %*s\n", repoWidth, repoHeader, countWidth, countHeader)
	fmt.Printf("%s  %s\n", strings.Repeat("-", repoWidth), strings.Repeat("-", countWidth))
	for repo, count := range result {
		fmt.Printf("%-*s  %*d\n", repoWidth, repo, countWidth, count)
	}

	if len(result) == 0 {
		fmt.Printf("ℹ️ No issues found with state '%s'\n", state)
	}

	fmt.Println()

	return nil
}

func NewIssueEventHandler(url string) IssueEventHandler {
	return &issueEventHandler{
		issueType: url,
	}
}
