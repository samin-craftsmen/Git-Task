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

func convertC() {
	var c float64
	fmt.Print("Enter Celsius: ")
	fmt.Scan(&c)
	fmt.Printf("%.2f°C = %.2f°F\n", c, cToF(c))
}

func convertF() {
	var f float64
	fmt.Print("Enter Fahrenheit: ")
	fmt.Scan(&f)
	fmt.Printf("%.2f°F = %.2f°C\n", f, fToC(f))
}

func main_loop() {
	for {
		showMenu()
		var choice int
		fmt.Print("Choose: ")
		fmt.Scan(&choice)

		switch choice {
		case 1:
			convertC()
		case 2:
			convertF()
		case 3:
			fmt.Println("Goodbye!")
			return
		default:
			fmt.Println("Invalid option")
		}
	}
}
