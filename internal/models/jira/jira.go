package jira

import (
	"errors"
	"report/internal/config"

	"github.com/andygrunwald/go-jira"
)

type Client struct {
	Client *jira.Client
}

type IssueChan struct {
	List []jira.Issue
	Err  error
}

type GroupUsersChan struct {
	List []string
	Err  error
}

// NewJira Авторизация Jira Cloud
func NewJira(c *config.Config) (*Client, error) {
	auth := jira.BasicAuthTransport{
		Username: c.UserName,
		Password: c.Token,
	}

	client, err := jira.NewClient(auth.Client(), c.Host)
	if err != nil {
		return nil, err
	}

	return &Client{client}, nil
}

// GetTasks получить список задач
func (c *Client) GetTasks(jql string, issueChan chan IssueChan) {
	defer close(issueChan)

	var (
		last  int
		issue []jira.Issue
	)

	for {
		opt := &jira.SearchOptions{
			MaxResults: 1000,
			StartAt:    last,
		}

		chunk, resp, err := c.Client.Issue.Search(jql, opt)
		if err != nil {
			issueChan <- IssueChan{nil, errors.New("неверная структура JQL")}
		}

		total := resp.Total
		if issue == nil {
			issue = make([]jira.Issue, 0, total)
		}

		issue = append(issue, chunk...)
		last = resp.StartAt + len(chunk)

		if last >= total {
			issueChan <- IssueChan{issue, nil}
		}
	}
}

// GetGroupUsers получить список пользователей из группы
func (c *Client) GetGroupUsers(nameGroup string, groupUserChan chan GroupUsersChan) {
	defer close(groupUserChan)

	list, _, err := c.Client.Group.Get(nameGroup)
	group := make([]string, 0, len(list))

	if err != nil {
		groupUserChan <- GroupUsersChan{nil, errors.New("группа не найдена")}
	}

	for _, l := range list {
		group = append(group, l.DisplayName)
	}

	groupUserChan <- GroupUsersChan{group, nil}
}
