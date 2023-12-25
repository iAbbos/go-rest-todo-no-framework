package models

type Status string

const (
	Planned Status = "planned"
	Doing   Status = "doing"
	Done    Status = "done"
)

type Task struct {
	ID     int
	Name   string
	Note   string
	Date   string
	Status Status
}
