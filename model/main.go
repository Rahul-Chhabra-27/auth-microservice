package model

import "time"

type Task struct {
	ID int
	Name string
	Description string
	Location string
	DateTime time.Time
	UserId int
}

var tasks [] Task = []Task{};

func (task Task) SaveTask() {
	tasks = append(tasks, task);
}

func GetAllTheTasks() [] Task {
	return tasks;
}