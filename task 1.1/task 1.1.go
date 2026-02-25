package task_1_1

import "fmt"

var tasks []string

func addTask() {
	var task string
	fmt.Print("Enter task: ")
	fmt.Scan(&task)

	tasks = append(tasks, task)
	fmt.Println("Task added!")
}
