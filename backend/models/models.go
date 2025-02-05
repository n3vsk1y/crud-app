package models

type Task struct {
	ID			int		`json:"id"`
	Title		string	`json:"title"`
	Complited	int		`json:"complited"`
}