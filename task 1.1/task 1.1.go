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

func deleteTask() {
	var index int
	fmt.Print("Enter task number to delete: ")
	fmt.Scan(&index)

	if index < 1 || index > len(tasks) {
		fmt.Println("Invalid number")
		return
	}

	tasks = append(tasks[:index-1], tasks[index:]...)
	fmt.Println("Task deleted!")
}

func showMenu() {
	fmt.Println("\n1. Add Task")
	fmt.Println("2. View Tasks")
	fmt.Println("3. Delete Task")
	fmt.Println("4. Exit")
}
