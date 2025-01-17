package excel

import (
	"fmt"
	"strings"

	"github.com/andygrunwald/go-jira"
)

// setTableIssueTitle заголовок таблицы
func (e *Excel) setTableIssueTitle(nameSheet string) error {
	if err := e.File.SetCellValue(nameSheet, "A1", "Ключ"); err != nil {
		return err
	}

	if err := e.File.SetCellValue(nameSheet, "B1", "№ Задача"); err != nil {
		return err
	}

	if err := e.File.SetCellValue(nameSheet, "C1", "URL задачи"); err != nil {
		return err
	}

	if err := e.File.SetCellValue(nameSheet, "D1", "Название задачи"); err != nil {
		return err
	}

	if err := e.File.SetCellValue(nameSheet, "E1", "Тип задачи"); err != nil {
		return err
	}

	if err := e.File.SetCellValue(nameSheet, "F1", "Статус задачи"); err != nil {
		return err
	}

	if err := e.File.SetCellValue(nameSheet, "G1", "Автор"); err != nil {
		return err
	}

	if err := e.File.SetCellValue(nameSheet, "H1", "Исполнитель"); err != nil {
		return err
	}

	if err := e.File.SetCellValue(nameSheet, "I1", "Задача Создана"); err != nil {
		return err
	}

	if err := e.File.SetCellValue(nameSheet, "J1", "Задача Закрыта"); err != nil {
		return err
	}

	if err := e.File.SetCellValue(nameSheet, "K1", "Задача Изменена"); err != nil {
		return err
	}

	return nil
}

// setTableIssue запись данных в ячейки
func (e *Excel) setTableIssueData(nameSheet, host string, issues *[]jira.Issue) error {
	for key, issue := range *issues {

		dateCreateTask, err := issue.Fields.Created.MarshalJSON()
		if err != nil {
			return err
		}

		dateResolutionTask, err := issue.Fields.Resolutiondate.MarshalJSON()
		if err != nil {
			return err
		}

		dateUpdateTask, err := issue.Fields.Updated.MarshalJSON()
		if err != nil {
			return err
		}

		lineNumber := key + 2 // Записываем со второй строчке

		if err := e.File.SetCellValue(nameSheet, fmt.Sprintf("A%d", lineNumber), keyTaskTag(issue.Key)); err != nil {
			return err
		}

		if err := e.File.SetCellValue(nameSheet, fmt.Sprintf("B%v", lineNumber), issue.Key); err != nil {
			return err
		}

		if err := e.File.SetCellFormula(nameSheet, fmt.Sprintf("C%v", lineNumber), `HYPERLINK("`+host+`/browse/`+issue.Key+`")`); err != nil {
			return err
		}

		if err := e.File.SetCellValue(nameSheet, fmt.Sprintf("D%v", lineNumber), issue.Fields.Summary); err != nil {
			return err
		}

		if err := e.File.SetCellValue(nameSheet, fmt.Sprintf("E%v", lineNumber), issue.Fields.Type.Name); err != nil {
			return err
		}

		if err := e.File.SetCellValue(nameSheet, fmt.Sprintf("F%v", lineNumber), issue.Fields.Status.Name); err != nil {
			return err
		}

		if err := e.File.SetCellValue(nameSheet, fmt.Sprintf("G%v", lineNumber), issue.Fields.Creator.DisplayName); err != nil {
			return err
		}

		if err := e.File.SetCellValue(nameSheet, fmt.Sprintf("H%v", lineNumber), issue.Fields.Assignee.DisplayName); err != nil {
			return err
		}

		if err := e.File.SetCellValue(nameSheet, fmt.Sprintf("I%v", lineNumber), parseDateTime(dateCreateTask)); err != nil {
			return err
		}

		if err := e.File.SetCellValue(nameSheet, fmt.Sprintf("J%v", lineNumber), parseDateTime(dateResolutionTask)); err != nil {
			return err
		}

		if err := e.File.SetCellValue(nameSheet, fmt.Sprintf("K%v", lineNumber), parseDateTime(dateUpdateTask)); err != nil {
			return err
		}
	}

	return nil
}

