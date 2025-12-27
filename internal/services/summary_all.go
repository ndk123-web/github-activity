package services

import "github.com/ndk123-web/github-activity/internal/models"

type SummaryAllService interface {
	GetAllSummary(limit int64) (map[string]map[string]int64, error)
}

type summaryAllService struct {
	jsonData []models.GitResponseObject
}

func IsGreaterThanLimit(currentCount int, limit int64) bool {
	return limit > 0 && int64(currentCount) >= limit
}

func (s *summaryAllService) GetAllSummary(limit int64) (map[string]map[string]int64, error) {
	data := s.jsonData
	var cnt int64 = 0

	pushEvents := make(map[string]int64)
	pullEvents := make(map[string]int64)
	issueEvents := make(map[string]int64)
	watchEvents := make(map[string]int64)

	for _, item := range data {

		switch item.Type {
		case "PushEvent":
			{
				if IsGreaterThanLimit(len(pushEvents), limit) {
					break
				}
				pushEvents[item.Repo.Name]++
			}
		case "PullRequestEvent":
			{
				if IsGreaterThanLimit(len(pullEvents), limit) {
					break
				}
				pullEvents[item.Repo.Name]++
			}
		case "IssuesEvent":
			{
				if IsGreaterThanLimit(len(issueEvents), limit) {
					break
				}
				issueEvents[item.Repo.Name]++
			}
		case "WatchEvent":
			{
				if IsGreaterThanLimit(len(watchEvents), limit) {
					break
				}
				watchEvents[item.Repo.Name]++
			}
		}
		cnt++
	}

	return map[string]map[string]int64{
		"pushEvents":  pushEvents,
		"pullEvents":  pullEvents,
		"issueEvents": issueEvents,
		"watchEvents": watchEvents,
	}, nil
}

func NewSummaryAllService(jsonData []models.GitResponseObject) SummaryAllService {
	return &summaryAllService{
		jsonData: jsonData,
	}
}
