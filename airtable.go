package airscreen

import (
	"io"
	"net/http"
	"net/url"
	"time"
	"encoding/json"
)

type airtableResp struct {
  Records []struct {
		Fields struct {
      CodeUrl string `json:"Code URL"`
			DemoUrl string `json:"Playable URL"`
			HackatimeProjs string `json:"Hackatime Project Name"`
			SlackID string `json:"Slack ID"`
			HackatimeID int `json:"Hackatime ID"`
		}
		ExternalID string `json:"id"`
	} `json:"records"`
}

func LoadAirtable() ([]Project, error) {
	url, err := url.JoinPath("https://api.airtable.com/v0", "(base)", "/", "(table)")
	req, err := http.NewRequest(http.MethodGet, url, nil)
	req.Header.Add("Authorization", "Bearer " + "(key)")
	if err != nil { return nil, err }
	client := http.Client{ Timeout: 10 * time.Second }
	resp, err := client.Do(req)
	if err != nil { panic (err) }
	defer resp.Body.Close()
	resBody, err := io.ReadAll(resp.Body)
	var body airtableResp
	if err := json.Unmarshal(resBody, &body); err != nil { panic (err) }
	if err != nil { panic (err) }
	return nil, nil
}