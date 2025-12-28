package models

type LicenseObject struct {
	Name string `json:"name"`
}

type RepoObject struct {
	Id          int64         `json:"id"`
	Description string        `json:"description"`
	Licence     LicenseObject `json:"license"`
	Topics      []string      `json:"topics"`
	Visibility  string        `json:"visibility"`
	Forks       int64         `json:"forks_count"`
	Stars       int64         `json:"watchers"`
	FullName    string        `json:"full_name"`
	Language    string        `json:"language"`
	OpenIssues  int64         `json:"open_issues"`
	CreatedAt   string        `json:"created_at"`
	UpdatedAt   string        `json:"updated_at"`
	PushedAt    string        `json:"pushed_at"`
}

type RepoEventsServiceResponse struct {
	PushEvents        int64
	WatchEvents       int64
	PullEventService  int64
	IssueEventService int64
}
