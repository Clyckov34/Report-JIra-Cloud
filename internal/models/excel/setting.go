package excel

import (
	"report/internal/config"
	"strings"
	"time"
)

// setStyle установка стилей
func (e *Excel) setStyle(nameSheet string) error {
	if err := e.File.SetColWidth(nameSheet, "A", "B", 20); err != nil {
		return err
	}
	if err := e.File.SetColWidth(nameSheet, "C", "D", 60); err != nil {
		return err
	}
	if err := e.File.SetColWidth(nameSheet, "E", "F", 15); err != nil {
		return err
	}
	if err := e.File.SetColWidth(nameSheet, "G", "K", 25); err != nil {
		return err
	}

	return nil
}

// keyTaskTag преобразовывает номер ключа ABCD-1234 -> ABCD
func keyTaskTag(nameKey string) string {
	key := strings.Split(nameKey, "-")
	return key[0]
}

// parseDateTime обработка даты в формат
func parseDateTime(datetime []byte) string {
	h, _ := time.Parse("\"2006-01-02T15:04:05.000-0700\"", string(datetime))
	timeString := h.Format(config.FormatDate)

	if timeString != "0001-01-01" {
		return timeString
	}

	return "-"
}

// addUnique добавляет уникальные элементы в массив
func addUnique(exists map[string]bool, array *[]string, name string) {
	if !exists[name] {
		*array = append(*array, name)
		exists[name] = true
	}
}
