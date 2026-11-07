package main

import (
	"fmt"
	"os"
	"strconv"
)

func main() {
	//Handle args length
	if len(os.Args) != 4 {
		fmt.Println("Usage: go run main.go [operation] [number1] [number2]")
		return
	}

	opt := os.Args[1]
	n1, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		fmt.Println("Error: Number 1 invalid")
		return
	}
	n2, err := strconv.ParseFloat(os.Args[3], 64)
	if err != nil {
		fmt.Println("Error: Number 2 invalid")
		return
	}
	var res float64

	switch opt {
	case "add":
		res = add(n1, n2)
	case "subs":
		res = sub(n1, n2)
	case "mult":
		res = mult(n1, n2)
	case "div":
		res, err = div(n1, n2)
		if err != nil {
			fmt.Println(err)
			return
		}
	default:
		fmt.Println("Invalid operation")
		return
	}

	fmt.Printf("Result = %.2f", res)
}

func add(n1, n2 float64) float64 {
	return n1 + n2
}
func sub(n1, n2 float64) float64 {
	return n1 - n2
}
func mult(n1, n2 float64) float64 {
	return n1 * n2
}
func div(n1, n2 float64) (float64, error) {
	if n2 == 0 {
		return 0, fmt.Errorf("Error: Cannot divide by zero")
	}

	return n1 / n2, nil
}
