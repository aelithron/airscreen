package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"golang.org/x/term"
)

func CheckProjects(projects []Project, config Config) {
	for index := range projects {
    project := projects[index]
  	fmt.Printf("\nproject #%d (of %d): %s (airtable id: %s)\n", index + 1, len(projects), project.HackatimeProjs[0], project.ExternalID)
		fmt.Print("- creator info: ")
		if project.SlackID != "" { fmt.Printf("(slack: %s)", project.SlackID) }
		if project.HackatimeID != 0 { fmt.Printf("(hackatime: %d)", project.HackatimeID) }
		fmt.Printf("\n- repo link: %s\n", project.CodeUrl)
		fmt.Printf("- demo link: %s\n", project.DemoUrl)
  	fmt.Println("- hackatime projects:")
		for hackatime := range project.HackatimeProjs { fmt.Printf("  - %s\n", project.HackatimeProjs[hackatime]) }
		fmt.Printf("send this project to fraud review? (y/n) ")
		fd := os.Stdin.Fd()
		state, err := term.MakeRaw(int(fd))
		if err != nil { panic (err) }
		defer term.Restore(int(fd), state)
		reader := bufio.NewReader(os.Stdin)
		letter, _, err := reader.ReadRune()
		if err != nil { panic (err) }
		term.Restore(int(fd), state)
  	fmt.Println(string(letter))
		if strings.ToLower(string(letter)) == "y" {
			fmt.Println("sending...")
			SendProject(project, config)
		} else { fmt.Println("skipping!") }
	}
}