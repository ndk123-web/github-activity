package handlers

import (
	"fmt"
	"strings"

	"github.com/ndk123-web/github-activity/internal/models"
	"github.com/ndk123-web/github-activity/internal/services"
)

type PullHandler interface {
	GetAllPullRequests(jsonData []models.GitResponseObject) error
	GetPullRequestRepoWise(limit int64, state string, jsonData []models.GitResponseObject) error
}

type pullHandler struct {
	url string
}

func (p *pullHandler) GetAllPullRequests(jsonData []models.GitResponseObject) error {
	pullEvetService := services.NewPullEventService(jsonData)

	totalPullRequests, err := pullEvetService.GetAllPullRequests()
	if err != nil {
		return err
	}

	fmt.Printf("\n🔀 Pull Request Overview\n")
	fmt.Printf("- Pull Requests: %v total\n", totalPullRequests)
	return nil
}

func (p *pullHandler) GetPullRequestRepoWise(limit int64, state string, jsonData []models.GitResponseObject) error {
	pullEventService := services.NewPullEventService(jsonData)

	if limit == 0 {
		limit = 2 // default limit
	}

	mapp, err := pullEventService.GetPullRequestsRepoWise(limit, state)

	if err != nil {
		return err
	}
	repoHeader := "REPOSITORY"
	countHeader := fmt.Sprintf("PULL_REQUESTS (%s)", state)
	maxRepoLen := len(repoHeader)
	for repo := range mapp {
		if len(repo) > maxRepoLen {
			maxRepoLen = len(repo)
		}
	}
	repoWidth := maxRepoLen
	countWidth := len(countHeader)

	fmt.Printf("\n%-*s  %*s\n", repoWidth, repoHeader, countWidth, countHeader)
	fmt.Printf("%s  %s\n", strings.Repeat("-", repoWidth), strings.Repeat("-", countWidth))
	for repo, prcnt := range mapp {
		fmt.Printf("%-*s  %*d\n", repoWidth, repo, countWidth, prcnt)
	}

	if len(mapp) == 0 {
		fmt.Printf("ℹ️ No Pull Requests found with state: %s\n", state)
	}

	fmt.Println()
	return nil
}

func NewPullHandler(url string) PullHandler {
	return &pullHandler{
		url: url,
	}
}
