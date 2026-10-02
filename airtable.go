package airscreen
	
import (
	"net/http"
	"net/url"
	"time"
)

type Project struct {
	CodeUrl string `json:"Code URL"`
	DemoUrl string `json:"Playable URL"`
	HackatimeProjs []string
	SlackID string `json:"Slack ID,omitempty"`
	HackatimeID int `json:"Hackatime ID,omitempty"`
}

func LoadAirtable() ([]Project, error) {
	url, err := url.Parse()
	req, err := http.NewRequest(http.MethodGet, url.String(), nil)
	if err != nil { return nil, err }
	client := http.Client{ Timeout: 10 * time.Second }
	resp, err := client.Do(req);
	return nil, nil
}