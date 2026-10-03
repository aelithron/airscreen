package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"path/filepath"
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
  home, err := os.UserConfigDir()
  if err != nil { panic(err) }
  path := filepath.Join(home, "airscreen", "config.json")
  var config Config
  configFile, err := os.ReadFile(path)
  if err != nil {
    if errors.Is(err, os.ErrNotExist) {
      fmt.Println("writing config file to " + `"` + path + `"` + "...")
      file, err := os.Create(path)
      if (err != nil) { panic (err) }
      if err := os.Chmod(path, 600); err != nil { panic (err) }
      encoder := json.NewEncoder(file)
      encoder.SetIndent("", "  ")
      empty := Config{ AirtablePAT: "", AirtableBase: "", AirtableTable: "", RobertKey: "" }
      if err := encoder.Encode(empty); (err != nil) { panic (err) }
      config = empty
    } else { panic (err) }
  }

}