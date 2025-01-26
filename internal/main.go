package internal

import (
	"errors"
	"fmt"
	"os"
	"report/internal/models/excel"
	jr "report/internal/models/jira"
	"report/pkg/config"
	"time"
)

// GetReport получить отчет. Запуск приложения
func GetReport(c *config.Config) error {
	timerStart := time.Now()

	statusApp("Этап 1/4. Проверка параметров")

	// Проверка конфига
	if err := checkConfig(c); err != nil {
		return err
	}

	statusApp("Этап 2/4. Запрос к API: " + c.Host)

	// Авторизация в Jira Cloud
	client, err := jr.NewJira(c)
	if err != nil {
		return err
	}

	issueChan := make(chan jr.IssueChan)
	groupUserChan := make(chan jr.GroupUsersChan)

	jql := jr.JQL{
		Status:    c.Status,
		DateStart: c.DateStart,
		DateEnd:   c.DateEnd,
		Group:     c.Group,
	}

	
	go client.GetTasks(jql.UpdateString(), issueChan) 	// Получить список задач
	go client.GetGroupUsers(c.Group, groupUserChan) 	// Получить список сотрудников состоящей в группе

	issue := <-issueChan
	if issue.Err != nil {
		return issue.Err
	}

	groupUser := <-groupUserChan
	if groupUser.Err != nil {
		return groupUser.Err
	}

	statusApp("Этап 3/4. Выгрузка задач в Excel файл")

	excelFile := excel.NewExcel()

	// Создаем общий лист задач
	if err := excelFile.CreateTodoList(&issue.List, c.Host); err != nil {
		return err
	}

	// CreateProjectDevOps создать лист с проектами разработчика
	if err := excelFile.CreateProjectDevOps(&issue.List, &groupUser.List); err != nil {
		return err
	}

	// Создаем лист на каждого пользователя
	if err := excelFile.CreateTodoListUsers(issue.List, groupUser.List, c.Host); err != nil {
		return nil
	}

	// Сохраняем файл
	fileName := fmt.Sprintf("%v %v - %v.xlsx", c.Group, c.DateStart, c.DateEnd)
	if err := excelFile.SaveFile(fileName); err != nil {
		return err
	}

	statusApp(fmt.Sprintf("Этап 4/4. Готово\n\nКол-во задач: %d шт.\nКол-во сотрудников: %d\nФайл: %v\nВремя выполнения скрипта %.2fs\n", len(issue.List), len(groupUser.List), fileName, time.Since(timerStart).Seconds()))

	return nil
}

// checkConfig проверка конфига (Параметры)
func checkConfig(c *config.Config) error {
	if len(os.Args) <= 1 {
		return errors.New("укажите параметры (Флаги). Подробно: --help")
	}

	if err := checkFormatDate(c.DateStart); err != nil {
		return err
	}

	if err := checkFormatDate(c.DateEnd); err != nil {
		return err
	}

	return nil
}

// checkFormatDate проверка формат даты в виде: ГГГГ-ММ-ДД
func checkFormatDate(date string) error {
	_, err := time.Parse(config.FormatDate, date)
	if err != nil {
		return errors.New("неверный формат даты. Образец: " + config.FormatDate + " 'ГГГГ-ММ-ДД'")
	}

	return nil
}

// statusApp выводит текущий статус в консоль
func statusApp(message string) {
	fmt.Println(message)
}
