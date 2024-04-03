package main
import (
	"fmt"
	"time"
)

var x float64
var y float64
func main() {
	var option string
	fmt.Print("What operation do you want to do? (multiply, divide, add, subtract, exit): ")
	fmt.Scan(&option)
	if option == "multiply" {
		multiply()
	}
	if option == "divide" {
		divide()
	}
	if option == "add" {
		add()
	}
	if option == "subtract" {
		subtract()
	}
	if option == "exit" {
		time.Sleep(1 * time.Second)
		fmt.Println("Exiting...")
	}
}
func multiply() {
	fmt.Println("Number Multiplier!!")
	fmt.Print("Type a number: ")
	fmt.Scan(&x)
	fmt.Print("Type another number: ")
	fmt.Scan(&y)
	fmt.Println("Result: ", x * y)
	main()
}
func divide() {
	fmt.Println("Number Divider!!")
	fmt.Print("Type a number: ")
	fmt.Scan(&x)
	fmt.Print("Type another number: ")
	fmt.Scan(&y)
	fmt.Println("Result: ", x / y)
	main()
}
func add() {
	fmt.Println("Number Adder!!")
	fmt.Print("Type a number: ")
	fmt.Scan(&x)
	fmt.Print("Type another number: ")
	fmt.Scan(&y)
	fmt.Println("Result: ", x + y)
	main()
}
func subtract() {
	fmt.Println("Number Subtractor!!")
	fmt.Print("Type a number: ")
	fmt.Scan(&x)
	fmt.Print("Type another number: ")	
	fmt.Scan(&y)
	fmt.Println("Result: ", x - y)
	main()
}