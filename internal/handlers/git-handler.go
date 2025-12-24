package handlers

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"

	customerror "github.com/ndk123-web/github-activity/internal/custom-error"
	"github.com/ndk123-web/github-activity/internal/models"
	"github.com/ndk123-web/github-activity/internal/services"
)

type GitHandler interface {
	GetAllResponseObjects(url string)
}

type gitHandler struct{}

func (g *gitHandler) GetAllResponseObjects(url string) {
	// declare service
	pushEventService := services.NewPushEventsService()

	if url == "" {
		url = "https://api.github.com/users/ndk123-web/events"
	}

	response, err := http.Get(url)
	customerror.GlobalError(&err)

	data, err := io.ReadAll(response.Body)

	var jsonData []models.GitResponseObject
	err = json.Unmarshal(data, &jsonData)
	customerror.GlobalError(&err)

	// fmt.Print(jsonData[0])

	var totalPushEvents int64
	totalPushEvents, err = pushEventService.GetTotalPushEvents(jsonData)
	customerror.GlobalError(&err)

	fmt.Println("Output")
	fmt.Printf("- Total Push Events: %v", totalPushEvents)
}

func NewGitHandler() GitHandler {
	return &gitHandler{}
}
