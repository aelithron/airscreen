package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

type robertReq struct {
	Name string `json:"name"`
	Kind string `json:"kind"`
	CodeUrl string `json:"codeUrl"`
	DemoUrl string `json:"playableUrl"`
	HackatimeProjs []string `json:"hackatimeProjs"`
	Submitter Submitter
	ExternalID string `json:"externalId"`
}
type Submitter struct {
  SlackID string `json:"slackId,omitempty"`
	HackatimeID int `json:"hackatimeId,omitempty"`
}

func SendProject(project Project) {
	url, err := url.Parse("https://telescreen.hackclub.com/api/v1/ysws/projects")
	reqBody := robertReq{ Name: project.HackatimeProjs[0], Kind: "hackatime", CodeUrl: project.CodeUrl, DemoUrl: project.DemoUrl, HackatimeProjs: project.HackatimeProjs, Submitter: Submitter{ SlackID: project.SlackID, HackatimeID: project.HackatimeID }, ExternalID: project.ExternalID }
	body, _ := json.Marshal(reqBody)
	req, err := http.NewRequest(http.MethodPost, url.String(), bytes.NewBuffer(body))
	if err != nil { panic (err) }
	req.Header.Add("Authorization", "Bearer " + "(key)")
	client := http.Client{ Timeout: 10 * time.Second }
	resp, err := client.Do(req)
	if err != nil { panic (err) }
	defer resp.Body.Close()
}