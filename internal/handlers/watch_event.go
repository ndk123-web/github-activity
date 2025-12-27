package handlers

import (
	"fmt"
	"strings"

	customerror "github.com/ndk123-web/github-activity/internal/custom-error"
	"github.com/ndk123-web/github-activity/internal/models"
	"github.com/ndk123-web/github-activity/internal/services"
)

type WatchEventHandler interface {
	GetAllWatchEvent(jsonData []models.GitResponseObject, limit int64) error
}

type watchEventHandler struct {
	url string
}

func (w *watchEventHandler) GetAllWatchEvent(jsonData []models.GitResponseObject, limit int64) error {
	watchEventService := services.NewWatchEventService(jsonData)

	mapp, err := watchEventService.GetAllWatchEvent(limit)
	if err != nil {
		customerror.Wrap("Error: ", fmt.Errorf("watch event handler: %s", err.Error()))
		return err
	}

	// Overview summary
	var total int64 = 0
	for _, v := range mapp {
		total += v
	}
	fmt.Printf("\n⭐ Watch Activity Overview\n")
	fmt.Printf("- Watch Events: %d total\n", total)

	// Tabular output similar to other handlers
	repoHeader := "REPOSITORY"
	countHeader := "WATCH_EVENTS"
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
	for repo, count := range mapp {
		fmt.Printf("%-*s  %*d\n", repoWidth, repo, countWidth, count)
	}

	fmt.Println()

	return nil
}

func NewWatchEventHandler(url string) WatchEventHandler {
	return &watchEventHandler{
		url: url,
	}
}
