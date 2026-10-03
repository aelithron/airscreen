package main

import (
	"encoding/json"
	"io"
	"net/http"
	"net/url"
	"strings"
	"time"
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
	if err != nil { return nil, err }
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil { return nil, err }
	req.Header.Add("Authorization", "Bearer " + "(key)")
	client := http.Client{ Timeout: 10 * time.Second }
	resp, err := client.Do(req)
	if err != nil { panic (err) }
	defer resp.Body.Close()
	resBody, err := io.ReadAll(resp.Body)
	if err != nil { return nil, err }
	var body airtableResp
	if err := json.Unmarshal(resBody, &body); err != nil { panic (err) }
	var projects []Project
	for recordIndex := range body.Records {
		rec := body.Records[recordIndex]
		var newProject = &Project{
			CodeUrl: rec.Fields.CodeUrl,
			DemoUrl: rec.Fields.DemoUrl,
			HackatimeProjs: strings.Split(rec.Fields.HackatimeProjs, ", "),
			SlackID: rec.Fields.SlackID,
			HackatimeID: rec.Fields.HackatimeID,
			ExternalID: rec.ExternalID,
		}
		projects = append(projects, *newProject)
	}
	return projects, nil
}