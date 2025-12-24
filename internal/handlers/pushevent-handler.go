package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"net/http"

	customerror "github.com/ndk123-web/github-activity/internal/custom-error"
	"github.com/ndk123-web/github-activity/internal/models"
	"github.com/ndk123-web/github-activity/internal/services"
)

type GitHandler interface {
	GetAllResponseObjects() error
	GetResponseRepoWise(limit int64, data []models.GitResponseObject) error
}

type gitHandler struct {
	url string
}

func (g *gitHandler) GetAllResponseObjects() error {

	if g.url == "" {
		return customerror.Wrap("Username Not Exist / Provide Username", errors.New("Username Not Exist / Provide Username"))
	}

	url := g.url

	response, err := http.Get(url)
	if err != nil {
		return customerror.Wrap("http get failed", err)
	}

	// close client socket
	defer response.Body.Close()

	data, err := io.ReadAll(response.Body)
	if err != nil {
		return customerror.Wrap("reading response body failed", err)
	}

	var jsonData []models.GitResponseObject
	if err := json.Unmarshal(data, &jsonData); err != nil {
		return customerror.Wrap("json unmarshal failed", err)
	}
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
