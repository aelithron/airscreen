package main

import (
	"io"
	"net/http"
	"net/url"
	"time"
	"encoding/json"
)

type robertReq struct {
  Records []struct {
		CodeUrl string `json:"Code URL"`
		DemoUrl string `json:"Playable URL"`
		HackatimeProjs []string `json:"Hackatime Project Name"`
		SlackID string `json:"Slack ID"`
		HackatimeID int `json:"Hackatime ID"`
		ExternalID string `json:"externalId"`
	} `json:"records"`
}

func SendProject(project Project) (error) {
	url, err := url.Parse("https://telescreen.hackclub.com/api/v1/ysws/projects")
	req, err := http.NewRequest(http.MethodPost, url.String(), nil)
	req.Header.Add("Authorization", "Bearer " + "(key)")
	if err != nil { return err }
	client := http.Client{ Timeout: 10 * time.Second }
	resp, err := client.Do(req)
	if err != nil { panic (err) }
	defer resp.Body.Close()
	resBody, err := io.ReadAll(resp.Body)
	var body airtableResp
	if err := json.Unmarshal(resBody, &body); err != nil { panic (err) }
	if err != nil { panic (err) }
	return nil
}