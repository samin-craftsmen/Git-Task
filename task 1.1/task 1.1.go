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

func viewTasks() {
	if len(tasks) == 0 {
		fmt.Println("No tasks available")
		return
	}

	for i, t := range tasks {
		fmt.Println(i+1, "-", t)
	}
}
