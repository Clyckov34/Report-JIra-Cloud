package jira

import "fmt"

type JQL struct {
	Status    string
	DateStart string
	DateEnd   string
	Group     string
}

// UpdateString jql код. Фильтр по обновленным задачам
func (j *JQL) UpdateString() string {
	const filter = `status in (%v) AND updated >= "%v 00:00" AND updated <= "%v 23:59" AND assignee in (membersOf("%v")) ORDER BY updated DESC`
	return fmt.Sprintf(filter, j.Status, j.DateStart, j.DateEnd, j.Group)
}

// // CreateString Фильтр по созданным задачам
// func (j *JQL) CreateString() string {
// 	const filter = `status in (%v) AND created >= "%v 00:00" AND created <= "%v 23:59" AND assignee in (membersOf("%v")) ORDER BY created DESC`
// 	return fmt.Sprintf(filter, j.Status, j.DateStart, j.DateEnd, j.Group)
// }
