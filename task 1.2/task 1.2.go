package task_1_2

import "fmt"

var counter int

func increment() {
	counter++
	fmt.Println("Counter:", counter)
}

func decrement() {
	counter--
	fmt.Println("Counter:", counter)
}

func reset() {
	counter = 0
	fmt.Println("Counter reset to 0")
}

func showMenu() {
	fmt.Println("\n1. Increment")
	fmt.Println("2. Decrement")
	fmt.Println("3. Reset")
	fmt.Println("4. Exit")
}
