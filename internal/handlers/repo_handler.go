package handlers

import (
	"fmt"
	"strings"
	"time"

	"github.com/ndk123-web/github-activity/internal/models"
	"github.com/ndk123-web/github-activity/internal/services"
)

type RepoHandler interface {
	HandleInfoRepo(limit int64, jsonData models.RepoObject) error
	HandleInfoRepoEvents(limit int64, jsonData []models.GitResponseObject, repoName string) error
}

type repoHandler struct {
	url string
}

func (r *repoHandler) HandleInfoRepo(limit int64, jsonData models.RepoObject) error {

	repoService := services.NewRepoService(jsonData)

	mapp, err := repoService.HandleInfoRepo(limit)
	if err != nil {
		return err
	}

	fmt.Printf("\n📦 Repository Info\n\n")

	desc := mapp.Description
	if strings.TrimSpace(desc) == "" {
		desc = "—"
	}

	// Build topics display as bracketed tags
	topicsDisplay := "—"
	if len(mapp.Topics) > 0 {
		parts := make([]string, 0, len(mapp.Topics))
		for _, t := range mapp.Topics {
			tt := strings.TrimSpace(t)
			if tt == "" {
				continue
			}
			parts = append(parts, "["+tt+"]")
		}
		if len(parts) > 0 {
			topicsDisplay = strings.Join(parts, " ")
		}
	}

	// Two-column table similar to push/watch handlers
	fieldHeader := "FIELD"
	valueHeader := "VALUE"
	maxFieldLen := len(fieldHeader)
	labels := []string{"Name", "Description", "Primary Language", "License", "Topics", "Visibility", "Forks", "Stars"}
	for _, l := range labels {
		if len(l) > maxFieldLen {
			maxFieldLen = len(l)
		}
	}
	// Set a tidy minimum width for field column
	fieldWidth := maxFieldLen
	if fieldWidth < 18 {
		fieldWidth = 18
	}

	fmt.Printf("%-*s  %s\n", fieldWidth, fieldHeader, valueHeader)
	fmt.Printf("%s  %s\n", strings.Repeat("-", fieldWidth), strings.Repeat("-", 44))

	fmt.Printf("%-*s  %s\n", fieldWidth, "Name", mapp.FullName)
	fmt.Printf("%-*s  %s\n", fieldWidth, "Description", desc)
	fmt.Printf("%-*s  %s\n", fieldWidth, "Primary Language", mapp.Language)
	fmt.Printf("%-*s  %s\n", fieldWidth, "License", mapp.Licence.Name)
	fmt.Printf("%-*s  %s\n", fieldWidth, "Topics", topicsDisplay)

	fmt.Printf("%-*s  %s\n", fieldWidth, "Visibility", mapp.Visibility)
	fmt.Printf("%-*s  %d\n", fieldWidth, "Forks", mapp.Forks)
	fmt.Printf("%-*s  %d\n", fieldWidth, "Stars", mapp.Stars)

	fmt.Printf("%-*s  %d\n", fieldWidth, "Open Issues", mapp.OpenIssues)

	t1, err := time.Parse(time.RFC3339, mapp.CreatedAt)
	if err != nil {
		return err
	}
	t1 = t1.Local()
	fmt.Printf("%-*s  %s\n", fieldWidth, "Created", t1.Format("2006-01-02"))

	t2, err := time.Parse(time.RFC3339, mapp.UpdatedAt)
	if err != nil {
		return err
	}
	t2 = t2.Local()
	fmt.Printf("%-*s  %s\n", fieldWidth, "Last Updated", t2.Format("2006-01-02"))

	t3, err := time.Parse(time.RFC3339, mapp.PushedAt)
	if err != nil {
		return err
	}
	t3 = t3.Local()
	fmt.Printf("%-*s  %s\n", fieldWidth, "Last Push", t3.Format("2006-01-02"))

	return nil
}

func (r *repoHandler) HandleInfoRepoEvents(limit int64, jsonData []models.GitResponseObject, repoName string) error {
	repoService := services.NewRepoService(models.RepoObject{})

	response, err := repoService.HandlerInfoRepoEvents(limit, repoName, jsonData)
	if err != nil {
		return err
	}
	fmt.Printf("\n📦 Repository Events (recent)\n")

	// Two-column table: EVENT | COUNT
	eventHeader := "EVENT"
	countHeader := "COUNT"
	labels := []string{"Push Events", "Watch Events", "Pull Requests", "Issues"}
	maxLabelLen := len(eventHeader)
	for _, l := range labels {
		if len(l) > maxLabelLen {
			maxLabelLen = len(l)
		}
	}
	eventWidth := maxLabelLen
	countWidth := len(countHeader)

	// Underline the section title
	fmt.Println(strings.Repeat("-", eventWidth+2+countWidth))
	fmt.Printf("%-*s  %*s\n", eventWidth, eventHeader, countWidth, countHeader)
	fmt.Printf("%s  %s\n", strings.Repeat("-", eventWidth), strings.Repeat("-", countWidth))

	fmt.Printf("%-*s  %*d\n", eventWidth, "Push Events", countWidth, response.PushEvents)
	fmt.Printf("%-*s  %*d\n", eventWidth, "Watch Events", countWidth, response.WatchEvents)
	fmt.Printf("%-*s  %*d\n", eventWidth, "Pull Requests", countWidth, response.PullEventService)
	fmt.Printf("%-*s  %*d\n", eventWidth, "Issues", countWidth, response.IssueEventService)

	fmt.Println()
	return nil
}

func NewRepoHandler(url string) RepoHandler {
	return &repoHandler{
		url: url,
	}
}
