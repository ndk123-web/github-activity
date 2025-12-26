package handlers

import (
	// "encoding/json"
	// "errors"
	"fmt"
	// "io"
	// "net/http"
	"strings"

	customerror "github.com/ndk123-web/github-activity/internal/custom-error"
	"github.com/ndk123-web/github-activity/internal/models"
	"github.com/ndk123-web/github-activity/internal/services"
)

type GitHandler interface {
	GetAllResponseObjects(jsonData []models.GitResponseObject) error
	GetResponseRepoWise(limit int64, data []models.GitResponseObject) error
}

type gitHandler struct {
	url string
}

func (g *gitHandler) GetAllResponseObjects(jsonData []models.GitResponseObject) error {

	pushEventService := services.NewPushEventsService(jsonData)

	totalPushEvents, err := pushEventService.GetTotalPushEvents()
	if err != nil {
		return customerror.Wrap("counting push events failed", err)
	}

	fmt.Printf("\n🚀 Push Activity Overview\n")
	fmt.Printf("- Pushes: %v events found\n", totalPushEvents)

	return nil
}

func (g *gitHandler) GetResponseRepoWise(limit int64, jsonData []models.GitResponseObject) error {

	pushEventService := services.NewPushEventsService(jsonData)

	mapp, err := pushEventService.GetPushEventsRepoWise(limit)
	if err != nil {
		return customerror.Wrap("Issue In GetPushEventRepoWise Handler", err)
	}

	repoHeader := "REPOSITORY"
	countHeader := "PUSH_EVENTS"
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
	for repo, pushcnt := range mapp {
		fmt.Printf("%-*s  %*d\n", repoWidth, repo, countWidth, pushcnt)
	}

	fmt.Println()
	return nil
}

func NewGitHandler(url string) GitHandler {
	return &gitHandler{
		url: url,
	}
}
