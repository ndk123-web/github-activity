package models

import "time"

type ActorModel struct {
	Id           string `json:"id"`
	Login        string `json:"login"`
	DisplayLogin string `json:"display_login"`
	GravatarId   string `json:"gravatar_id"`
	URL          string `json:"url"`
	AvtarUrl     string `json:"avatar_url"`
}

type RepoModel struct {
	Id   string `json:"id"`
	Name string `json:"name"`
	Url  string `json:"url"`
}

type PayloadModel struct {
	RepositoryId string `json:"repository_id"`
	Ref          string `json:"ref"`
	PushId       string `json:"push_id"`
	Head         string `json:"head"`
	Before       string `json:"before"`
}

type GitResponseObject struct {
	Id        string       `json:"id"`
	Type      string       `json:"type"`
	Actor     ActorModel   `json:"actor"`
	Repo      RepoModel    `json:"repo"`
	Payload   PayloadModel `json:"payload"`
	Public    bool         `json:"public"`
	CreatedAt time.Time    `json:"created_at"`
}
