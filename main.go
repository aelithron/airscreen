package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
	"strings"
  "golang.org/x/term"
)

type Project struct {
	CodeUrl string
	DemoUrl string
	HackatimeProjs []string
	SlackID string
	HackatimeID int
	ExternalID string
}
type Config struct {
  AirtablePAT string `json:"airtablePAT"`
  AirtableBase string `json:"airtableBase"`
  AirtableTable string `json:"airtableTable"`
  RobertKey string `json:"robertKey"`
}

func main() {
	fmt.Println("welcome to airscreen! :3")
  home, err := os.UserHomeDir()
  if err != nil { panic(err) }
  path := filepath.Join(home, ".airscreen", "config.json")
  var config Config
  configFile, err := os.ReadFile(path)
  if err != nil {
    if errors.Is(err, os.ErrNotExist) {
      fmt.Printf(`writing empty config file to "%s"...` + "\n", path)
      err := os.MkdirAll(filepath.Dir(path), 0700)
      file, err := os.Create(path)
      if (err != nil) { panic (err) }
      if err := os.Chmod(path, 0600); err != nil { panic (err) }
      encoder := json.NewEncoder(file)
      encoder.SetIndent("", "  ")
      empty := Config{ AirtablePAT: "", AirtableBase: "", AirtableTable: "", RobertKey: "" }
      if err := encoder.Encode(empty); (err != nil) { panic (err) }
      config = empty
    } else { panic (err) }
  } else if err := json.Unmarshal(configFile, &config); err != nil { panic (err) }

  reader := bufio.NewReader(os.Stdin)
  if config.AirtablePAT == "" {
    fmt.Println("\n[setup] enter your Airtable Personal Access Token below! (will not display in terminal)")
    fmt.Print("Airtable PAT: ")
    pat, err := term.ReadPassword(int(os.Stdin.Fd()))
    if (err != nil) { panic (err) }
    config.AirtablePAT = string(pat)
  }
  if config.AirtableBase == "" {
    fmt.Println("\n[setup] enter your Airtable Base ID below!")
    fmt.Print("Airtable Base ID: ")
    base, _ := reader.ReadString('\n')
    config.AirtableBase = strings.TrimSpace(base)
  }
  if config.AirtableTable == "" {
    fmt.Println("\n[setup] enter your Airtable Table ID below!")
    fmt.Print("Airtable Table ID: ")
    table, _ := reader.ReadString('\n')
    config.AirtableTable = strings.TrimSpace(table)
  }
  if config.RobertKey == "" {
    fmt.Println("\n[setup] enter your Robert API key below! (will not display in terminal)")
    fmt.Print("Robert Key: ")
    robert, err := term.ReadPassword(int(os.Stdin.Fd()))
    if (err != nil) { panic (err) }
    config.RobertKey = string(robert)
  }
  file, err := os.OpenFile(path, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
  if (err != nil) { panic (err) }
  encoder := json.NewEncoder(file)
  encoder.SetIndent("", "  ")
  encoder.Encode(config)

  fmt.Println("\nloading projects from Airtable...")
  projects, err := LoadAirtable(config)
  if err != nil { panic (err) }
  fmt.Printf("found %d projects!\n", len(projects))
  for index := range projects {
    CheckProject(projects[index], index, len(projects))
  }
}