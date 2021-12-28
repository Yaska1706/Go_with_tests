package tasks

import "fmt"

type task struct {
	name   string
	status bool
}

var tasks []task

//Add task
func (t *task) AddTask(tk task) []task {
	tasks = append(tasks, tk)
	return tasks
}

//display task
func (t *task) Display() {
	for _, task := range tasks {
		fmt.Sprintf("Name: %s \n Status: %t \n", task.name, task.status)
	}
}