// setTableIssue запись данных в ячейки
func (e *Excel) setTableIssueUsersData(nameSheetUser, host string, issues *[]jira.Issue) error {
	userIssues := make([]jira.Issue, 0)
	
	// Фильтрируем задачи. Выбираем задчи для текущего пользователя 
	for _, issue := range *issues {
		if nameSheetUser != issue.Fields.Assignee.DisplayName {
			continue
		} else {
			userIssues = append(userIssues, issue)
		}
	}

	// Записываем задачи данные
	for key, issue := range userIssues {
		dateCreateTask, err := issue.Fields.Created.MarshalJSON()
		if err != nil {
			return err
		}

		dateResolutionTask, err := issue.Fields.Resolutiondate.MarshalJSON()
		if err != nil {
			return err
		}

		dateUpdateTask, err := issue.Fields.Updated.MarshalJSON()
		if err != nil {
			return err
		}

		lineNumber := key + 2 // Записываем со второй строчке

		if err := e.File.SetCellValue(nameSheetUser, fmt.Sprintf("A%d", lineNumber), keyTaskTag(issue.Key)); err != nil {
			return err
		}

		if err := e.File.SetCellValue(nameSheetUser, fmt.Sprintf("B%v", lineNumber), issue.Key); err != nil {
			return err
		}

		if err := e.File.SetCellFormula(nameSheetUser, fmt.Sprintf("C%v", lineNumber), `HYPERLINK("`+host+`/browse/`+issue.Key+`")`); err != nil {
			return err
		}

		if err := e.File.SetCellValue(nameSheetUser, fmt.Sprintf("D%v", lineNumber), issue.Fields.Summary); err != nil {
			return err
		}

		if err := e.File.SetCellValue(nameSheetUser, fmt.Sprintf("E%v", lineNumber), issue.Fields.Type.Name); err != nil {
			return err
		}

		if err := e.File.SetCellValue(nameSheetUser, fmt.Sprintf("F%v", lineNumber), issue.Fields.Status.Name); err != nil {
			return err
		}

		if err := e.File.SetCellValue(nameSheetUser, fmt.Sprintf("G%v", lineNumber), issue.Fields.Creator.DisplayName); err != nil {
			return err
		}

		if err := e.File.SetCellValue(nameSheetUser, fmt.Sprintf("H%v", lineNumber), issue.Fields.Assignee.DisplayName); err != nil {
			return err
		}

		if err := e.File.SetCellValue(nameSheetUser, fmt.Sprintf("I%v", lineNumber), parseDateTime(dateCreateTask)); err != nil {
			return err
		}

		if err := e.File.SetCellValue(nameSheetUser, fmt.Sprintf("J%v", lineNumber), parseDateTime(dateResolutionTask)); err != nil {
			return err
		}

		if err := e.File.SetCellValue(nameSheetUser, fmt.Sprintf("K%v", lineNumber), parseDateTime(dateUpdateTask)); err != nil {
			return err
		}
	}

	return nil
}

// setTableProjectTitle заголовок таблице в листе в проект разработчика
func (e *Excel) setTableProjectTitle(nameSheet string) error {
	if err := e.File.SetCellValue(nameSheet, "A1", "Сотрудники"); err != nil {
		return err
	}
	if err := e.File.SetCellValue(nameSheet, "B1", "Ключи задач"); err != nil {
		return err
	}
	if err := e.File.SetCellValue(nameSheet, "C1", "Продукт"); err != nil {
		return err
	}

	return nil
}

// setTableProject запись данных в лист проект разработчиков
func (e *Excel) setTableProjectData(nameSheet string, issues *[]jira.Issue, groupUsers *[]string) error {
	for keyUser, user := range *groupUsers {

		// Уникальные Продукты
		productList := make([]string, 0)
		productExists := make(map[string]bool)

		// Уникальные Ключи задач (Теги)
		taskTagList := make([]string, 0)
		taskTagExists := make(map[string]bool)

		for _, issue := range *issues {
			if user == issue.Fields.Assignee.DisplayName {

				nameKey := keyTaskTag(issue.Key)

				switch nameKey {
				case "BFBV2", "SPR", "DATA":
					addUnique(productExists, &productList, "DSA")
					addUnique(taskTagExists, &taskTagList, nameKey)

				case "BIT", "ETP", "IP", "PIM", "DEVD":
					addUnique(productExists, &productList, "PIM")
					addUnique(taskTagExists, &taskTagList, nameKey)

				case "MQ":
					addUnique(productExists, &productList, "Mediaquad")
					addUnique(taskTagExists, &taskTagList, nameKey)

				case "INCOME":
					addUnique(productExists, &productList, "Любой продукт, зависит от задачи")
					addUnique(taskTagExists, &taskTagList, nameKey)

				default:
					productList = nil
					taskTagList = nil
				}
			}
		}

		lineNumber := keyUser + 2 // Записываем со второй строчке

		if err := e.File.SetCellValue(nameSheet, fmt.Sprintf("A%d", lineNumber), user); err != nil {
			return err
		}

		if err := e.File.SetCellValue(nameSheet, fmt.Sprintf("B%d", lineNumber), strings.Join(taskTagList, ", ")); err != nil {
			return err
		}

		if err := e.File.SetCellValue(nameSheet, fmt.Sprintf("C%d", lineNumber), strings.Join(productList, ", ")); err != nil {
			return err
		}

	}

	return nil
}


