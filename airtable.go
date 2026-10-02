package airscreen

type Project struct {
	CodeUrl string `json:"Code URL"`
	DemoUrl string `json:"Playable URL"`
	HackatimeProjs []string
	SlackID string `json:"Slack ID,omitempty"`
	HackatimeID int `json:"Hackatime ID,omitempty"`
}