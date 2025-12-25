package handlers

import (
	// "encoding/json"
	// "errors"
	"fmt"
	// "io"
	// "net/http"

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

	// fmt.Println("Output")
	fmt.Printf("- Total Push Events: %v\n", totalPushEvents)

	return nil
}

func (g *gitHandler) GetResponseRepoWise(limit int64, jsonData []models.GitResponseObject) error {

	pushEventService := services.NewPushEventsService(jsonData)

	mapp, err := pushEventService.GetPushEventsRepoWise(limit)
	if err != nil {
		return customerror.Wrap("Issue In GetPushEventRepoWise Handler", err)
	}

	for repo, pushcnt := range mapp {
		fmt.Printf("- Total Push On Repository: %s is %v\n", repo, pushcnt)
	}

	return nil
}

func NewGitHandler(url string) GitHandler {
	return &gitHandler{
		url: url,
	}
}
