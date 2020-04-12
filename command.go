package main

import (
	"fmt"
	"github.com/andygrunwald/go-jira"
)

// RemoveAssigneesFromDoneTickets is a function for remove assignee from done tickets for user from configs
func RemoveAssigneesFromDoneTickets() error {
	config, err := GetConfiguration()

	if err != nil {
		return fmt.Errorf("error when getting configurations: %s", err)
	}

	client, err := createClient(config)
	if err != nil {
		return fmt.Errorf("error when creating jira client: %s", err)
	}

	user, _, err := client.User.GetSelf()
	if err != nil {
		return fmt.Errorf("error when get self %s", err)
	}

	const maxResults = 50

	jql := fmt.Sprintf("assignee = '%s' and status = Done ORDER BY updatedDate ASC", user.DisplayName)
	errorCount := 0

	for {
		issues, _, err := client.Issue.Search(jql, &jira.SearchOptions{StartAt: errorCount, MaxResults: maxResults})

		if err != nil {
			return fmt.Errorf("error when search issues %s", err)
		}

		if len(issues) == 0 {
			break
		}

		errorCount += updateIssues(client, issues)
	}

	return nil
}

func updateIssues(client *jira.Client, issues []jira.Issue) (errorCount int) {
	errorCount = 0

	for _, issue := range issues {
		_, err := client.Issue.UpdateAssignee(issue.ID, nil)

		if err == nil {
			fmt.Printf("the issue was updated successfully: %s\r\n", issue.Key)
		} else {
			errorCount++
			fmt.Printf("the issue was not updated: %s. Reason: %s", issue.Key, err)
		}
	}

	return
}

func createClient(config Configuration) (client *jira.Client, err error) {
	tp := jira.BasicAuthTransport{
		Username: config.Login,
		Password: config.Token,
	}

	client, err = jira.NewClient(tp.Client(), config.Domain)
	return
}
