package handlers

import (
	"fmt"
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

	// fmt.Println("Output")
	fmt.Printf("- Total Pull Requests: %v\n", totalPullRequests)
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
	// fmt.Println("Output")
	for repo, prcnt := range mapp {
		fmt.Printf("- Total Pull Requests that are %s  On Repository: %s is %v\n", state, repo, prcnt)
	}
	return nil
}

func NewPullHandler(url string) PullHandler {
	return &pullHandler{
		url: url,
	}
}
