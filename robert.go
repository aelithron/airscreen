package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"time"
)

type robertReq struct {
	Name string `json:"name"`
	Kind string `json:"kind"`
	CodeUrl string `json:"codeUrl"`
	DemoUrl string `json:"demoUrl"`
	HackatimeProjs []string `json:"hackatimeProjects"`
	Submitter submitter `json:"submitter"`
	ExternalID string `json:"externalId"`
}
type submitter struct {
  SlackID string `json:"slackId,omitempty"`
	HackatimeID int `json:"hackatimeId,omitempty"`
}
type robertRes struct {
	Error bool `json:"error"`
	Message string `json:"message"`
	ID string `json:"id"`
	Created bool `json:"created"`
}

func SendProject(project Project, config Config) {
	url, err := url.Parse("https://telescreen.hackclub.com/api/v1/ysws/projects")
	reqBody := robertReq{ Name: project.HackatimeProjs[0], Kind: "hackatime", CodeUrl: project.CodeUrl, DemoUrl: project.DemoUrl, HackatimeProjs: project.HackatimeProjs, Submitter: submitter{ SlackID: project.SlackID, HackatimeID: project.HackatimeID }, ExternalID: project.ExternalID }
	jsonBody, _ := json.Marshal(reqBody)
	req, err := http.NewRequest(http.MethodPost, url.String(), bytes.NewBuffer(jsonBody))
	if err != nil { panic (err) }
	req.Header.Add("Authorization", "Bearer " + config.RobertKey)
	client := http.Client{ Timeout: 10 * time.Second }
	resp, err := client.Do(req)
	if err != nil { panic (err) }
	defer resp.Body.Close()
	resBody, err := io.ReadAll(resp.Body)
	if err != nil { panic (err) }
	var body robertRes
	if err := json.Unmarshal(resBody, &body); err != nil { panic (err) }
	if body.Error == true {
		fmt.Printf("[error] issue submitting to robert: %s (status: %d)\n", body.Message, resp.StatusCode)
		return
	}
	if body.Created == false {
		fmt.Printf("[error] this project was already submitted to robert! id: %s\n", body.ID)
		return
	}
	fmt.Printf("submitted to robert! id: %s\n", body.ID)
}