package main

type Actor struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

type Repo struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

type Event struct {
	Id    string `json:"id"`
	Type  string `json:"type"`
	Actor Actor  `json:"actor"`
	Repo  Repo   `json:"repo"`
}

// Repository and Branch events
const (
	EventTypeCreate  = "CreateEvent"
	EventTypeDelete  = "DeleteEvent"
	EventTypePublic  = "PublicEvent"
	EventTypePush    = "PushEvent"
	EventTypeRelease = "ReleaseEvent"
)
