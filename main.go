package main

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/andygrunwald/go-jira"
)

// Configuration for the application
type Configuration struct {
	login  string
	token  string
	domain string
}

// GetConfiguration is a function for get configurations from config.json
func GetConfiguration() Configuration {
	file, _ := os.Open("config.json")
	defer file.Close()
	decoder := json.NewDecoder(file)
	configuration := Configuration{}
	err := decoder.Decode(&configuration)

	return configuration
}

func main() {
	config := GetConfiguration()

	tp := jira.BasicAuthTransport{
		Username: config.login,
		Password: config.token,
	}

	client, _ := jira.NewClient(tp.Client(), config.domain)

	for {
		opt := &jira.SearchOptions{MaxResults: 100}
		issues, _, err := client.Issue.Search("assignee = akirinyuk and status = Done ORDER BY updatedDate ASC", opt)
		if err != nil {
			fmt.Errorf("Error: ", err)
		}
		if len(issues) == 0 {
			break
		}

		for _, issue := range issues {
			client.Issue.UpdateAssignee(issue.ID, nil)
			fmt.Printf("Updated issue:\n\r", issue.Key)
		}
	}
}
