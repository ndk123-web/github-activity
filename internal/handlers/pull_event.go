package handlers

import (
	"fmt"
	"github.com/ndk123-web/github-activity/internal/models"
	"github.com/ndk123-web/github-activity/internal/services"
)

type PullHandler interface {
	GetAllPullRequests(jsonData []models.GitResponseObject) error
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

func NewPullHandler(url string) PullHandler {
	return &pullHandler{
		url: url,
	}
}
