package config

import (
	"flag"
	"time"
)

const FormatDate = "2006-01-02" // Формат даты в виде: ГГГГ-ММ-ДД

type Config struct {
	Host      string
	UserName  string
	Token     string
	DateStart string
	DateEnd   string
	Group     string
	Status    string
}

// GetConfig Параметры запуска приложения
func GetConfig() *Config {
	flags := Config{}

	flag.StringVar(&flags.Host, "host", "", "Хост Jira Cloud")
	flag.StringVar(&flags.UserName, "user", "", "Логин пользователя")
	flag.StringVar(&flags.Token, "token", "", "Токен")
	flag.StringVar(&flags.DateStart, "date_start", nowDate(), "Начальная дата запроса. Формат: YYYY-MM-DD")
	flag.StringVar(&flags.DateEnd, "date_end", nowDate(), "Конечная дата запроса. Формат: YYYY-MM-DD")
	flag.StringVar(&flags.Group, "group", "", "Группа пользователей в Jira")
	flag.StringVar(&flags.Status, "status", `"To Do", "In Progress", "Done", "Denied", "Pause"`, "Статус задач")

	flag.Parse()
	return &flags
}

// nowDate текущая дата
func nowDate() string {
	return time.Now().Format(FormatDate)
}
