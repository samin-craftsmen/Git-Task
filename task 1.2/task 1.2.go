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
