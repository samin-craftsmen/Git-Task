package task_1_3

import "fmt"

func cToF(c float64) float64 {
	return (c*9/5 + 32)
}

func fToC(f float64) float64 {
	return (f - 32) * 5 / 9
}

func showMenu() {
	fmt.Println("\n1. Celsius to Fahrenheit")
	fmt.Println("2. Fahrenheit to Celsius")
	fmt.Println("3. Exit")
}
