package handlers

import (
	"fmt"
	"os"
	"sort"
	"text/tabwriter"

	"github.com/ndk123-web/github-activity/internal/models"
	"github.com/ndk123-web/github-activity/internal/services"
)

type SummaryAllHandler interface {
	GetAllSummary(limit int64, jsonData []models.GitResponseObject) error
}

type summaryAllHandler struct {
	url string
}

func (s *summaryAllHandler) GetAllSummary(limit int64, jsonData []models.GitResponseObject) error {
	summaryService := services.NewSummaryAllService(jsonData)

	mapp, err := summaryService.GetAllSummary(limit)
	if err != nil {
		return err
	}

	var totalEvents int64 = 0
	var totalPushes int64 = 0
	var totalPulls int64 = 0
	var totalIssues int64 = 0
	var totalWatches int64 = 0

	data, ok := mapp["pushEvents"]
	if ok {
		for _, c := range data {
			totalPushes += c
		}
	}

	data, ok = mapp["pullEvents"]
	if ok {
		for _, c := range data {
			totalPulls += c
		}
	}

	data, ok = mapp["issueEvents"]
	if ok {
		for _, c := range data {
			totalIssues += c
		}
	}

	data, ok = mapp["watchEvents"]
	if ok {
		for _, c := range data {
			totalWatches += c
		}
	}

	totalEvents = totalPushes + totalPulls + totalIssues + totalWatches

	fmt.Printf("📊 Activity Summary (last %v events)\n\n", limit)

	// Summary metrics table
	tw := tabwriter.NewWriter(os.Stdout, 0, 0, 2, ' ', 0)
	fmt.Fprintln(tw, "Metric\tCount")
	fmt.Fprintf(tw, "Total Events\t%d\n", totalEvents)
	fmt.Fprintf(tw, "Total Push Events\t%d\n", totalPushes)
	fmt.Fprintf(tw, "Total Pull Request Events\t%d\n", totalPulls)
	fmt.Fprintf(tw, "Total Issues Events\t%d\n", totalIssues)
	fmt.Fprintf(tw, "Total Watch Events\t%d\n", totalWatches)
	tw.Flush()
	fmt.Println()

	// Aggregated top repositories by total activity across all event types
	fmt.Println("Top Repositories (by total activity):")
	fmt.Println("------------------------------------------")

	totalByRepo := make(map[string]int64)
	orderedEventTypes := []string{"pushEvents", "pullEvents", "issueEvents", "watchEvents"}
	for _, eventType := range orderedEventTypes {
		repoData, ok := mapp[eventType]
		if !ok {
			continue
		}
		for repo, count := range repoData {
			totalByRepo[repo] += count
		}
	}

	type aggPair struct {
		repo  string
		count int64
	}
	topAgg := make([]aggPair, 0, len(totalByRepo))
	for repo, count := range totalByRepo {
		topAgg = append(topAgg, aggPair{repo: repo, count: count})
	}
	sort.Slice(topAgg, func(i, j int) bool {
		if topAgg[i].count == topAgg[j].count {
			return topAgg[i].repo < topAgg[j].repo
		}
		return topAgg[i].count > topAgg[j].count
	})

	for _, p := range topAgg {
		fmt.Printf("%-30s %d events\n", p.repo, p.count)
	}
	fmt.Printf("\n")

	return nil
}

func NewSummaryAllHandler(url string) SummaryAllHandler {
	return &summaryAllHandler{
		url: url,
	}
}
