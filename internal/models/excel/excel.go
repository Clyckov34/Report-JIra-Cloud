package excel

import (
	"github.com/andygrunwald/go-jira"
	"github.com/xuri/excelize/v2"
)

type Excel struct {
	File *excelize.File
}

func NewExcel() *Excel {
	return &Excel{
		File: excelize.NewFile(),
	}
}

// CreateTodoList создать общий лист с задачами
func (e *Excel) CreateTodoList(issues *[]jira.Issue, host string) error {
	const nameSheet = "Sheet1"

	if err := e.setTableIssueTitle(nameSheet); err != nil {
		return err
	}

	if err := e.setTableIssueData(nameSheet, host, issues); err != nil {
		return err
	}

	if err := e.setStyle(nameSheet); err != nil {
		return err
	}

	if err := e.File.SetSheetName(nameSheet, "Общий список задач"); err != nil {
		return err
	}

	return nil
}

// CreateProjectDevOps создать проектный лист разработчика
func (e *Excel) CreateProjectDevOps(issues *[]jira.Issue, groupUsers *[]string) error {
	nameSheet := "Проекты разработчика"

	if _, err := e.File.NewSheet(nameSheet); err != nil {
		return err
	}

	if err := e.setTableProjectTitle(nameSheet); err != nil {
		return err
	}

	if err := e.setTableProjectData(nameSheet, issues, groupUsers); err != nil {
		return err
	}

	if err := e.setStyle(nameSheet); err != nil {
		return err
	}

	return nil
}

// CreateTodoListUsers создать лист с задачами по отдельности (По пользователям)
func (e *Excel) CreateTodoListUsers(issue []jira.Issue, users []string, host string) error {
	for _, user := range users {

		if _, err := e.File.NewSheet(user); err != nil {
			return err
		}

		if err := e.setTableIssueTitle(user); err != nil {
			return err
		}

		if err := e.setTableIssueUsersData(user, host, &issue); err != nil {
			return err
		}

		if err := e.setStyle(user); err != nil {
			return err
		}

	}

	return nil
}

// SaveFile сохранить файл. (Так же можно указать параметры файла)
func (e *Excel) SaveFile(fileName string) error {
	return e.File.SaveAs(fileName, excelize.Options{})
}
